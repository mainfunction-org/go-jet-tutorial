package main

import (
	"fmt"

	"github.com/mainfunction-org/go-jet-tutorial/config"
	"github.com/mainfunction-org/go-jet-tutorial/internal/generated/pet_store/public/model"
	"github.com/mainfunction-org/go-jet-tutorial/internal/generated/pet_store/public/table"
)

func main() {
	db := config.MustConnectToDB()
	defer db.Close()

	var petItems []model.Pets
	statement := table.Pets.SELECT(table.Pets.AllColumns)
	err := statement.Query(db, &petItems)
	if err != nil {
		panic(err.Error())
	}

	for _, item := range petItems {
		fmt.Printf("%#v \n", item)
	}
}
