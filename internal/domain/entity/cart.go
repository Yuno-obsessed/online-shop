package entity

import (
	"github.com/google/uuid"
)

type Carts struct {
	UUID        uuid.UUID
	UserUUID    uuid.UUID
	ProductUUID uuid.UUID
	Quantity    int
}
