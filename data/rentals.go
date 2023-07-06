package data

type Rentals struct {
	Id              uint   `gorm:"primary_key","column:id"`
	UserId          int    `gorm:"column:user_id"`
	Name            string `gorm:"column:name"`
	Type            string `gorm:"column:type"`
	Description     string `gorm:"column:description"`
	PricePerDay     int    `gorm:"column:price_per_day"`
	Sleeps          int    `gorm:"column:sleeps"`
	HomeCity        string `gorm:"column:home_city"`
	HomeState       string `gorm:"column:home_state"`
	HomeZip         string `gorm:"column:home_zip"`
	HomeCountry     string `gorm:"column:home_country"`
	VehicleMake     string `gorm:"column:vehicle_make"`
	VehicleModel    string `gorm:"column:vehicle_year"`
	VehicleYear     int    `gorm:"column:vehicle_make"`
	VehicleLength   int    `gorm:"column:vehicle_length"`
	Lat             int    `gorm:"column:lat"`
	Lng             int    `gorm:"column:lng"`
	PrimaryImageUrl string `gorm:"column:primary_image_url"`
}
