package main

import(
	"fmt"
	"net/http"	
	"github.com/labstack/echo"
	"github.com/spf13/viper"	
	"gopkg.in/mgo.v2"


	_userHandler "qlass-native/user"
	_userUcase "qlass-native/user/usecase"		
	_userRepo "qlass-native/user/repository"	
	_userHttpDeliver "qlass-native/user/delivery/http"	
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
	dbConn, err := mgo.Dial("localhost")

	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	e := echo.New()	
	http.handleFunc("/",Coba)

	dataUser := _userRepo.NewMongoUserRepository(dbConn)	
	au := _userUcase.NewUserUsecase(dataUser)	
	_userHttpDeliver.NewUserHandler(e, au)	

	fmt.Println("server started at PORT"+viper.GetString("server.address"))
	http.ListenAndServe(viper.GetString("server.address"), nil)
}
