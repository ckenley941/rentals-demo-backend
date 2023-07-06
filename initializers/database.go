package initializers

//"fmt"
// "gorm.io/driver/postgres"
import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error
	//config := "postgres://postgres:postgres@localhost:5432/postgres"
	config := "host=127.0.0.1 user=root password=root dbname=testingwithrentals port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	fmt.Println(("Connecting to db"))
	if err != nil {
		//log.Fatal("Failed to connect to database");
		fmt.Println(("Failed to connect to db"))
	}

	return db
}
