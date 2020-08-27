package Models

import (
	"log"
	"myapp/app/mappers"
)

type EquipmentModel struct {
	Id              int    ` json:"id" `
	Fk_class        int    ` json:"class" `
	Fk_subclass     int    ` json:"sublcass" `
	UserFIO         string ` json:"user" `
	InventoryNumber string ` json:"inventoryNumber" `
	EquipmentName   string ` json:"name" `
	Status          int    ` json:"status" `
}

func (e *EquipmentModel) WriteEquipmentModel(id int) (status string) {
	eqMapper := mappers.EquipmentTable{}
	result, err := eqMapper.WriteEquipment(id)
	if err != nil {
		log.Println(err)
	}
	if result > 0 {
		row, err := eqMapper.GetEquipmentById(id)
		if err != nil {
			log.Println(err)
		}
		if row.Status == 2 {
			status = "Списан"
		}
		return
	}
	return
}

func (e *EquipmentModel) GetAllEquipments() []EquipmentModel {
	eqMapper := mappers.EquipmentTable{}
	dbEqupments, err := eqMapper.GetAllEquipments()
	if err != nil {
		log.Println(err)
	}
	equipments := make([]EquipmentModel, 0)
	var temp EquipmentModel
	for _, v := range dbEqupments {
		temp.Id = v.Id
		temp.Fk_class = v.Id
		temp.InventoryNumber = v.InventoryNumber
		temp.EquipmentName = v.EquipmentName
		temp.Status = v.Status
		equipments = append(equipments, temp)
	}
	return equipments
}
