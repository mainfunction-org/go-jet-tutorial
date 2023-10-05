package main

import (
	"github.com/mainfunction-org/go-jet-tutorial/config"
)

func main() {
	db := config.MustConnectToDB()
	defer db.Close()

	// put your codmain.goe here
}
