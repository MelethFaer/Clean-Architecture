package entity

type User struct {
	ID     uint     `json:"id"`
	Name   string   `json:"name"`
	Tracks *[]Track `json:"tracks" gorm:"many2many:user_tracks"`
}
