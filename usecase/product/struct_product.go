package structProduct

import "time"

type InputCreateProduct struct {
	Id_Product string `json:"id_product" validate:"required"`
	Name       string `json:"name" validate:"required" unique:"name"`
	Price      int    `json:"price" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required"`
}

type InputDeleteProduct struct {
	Id string `validate:"required,uuid"`
}

type InputUpdateProduct struct {
	Id        string    `json:"id" validate:"required,uuid"`
	Name      string    `json:"name" validate:"required"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}
