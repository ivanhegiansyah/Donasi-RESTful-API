package main

import (
	
	_userUsecase "finalproject-BE/business/users"
	_userController "finalproject-BE/controllers/users"
	_userRepository "finalproject-BE/drivers/databases/users"
	_userdb "finalproject-BE/drivers/databases/users"

	_donationUsecase "finalproject-BE/business/donations"
	_donationController "finalproject-BE/controllers/donations"
	_donationRepository "finalproject-BE/drivers/databases/donations"
	_donationdb "finalproject-BE/drivers/databases/donations"

	_donationDetailUsecase "finalproject-BE/business/donationDetails"
	_donationDetailController "finalproject-BE/controllers/donationDetails"
	_donationDetailRepository "finalproject-BE/drivers/databases/donationDetails"
	_donationdetaildb "finalproject-BE/drivers/databases/donationDetails"

	_donationTypeUsecase "finalproject-BE/business/donationTypes"
	_donationTypeController "finalproject-BE/controllers/donationTypes"
	_donationTypeRepository "finalproject-BE/drivers/databases/donationTypes"
	_donationtypedb "finalproject-BE/drivers/databases/donationTypes"

	_route "finalproject-BE/app/routes"
	_middleware "finalproject-BE/app/middlewares"
	_mysqlDriver "finalproject-BE/drivers/mysql"
	"log"
	"time"

	
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
	db.AutoMigrate(&_donationdetaildb.DonationDetails{})
	db.AutoMigrate(&_donationtypedb.DonationTypes{})
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

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext,  &configJWT)
	userController := _userController.NewUserController(userUseCase)

	donationRepository := _donationRepository.NewMysqlDonationRepository(Conn)
	donationUseCase := _donationUsecase.NewDonationUsecase(donationRepository, timeoutContext)
	donationController := _donationController.NewDonationController(donationUseCase)

	donationDetailRepository := _donationDetailRepository.NewMysqlDonationDetailRepository(Conn)
	donationDetailUseCase := _donationDetailUsecase.NewDonationDetailUsecase(donationDetailRepository, timeoutContext)
	donationDetailController := _donationDetailController.NewDonationDetailController(donationDetailUseCase)

	donationTypeRepository := _donationTypeRepository.NewMysqlDonationTypeRepository(Conn)
	donationTypeUseCase := _donationTypeUsecase.NewDonationTypeUsecase(donationTypeRepository, timeoutContext)
	donationTypeController := _donationTypeController.NewDonationTypeController(donationTypeUseCase)

	routesInit := _route.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:           *userController,
		DonationController:       *donationController,
		DonationDetailController: *donationDetailController,
		DonationTypeController:   *donationTypeController,
	}

	routesInit.RouteUser(e)
	routesInit.RouteDonation(e)
	routesInit.RouteDonationDetail(e)
	routesInit.RouteDonationType(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
