package entityProduct

type InputCreateProduct struct {
	Id_Product string `json:"id_product" validate:"required"`
	Name       string `json:"name" validate:"required" unique:"name"`
	Price      int    `json:"price" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required"`
}

type InputDeleteProduct struct {
	Id string `validate:"required,uuid"`
}
