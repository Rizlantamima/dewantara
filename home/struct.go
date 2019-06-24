package home


type student struct {
    Name  string `bson:"name"`
    Grade int    `bson:"Grade"`
}