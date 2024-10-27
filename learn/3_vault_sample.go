package learn

import (
	"log"
	"proto_learn/protogen/user"
	"proto_learn/protogen/vault"
)

func Example3() {
	vaultid := 1
	vault1 := vault.Vault{
		Id:          uint32(vaultid),
		Name:        "My Login Password",
		Description: "This is my vault",
		Secret:      "adminpassword123",
		Url:         "https://localhost",
	}
	user1 := user.User{
		Id:       1,
		Username: "user1",
		Email:    "admin@root.com",
		Password: "mypassword",
		IsVip:    true,
		Vaults:   []*vault.Vault{&vault1},
	}
	log.Println(&user1)
}
