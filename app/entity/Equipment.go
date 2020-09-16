package entity

import "database/sql"

type Equipment struct {
	Id              int ` json:"id" `
	Fk_parent       int ` json:"class" `
	Fk_class        int ` json:"subclass" `
	StatusI         int ` json:"statusI" `
	Fk_user         sql.NullInt64
	Fk_user1        int    ` json:"fk_user" `
	InventoryNumber string ` json:"inventoryNumber" `
	EquipmentName   string ` json:"name" `
	Status          string `json:"status"`
	Subclass        string `json:"Subclass"`
	Class           string `json:"Class"`
	UserFIO         string     ` json:"user" `
	ClassName       string     ` json:"value" `
	Data            []Subclass ` json:"data" `
	Open            bool       `json:"open"`
}
type Subclass struct {
	Id           int    ` json:"subclass" `
	SubclassName string ` json:"value" `
}
type FullTree struct {
	Id int `json:"class"`
	Value     string      ` json:"value" `
	Equipment []Equipment ` json:"data" `
	Open      bool        `json:"open"`
}
