package entity

type User struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Track *[]Track `json:"track" gorm:"many2many:user_tracks"`
}
