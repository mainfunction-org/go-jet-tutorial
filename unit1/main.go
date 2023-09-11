package main

import (
	"fmt"
	"github.com/mainfunction-org/go-jet-tutorial/config"
)

func main() {
	db := config.MustConnectToDB()
	defer db.Close()

	fmt.Println(db.Ping())
}
