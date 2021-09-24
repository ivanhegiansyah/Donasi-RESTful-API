package main

import (
	"finalproject-BE/app/routes"
	_userUsecase "finalproject-BE/business/users"
	_userController "finalproject-BE/controllers/users"
	_userdb "finalproject-BE/drivers/databases/users"
	_userRepository "finalproject-BE/drivers/databases/users"

	_donationUsecase "finalproject-BE/business/donations"
	_donationController "finalproject-BE/controllers/donations"
	_donationdb "finalproject-BE/drivers/databases/donations"
	_donationRepository "finalproject-BE/drivers/databases/donations"
	
	_mysqlDriver "finalproject-BE/drivers/mysql"
	"time"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userdb.Users{})
	db.AutoMigrate(&_donationdb.Donations{})

}

func main() {
	//init koneksi database
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	donationRepository := _donationRepository.NewMysqlDonationRepository(Conn)
	donationUseCase := _donationUsecase.NewDonationUsecase(donationRepository, timeoutContext)
	donationController := _donationController.NewDonationController(donationUseCase)

	routesInit := routes.ControllerList{
		UserController: *userController,
		DonationController: *donationController,
	}

	routesInit.RouteUser(e)
	routesInit.RouteDonation(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}