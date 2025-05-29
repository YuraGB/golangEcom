package models

type OrderProductInput struct {
	ProductID int     `json:"productId" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required,min=1"`
	Price     float64 `json:"price" validate:"required,gt=0"`
}

type Order struct {
	PaymentType   string              `json:"paymentType" validate:"required,oneof=CASH CARD"`
	Address       string              `json:"address" validate:"required"`
	City          string              `json:"city" validate:"required"`
	State         string              `json:"state" validate:"required"`
	Zip           string              `json:"zip" validate:"required"`
	UserID        int                 `json:"userId" validate:"required"`
	OrderProducts []OrderProductInput `json:"orderProducts" validate:"required,dive"`
}

type OrderInput struct {
	Order
	TotalPrice float64 `json:"totalPrice" validate:"required,gt=0"`
	Status     string  `json:"status" validate:"required,oneof=NEW IN_PROGRESS COMPLETED CANCELED"`
}

type OrderRequestInput struct {
	UserId        int                 `json:"userId" validate:"required"`
	PaymentMethod string              `json:"paymentType" validate:"required,oneof=CASH CARD"`
	Address       string              `json:"address" validate:"required"`
	City          string              `json:"city" validate:"required"`
	State         string              `json:"state" validate:"required"`
	Zip           string              `json:"zip" validate:"required"`
	OrderProducts []OrderProductInput `json:"orderProducts" validate:"required,dive"`
}
