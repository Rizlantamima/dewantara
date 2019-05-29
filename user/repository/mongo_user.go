package repository

import(
	"fmt"

	"github.com/spf13/viper"	
	"gopkg.in/mgo.v2"

)

type mongoUserRepository struct {
	Conn *sql.DB
}


// NewMongoUserRepository will create an object that represent the user.Repository interface
func NewMongoUserRepository(Conn *mgo.Session) user.Repository {
	return &mongoUserRepository{Conn}
}