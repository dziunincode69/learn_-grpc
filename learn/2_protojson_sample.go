package learn

import (
	"log"
	"proto_learn/protogen/user"

	"google.golang.org/protobuf/encoding/protojson"
)

func Example2() {
	UserAddress := user.Address{
		Street:  "street 1",
		Country: "Indonesia",
	}
	User := user.User{
		Id:       1,
		Username: "user1",
		Email:    "admin@gmail.com",
		Password: "mypassword123123",
		Gender:   user.Gender_GENDER_MALE,
		IsVip:    false,
		Hobbies:  []string{"reading", "coding"},
		Address:  &UserAddress,
	}
	m, err := ProtoToJSON(&User)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(m))

	m2, err := JSONToProto(m)
	if err != nil {
		log.Println(err)
	}
	log.Println(m2)
}

func ProtoToJSON(data *user.User) ([]byte, error) {
	m, err := protojson.Marshal(data)
	if err != nil {
		return nil, err
	}
	return m, nil
}
func JSONToProto(data []byte) (*user.User, error) {
	var User user.User
	err := protojson.Unmarshal(data, &User)
	if err != nil {
		return nil, err
	}
	return &User, nil
}
