package main

import(
	"fmt"
	"net/http"	

	"github.com/spf13/viper"	
	 "gopkg.in/mgo.v2"


	_userHandler "qlass-native/user"
	_userUcase "qlass-native/user/usecase"		
	_userRepo "qlass-native/user/repository"	
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


	dataUser := _userRepo.NewMongoUserRepository(dbConn)	
	au := _userUcase.NewUserUsecase(dataUser)	
	_userHandler.HttpHandler()


    fmt.Println("server started at PORT"+viper.GetString("server.address"))
    http.ListenAndServe(viper.GetString("server.address"), nil)
}

