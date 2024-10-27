package learn

import (
	"log"
	"proto_learn/protogen/user"
)

func Hello() {
	User := user.User{
		Id:       1,
		Username: "user1",
		Email:    "test@gmail.com",
		Password: "mypassword123123",
		Gender:   user.Gender_GENDER_FEMALE,
	}
	log.Println(&User)
}
