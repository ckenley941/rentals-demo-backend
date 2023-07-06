package data

type Users struct {
	Id        uint   `gorm:"primary_key","column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}
