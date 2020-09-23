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
	employeeMapper mappers.Employee
}

//получение всех сотрудников
func (e *Employee) GetAll(DB *sql.DB) (employees []entity.Employee, err error) {
	employees, err = e.employeeMapper.GetAll(DB)
	if err != nil {
		return
	}
	return
}

//изменение данных о сотруднике
func (e *Employee) Update(db *sql.DB, employeeIn entity.Employee) (employeeOut entity.Employee, err error) {
	var password string
	if employeeIn.Password == "" {
		password, err = e.employeeMapper.GetPasswordById(db, employeeIn.Id)
		employeeIn.Password = password

		if err != nil {
			return
		}
	} else {
		employeeIn.Password = hashAndSalt([]byte(employeeIn.Password)) //получаем хеш пароля
	}

	lastUpdateId, err := e.employeeMapper.Update(db, employeeIn)
	if err != nil {

		return
	}

	employeeOut, err = e.employeeMapper.GetById(db, lastUpdateId)
	if err != nil {
		return
	}
	return
}

//удаление сотрудника
func (e *Employee) Delete(db *sql.DB, employeeIn entity.Employee) (employeeOut entity.Employee, err error) {
	eventMapper := mappers.InventoryEvent{}
	equipmentMapper := mappers.Equipment{}
	result, err := equipmentMapper.GetByUserId(db, entity.Equipment{Id: employeeIn.Id})
	if result != nil {
		err = errors.New("Невозможно удалить, на сотруднике закреплено оборудование")
		return
	}
	_, err = eventMapper.DeleteByFkUser(db, employeeIn)
	if err != nil {
		return
	}
	lastDeleteId, err := e.employeeMapper.Delete(db, employeeIn)
	if err != nil {
		return
	}
	employeeOut.Id = lastDeleteId
	return
}

//добавление сотрудника
func (e *Employee) Add(db *sql.DB, employeeIn entity.Employee) (employeeOut entity.Employee, err error) {
	employeeIn.Password = hashAndSalt([]byte(employeeIn.Password)) //получение хеша пароля

	lastInsertedId, err := e.employeeMapper.Add(db, employeeIn)
	if err != nil {
		return
	}

	employeeOut, err = e.employeeMapper.GetById(db, lastInsertedId)
	if err != nil {
		return
	}
	return
}

//сброс пароля пользователя
func (e *Employee) ResetPassword(db *sql.DB, employeeIn entity.Employee) (employeeOut entity.Employee, err error) {
	employeeIn.Password = hashAndSalt([]byte("123")) //получение хеша пароля

	updatedRowId, err := e.employeeMapper.ResetPassword(db, employeeIn)
	if err != nil {
		return
	}

	if employeeOut.Id != updatedRowId {
		err = errors.New("not reset")
		return
	}
	return
}

//авторизация пользователя
func (e *Employee) Auth(db *sql.DB, authIn entity.Authorization) (authOut entity.Authorization, err error) {
	user, err := e.employeeMapper.GetByLogin(db, authIn.Login)
	if err != nil {
		return
	}

	if !comparePasswords(user.Password, []byte(authIn.Password)) { //сравение паролей
		err = errors.New("Неверное имя пользователя или пароль")
		return
	}

	userRole, err := e.employeeMapper.GetUserRoleById(db, user.Id)
	if err != nil {
		return
	}

	newToken := gravatarMD5(authIn.Login)
	authOut.Id = user.Id
	authOut.Role = userRole
	authOut.Token = newToken
	return
}

//функция создания токена
func gravatarMD5(login string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(login)))
	return hex.EncodeToString(h.Sum(nil))
}

//функция создания хеша пароля
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//функция сравнения паролей
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("err compare passwords", err)
		return false
	}
	return true
}
