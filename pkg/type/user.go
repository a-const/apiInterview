package user

type User struct {
	Username    string `bson:"username"`
	Password    string `bson:"password"`
	Description string `bson:"description"`
}
