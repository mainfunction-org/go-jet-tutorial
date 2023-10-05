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

	newPet := model.Pets{
		Name:   "Aura",
		Status: model.PetStatus_Pending,
	}

	res, err := table.Pets.
		INSERT(table.Pets.MutableColumns).
		MODEL(newPet).
		Exec(db)

	if err != nil {
		panic(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(affected)
}
