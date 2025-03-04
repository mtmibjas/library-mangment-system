package entities

type Role struct {
	ID   uint   `json:"primaryKey"`
	Name string `json:"unique;not null"`
}
