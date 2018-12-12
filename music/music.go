package music

type Song struct {
	Id      string `bson:"_id"`
	Title    string `bson:"title"`
	Artist string `bson:"artist"`
}