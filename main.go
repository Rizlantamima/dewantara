package main

import(
	"fmt"
	"net/http"	
	"github.com/labstack/echo"
	"github.com/spf13/viper"	
	// "gopkg.in/mgo.v2"


	// _userHandler "qlass-native/user"
	// _userUcase "qlass-native/user/usecase"		
	// _userRepo "qlass-native/user/repository"	
	// _userHttpDeliver "qlass-native/user/delivery/http"	
)


func init() {
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}


func main(){
	// dbConn, err := mgo.Dial("localhost")

	// if err != nil && viper.GetBool("debug") {
		// 	fmt.Println(err)
		// }

		e := echo.New()	
		e.GET("/",func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})	
		e.GET("/coba",handleIndex)
		e.POST("/coba",handleStore)	

		// dataUser := _userRepo.NewMongoUserRepository(dbConn)	
		// au := _userUcase.NewUserUsecase(dataUser)	
		// _userHttpDeliver.NewUserHandler(e, au)	

		var server_posrt string = viper.GetString("server.address");

		fmt.Println("server started at PORT"+server_posrt)
		e.Logger.Fatal(e.Start(server_posrt))
	}


	func handleIndex(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}

	func handleStore(c echo.Context) error {
		u := &user{}
		if err := c.Bind(u); err != nil {
			return err
		}

		fmt.Println(u)
		return c.String(http.StatusOK, "Hello, World!")
	}

	func handleUpdate(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}

	func handleDelete(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}


	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
