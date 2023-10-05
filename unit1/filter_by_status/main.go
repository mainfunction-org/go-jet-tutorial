package main

import (
	"fmt"
	"github.com/mainfunction-org/go-jet-tutorial/internal/generated/pet_store/public/enum"

	"github.com/mainfunction-org/go-jet-tutorial/config"
	"github.com/mainfunction-org/go-jet-tutorial/internal/generated/pet_store/public/model"
	"github.com/mainfunction-org/go-jet-tutorial/internal/generated/pet_store/public/table"
)

func main() {
	db := config.MustConnectToDB()
	defer db.Close()

	var availablePets []model.Pets
	statement := table.Pets.
		SELECT(table.Pets.AllColumns).
		WHERE(table.Pets.Status.EQ(enum.PetStatus.Available))

	err := statement.Query(db, &availablePets)
	if err != nil {
		panic(err.Error())
	}

	for _, item := range availablePets {
		fmt.Printf("%#v \n", item)
	}
}
