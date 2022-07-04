package entity

//User represents users table in database
type Role struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
}
