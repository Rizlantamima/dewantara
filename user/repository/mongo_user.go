package repository

import(
	"fmt"

	// "github.com/spf13/viper"	
	"gopkg.in/mgo.v2"

	"qlass-native/models"
)

type mongoUserRepository struct {
	Conn *mgo.Session
}


// NewMongoUserRepository will create an object that represent the user.Repository interface
func NewMongoUserRepository(Conn *mgo.Session) user.Repository {
	return &mongoUserRepository{Conn}
}


func (m *mongoUserRepository) Store(data *models.User) error {
	defer m.Conn.Close()
	collection := m.Conn.DB("belajar_golang").C("student")

	err = collection.Insert(&data{Name:"hoho"})
    if err != nil {
        fmt.Println("Error!", err.Error())
        return
    }

    fmt.Println("Insert success!")
    return nil
}