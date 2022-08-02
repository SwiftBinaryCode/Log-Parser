//structs with json

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type permissions map[string]bool
type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {

	//decoding json

	type user struct {
		Name        string          `json:"username"`
		Permissions map[string]bool `json:"perms"`
	}

	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {

		input = append(input, in.Bytes()...)
	}

	var users []user

	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {

		case p == nil:
			fmt.Print("has no power")

		case p["admin"]:
			fmt.Print("is an admin")

		case p["write"]:
			fmt.Print("can write")
		}

		fmt.Println()

	}

	//------------
	// users := []user{
	// 	{"inanc", "1234", nil},
	// 	{"god", "42", permissions{"admin": true}},
	// 	{"devil", "666", permissions{"write": true}},
	// }

	// out, err := json.MarshalIndent(users, "", "\t")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(out))

}
