package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"log"
	"myapp/app/mappers"
	"strings"
)

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Fk_role    int    `json:"fk_role"`
}

//получение всех сотрудников
func (e *Employee) GetAllEmployees(DB *sql.DB) (employeeArray []Employee, err error) {
	emMapper := mappers.Employee{}
	dbEmployees, err := emMapper.GetAllEmployees(DB)
	if err != nil {
		return
	}
	var temp Employee
	for _, v := range dbEmployees {
		temp.Id = v.Id
		temp.Name = v.Name
		temp.Surname = v.Surname
		temp.Patronymic = v.Patronymic
		temp.Login = v.Login
		employeeArray = append(employeeArray, temp)
	}
	return
}

//изменение данных о сотруднике
func (e *Employee) UpdateEmployee(DB *sql.DB, params *revel.Params) (employee Employee, err error) {
	employeeMapper := mappers.Employee{}
	err = json.Unmarshal(params.JSON, &employeeMapper)
	if err != nil {
		return
	}
	lastUpdateId, err := employeeMapper.UpdateEmployee(DB)
	if err != nil {
		return
	}
	employeeMapper, err = employeeMapper.GetEmployeeById(DB, lastUpdateId)
	employee.Id = employeeMapper.Id
	employee.Name = employeeMapper.Name
	employee.Surname = employeeMapper.Surname
	employee.Patronymic = employeeMapper.Patronymic
	employee.Login = employeeMapper.Login
	if err != nil {
		return
	}
	return
}

//удаление сотрудника
func (e *Employee) DeleteEmployee(DB *sql.DB, params *revel.Params) (employee Employee, err error) {
	emMapper := mappers.Employee{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}
	lastDeleteId, err := emMapper.DeleteEmployee(DB, id)
	if err != nil {
		return
	}
	employee.Id = lastDeleteId
	return
}

func (e *Employee) AddEmployee(DB *sql.DB, params *revel.Params) (employee Employee, err error) {
	employeeMapper := mappers.Employee{}
	err = json.Unmarshal(params.JSON, &employeeMapper)
	if err != nil {
		return
	}
	employeeMapper.Password = HashAndSalt([]byte(employeeMapper.Password))
	lastInsertedId, err := employeeMapper.AddEmployee(DB)
	if err != nil {
		return
	}
	employeeMapper, err = employeeMapper.GetEmployeeById(DB, lastInsertedId)
	if err != nil {
		return
	}
	employee.Id = employeeMapper.Id
	employee.Name = employeeMapper.Name
	employee.Surname = employeeMapper.Surname
	employee.Patronymic = employeeMapper.Patronymic
	employee.Login = employeeMapper.Login
	return
}

func (e *Employee) ResetPassEmployee(DB *sql.DB, params *revel.Params) (employee Employee, err error) {
	emMapper := mappers.Employee{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	updatedRowId, err := emMapper.ResetPassEmployee(DB, id)
	if err != nil {
		return
	}
	if id != updatedRowId {
		err = errors.New("not reset")
		return
	}
	return
}

type Authorization struct {
	Id int
	Token string
	Role  string
}

func (e *Employee) Auth(DB *sql.DB, c *revel.Controller) (authStruct Authorization, err error) {
	employeeMap := mappers.Employee{}
	var login, password string
	cookies, _ := c.Request.Cookie("auth")
	splitCookie := strings.Split(cookies.GetValue(), ":")
	decoded, _ := base64.StdEncoding.DecodeString(splitCookie[0])
	authData := strings.Split(string(decoded), ":")
	login = authData[0]
	password = authData[1]
	user, err := employeeMap.Auth(DB, login)
	if err != nil {
		return
	}
	if ComparePasswords(password, []byte(user.Password)) {
		err = errors.New("Неверное имя пользователя или пароль")
		return
	}
	userRole, err := employeeMap.GetUserRoleById(DB, user.Id)
	fmt.Println("ROLE", userRole)
	fmt.Println("id", user.Id)
	if err != nil {
		return
	}
	newToken := gravatarMD5(login)
	authStruct.Id=user.Id
	authStruct.Role = userRole
	authStruct.Token = newToken
	return
}
func gravatarMD5(login string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(login)))
	return hex.EncodeToString(h.Sum(nil))
}

func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
