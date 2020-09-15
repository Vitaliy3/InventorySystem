package entity

import "database/sql"

type InventoryEvent struct {
	Id           int    `json:"id"`
	UserFIO      string `json:"user"`
	Date         string `json:"date"`
	Equipment    string `json:"equipment"`
	Fk_user      sql.NullInt64
	Fk_equipment int
	ActionEvent  string `json:"event"`
}
