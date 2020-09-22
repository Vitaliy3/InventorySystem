package entity

import "database/sql"

//структура событий выдачи
type InventoryEvent struct {
	Id           int    `json:"id"`
	UserFIO      string `json:"user"`      //ФИО сотрудника
	Date         string `json:"date"`      //дата регистрации события
	Equipment    string `json:"equipment"` //название товара
	Fk_user      sql.NullInt64             //внешний ключ на пользователя
	Fk_userI     int                       //внешний ключ на пользователя
	Fk_equipment int                       //внешний ключ на оборудование
	ActionEvent  string `json:"event"`     //описание события
}
