package model

type StatusType string

const (
	STATUS_ON_THE_WAY       StatusType = "В пути"
	STATUS_IN_STORAGE                  = "На складе"
	STATUS_SOLD                        = "Продан"
	STATUS_REMOVE_FROM_SALE            = "Снят с продаж"
)

type Car struct {
	Brand  string
	Model  string
	Price  int
	Status StatusType
	Milage int
}
