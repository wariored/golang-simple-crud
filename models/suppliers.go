package models

import (
	"errors"
	"supplier/db"

	"github.com/patrickmn/go-cache"
)

var ErrRecordNotFound = errors.New("record not found")

type Supplier struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	LogoURL string `json:"logoURL"`
}

type SupplierUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	LogoURL string `json:"logoURL"`
}

func GetSupplierByID(supplierID string) (Supplier, error) {
	if value, found := db.Database.Get(supplierID); found {
		return value.(Supplier), nil
	}
	return Supplier{}, ErrRecordNotFound
}

func (supplier *Supplier) Update(reqObj SupplierUpdateRequest) Supplier {
	newObj := Supplier{
		ID:      supplier.ID,
		Name:    reqObj.Name,
		Address: reqObj.Address,
		LogoURL: reqObj.LogoURL,
	}
	// should have checked the error and return 2 values (object, error)
	// in the case the id doesn't exists
	db.Database.Set(supplier.ID, newObj, cache.NoExpiration)
	
	return newObj
}

func (supplier *Supplier) Delete() int {
	// should have checked the error
	// decision made by following the current DB used
	db.Database.Delete(supplier.ID)
	// the deleted count is usually returned by the database delete function
	deleted_count := 1
	return deleted_count
}

func CreateSupplier(supplier Supplier) {
	db.Database.Set(supplier.ID, supplier, cache.NoExpiration)
}
