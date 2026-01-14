package validations

type ItemRequestValidation struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,max=500"`
	Price       int    `json:"price" validate:"required,gt=0"`
	Stock       int    `json:"stock" validate:"required,gte=0"`
}
