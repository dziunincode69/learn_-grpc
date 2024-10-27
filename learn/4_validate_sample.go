package learn

import (
	"log"
	"proto_learn/protogen/user"
)

func Example4() {
	user1 := user.User{
		Id:       1,
		Username: "user1",
		Email:    "admin@asd",
		Password: "mypassword123123",
		IsVip:    true,
	}

	err := user1.ValidateAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(&user1)
}
