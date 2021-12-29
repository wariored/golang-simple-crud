package handlers

import (
	"encoding/json"
	"net/http"
	"supplier/models"
	"supplier/utils"

	"github.com/go-chi/render"
)

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	err := json.NewDecoder(r.Body).Decode(&supplier)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest)
		return
	}
	
	models.CreateSupplier(supplier)
	render.JSON(w, r, supplier)	
}
func GetSupplier(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	supplier, _ := ctx.Value("supplier").(models.Supplier)
	render.JSON(w, r, supplier)
}

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	var newSupplier models.SupplierUpdateRequest
	ctx := r.Context()
	supplier, _ := ctx.Value("supplier").(models.Supplier)
	err := json.NewDecoder(r.Body).Decode(&newSupplier)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest)
		return
	}
	
	supplier = supplier.Update(newSupplier)
	render.JSON(w, r, supplier)
}
func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	supplier, _ := ctx.Value("supplier").(models.Supplier)
	count := supplier.Delete()
	render.JSON(w, r, map[string]int {
		"deleted": count,
	})
}
