package routes

import (
	"tugas/praktikum/controller"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController, secret string) {
	// var user = e.Group("/users")
	// user.POST(c.Register(), )
	e.POST("/Login", uc.Login())
	e.POST("/users", uc.Register())
	// e.GET("/users", uc.GetAllUsers(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/users/:id", uc.GetUserByID(), echojwt.JWT([]byte(secret)))
	e.DELETE("/users/:id", uc.DeleteUser(), echojwt.JWT([]byte(secret)))
}

func AspirationRoutes(e *echo.Echo, ac controller.AspirationsController, secret string) {
	e.POST("/aspirations", ac.InsertAspiration(), echojwt.JWT([]byte(secret)))
	e.DELETE("/aspirations/:id", ac.DeleteAspiration(), echojwt.JWT([]byte(secret)))
	e.GET("/aspirations", ac.GetAllAspirations())
	e.GET("/aspirations/:id", ac.GetAspirationByID())
}

func DrivertrainRoutes(e *echo.Echo, dc controller.DrivetrainController, secret string) {
	e.POST("/drivetrains", dc.InsertDrivetrain(), echojwt.JWT([]byte(secret)))
	e.GET("/drivetrains", dc.GetAllDrivetrain())
	e.GET("/drivetrains/:name", dc.GetDrivetrainByName())
	e.DELETE("/drivetrains/:name", dc.DeletedriveTrain(), echojwt.JWT([]byte(secret)))
}

func FuelRoutes(e *echo.Echo, fc controller.FuelController, secret string) {
	e.POST("/fuels", fc.InsertFuel())
	e.GET("/fuels", fc.GetAllFuel())
	e.GET("/fuels/:name", fc.GetFuelByName())
	e.DELETE("/fuels/:name", fc.DeleteFuel())
}

func TransmissionRoutes(e *echo.Echo, tc controller.TransmissionsController, secret string) {
	e.GET("/transmissions", tc.GetAllTransmissions())
	e.POST("/transmissions", tc.InsertTransmission(), echojwt.JWT([]byte(secret)))
	e.GET("/transmissions/:name", tc.GetTransmissionByName())
	e.DELETE("/transmissions/:name", tc.DeleteTransmission(), echojwt.JWT([]byte(secret)))
}

func ManufacturersRoutes(e *echo.Echo, mc controller.ManufacturersController, secret string) {
	e.GET("/manufacturers", mc.GetAllManufacturers())
	e.GET("/manufacturers/:name", mc.GetManufacturersByName())
	e.POST("/manufacturers", mc.InsertManufacturer(), echojwt.JWT([]byte(secret)))
	e.DELETE("/manufacturers/:name", mc.DeleteManufacturer(), echojwt.JWT([]byte(secret)))
}

func TypesRoutes(e *echo.Echo, tc controller.TypesController, secret string) {
	e.GET("/types", tc.GetAllTypes())
	e.GET("/types/:name", tc.GetTypeByName())
	e.POST("/types", tc.InsertTypes(), echojwt.JWT([]byte(secret)))
	e.DELETE("/types/:name", tc.DeleteTypes(), echojwt.JWT([]byte(secret)))
}

func CarRoutes(e *echo.Echo, cc controller.CarsController, secret string) {
	e.GET("/cars", cc.GetAllCars())
	e.GET("/cars/:name", cc.GetCarByName())
	e.POST("/cars", cc.InsertCars(), echojwt.JWT([]byte(secret)))
	e.DELETE("/cars/:name", cc.DeleteCar(), echojwt.JWT([]byte(secret)))
}

func EngineRoutes(e *echo.Echo, ec controller.EngineController, secret string) {
	e.GET("/engines", ec.GetAllEngines())
	e.GET("/engines/:id", ec.GetEngineByID())
	e.POST("/engines", ec.InsertEngine(), echojwt.JWT([]byte(secret)))
	e.DELETE("engines/:id", ec.DeleteEngine(), echojwt.JWT([]byte(secret)))
}
