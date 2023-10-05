package main

import (
	"fmt"
	"github.com/go-jet/jet/v2/postgres"

	"github.com/mainfunction-org/go-jet-tutorial/config"
	"github.com/mainfunction-org/go-jet-tutorial/internal/generated/pet_store/public/model"
	"github.com/mainfunction-org/go-jet-tutorial/internal/generated/pet_store/public/table"
)

func main() {
	db := config.MustConnectToDB()
	defer db.Close()

	var pet model.Pets
	statement := table.Pets.
		SELECT(table.Pets.AllColumns).
		WHERE(table.Pets.ID.EQ(postgres.Int(2)))

	err := statement.Query(db, &pet)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v", pet)
}
