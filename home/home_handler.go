package home

import(
	"net/http"	
	"encoding/json"

	Response "qlass-native/common"		

)


func HttpHandler(){
	http.HandleFunc("/", handlerIndex)
}
func handlerIndex(res http.ResponseWriter, req *http.Request){

	var result, err = json.Marshal(Response.Error{Message: "lalaal"})

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(result)
	return
}


func insert() {
    var session, err = connect()
    if err != nil {
        fmt.Println("Error!", err.Error())
        return
    }
    defer session.Close()

    var collection = session.DB("belajar_golang").C("student")
    err = collection.Insert(&student{"Wick", 2}, &student{"Ethan", 2})
    if err != nil {
        fmt.Println("Error!", err.Error())
        return
    }

    fmt.Println("Insert success!")
}