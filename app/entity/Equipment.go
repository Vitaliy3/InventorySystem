package entity

import "database/sql"

//структура оборудования
type Equipment struct {
	Id              int ` json:"id" `
	Fk_parent       int ` json:"class" `              //внешний ключ на класс
	Fk_class        int ` json:"subclass" `           //внешний ключ на подкласс
	StatusI         int ` json:"statusI" `            //статус
	Fk_user         sql.NullInt64                     //внешний ключ на пользователя
	Fk_userI        int    ` json:"fk_user" `         //внешний ключ на пользователя
	InventoryNumber string ` json:"inventoryNumber" ` //инвентарный номер
	EquipmentName   string ` json:"name" `            //назнавание оборудования
	Status          string `json:"status"`            //статус
	Subclass        string `json:"Subclass"`          //название подкласса
	Class           string `json:"Class"`             //название класса
	UserFIO         string ` json:"user" `            //ФИО сотрудника
}

//структура древовидного списка
type TreeClass struct {
	Fk_parent int        ` json:"class" `    //внешний ключ на клас
	Fk_class  int        ` json:"subclass" ` //внешний ключ на подкласс
	Subclass  string     `json:"Subclass"`   //название подкласса
	ClassName string     ` json:"value" ` //название класса
	Data      []Subclass ` json:"data" `
	Open      bool       `json:"open"` //развернутый/свернутый узед дерева
}

//структура подкласса
type Subclass struct {
	Id           int    ` json:"subclass" `
	SubclassName string ` json:"value" `  //название подкласса
}

//структура итогового дерева
type FullTree struct {
	Id    int         `json:"class"`
	Value string      ` json:"value" ` //самая верхняя папка,содержащая все узлы дерева
	Tree  []TreeClass ` json:"data" `
	Open  bool        `json:"open"` //развернутый/свернутый узед дерева
}
