package providers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"myapp/app/entity"
	"myapp/app/mappers"
	"strings"
)

type Employee struct {
	entity.Employee
}

//получение всех сотрудников
func (e *Employee) GetAllEmployees(DB *sql.DB) (employees []entity.Employee, err error) {
	mapper := mappers.Employee{}
	employees, err = mapper.GetAllEmployees(DB)
	if err != nil {
		return
	}
	return
}

//изменение данных о сотруднике
func (e *Employee) UpdateEmployee(DB *sql.DB, employeeIn entity.Employee) (employeeOut entity.Employee, err error) {
	mapper := mappers.Employee{}
	employeeIn.Password = HashAndSalt([]byte(employeeIn.Password))

	lastUpdateId, err := mapper.UpdateEmployee(DB, employeeIn)
	if err != nil {
		return
	}
	employeeOut, err = mapper.GetEmployeeById(DB, lastUpdateId)
	if err != nil {
		return
	}
	return
}

//удаление сотрудника
func (e *Employee) DeleteEmployee(DB *sql.DB, employeeIn entity.Employee) (employeeOut entity.Employee, err error) {
	emMapper := mappers.Employee{}
	eventMapper := mappers.InventoryEvent{}

	_, err = eventMapper.DeleteEventByFkUser(DB, employeeIn)
	if err != nil {
		return
	}
	lastDeleteId, err := emMapper.DeleteEmployee(DB, employeeIn)
	if err != nil {
		return
	}
	employeeOut.Id = lastDeleteId
	return
}

func (e *Employee) AddEmployee(DB *sql.DB, employeeIn entity.Employee) (employeeOut	 entity.Employee, err error) {
	employeeIn.Password = HashAndSalt([]byte(employeeIn.Password))
	employeeMap := mappers.Employee{}
	lastInsertedId, err := employeeMap.AddEmployee(DB, employeeIn)
	if err != nil {
		return
	}
	employeeOut, err = employeeMap.GetEmployeeById(DB, lastInsertedId)
	if err != nil {
		return
	}
	return
}

func (e *Employee) ResetPassEmployee(DB *sql.DB,employeeIn entity.Employee) (employeeOut entity.Employee, err error) {
	emMapper := mappers.Employee{}
	employeeIn.Password=HashAndSalt([]byte("123"))
	updatedRowId, err := emMapper.ResetPassEmployee(DB, employeeIn)
	if err != nil {
		return
	}
	if employeeOut.Id != updatedRowId {
		err = errors.New("not reset")
		return
	}
	return
}

func (e *Employee) Auth(DB *sql.DB, authIn entity.Authorization) (authOut entity.Authorization, err error) {
	employeeMap := mappers.Employee{}
	user, err := employeeMap.GetEmployeeByLogin(DB, authIn.Login)
	if err != nil {
		return
	}
	if !ComparePasswords(user.Password, []byte(authIn.Password)) {
		err = errors.New("Неверное имя пользователя или пароль")
		return
	}
	userRole, err := employeeMap.GetUserRoleById(DB, user.Id)
	if err != nil {
		return
	}
	newToken := gravatarMD5(authIn.Login)
	authOut.Id = user.Id
	authOut.Role = userRole
	authOut.Token = newToken
	return
}

//создает токен
func gravatarMD5(login string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(login)))
	return hex.EncodeToString(h.Sum(nil))
}

//создает хеш для пароля
func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//сравнивает хеш пароля и пароль
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("err in compare", err)
		return false
	}
	return true
}
