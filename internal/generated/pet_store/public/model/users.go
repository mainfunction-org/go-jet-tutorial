//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Users struct {
	ID        int64 `sql:"primary_key"`
	Username  string
	FirstName *string
	LastName  *string
	Email     string
	Password  string
	Phone     *string
	Status    *UserStatus
}
