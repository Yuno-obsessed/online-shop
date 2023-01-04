package entity

import "github.com/google/uuid"

type Liked struct {
	UUID        uuid.UUID
	UserUUID    uuid.UUID
	ProductUUID uuid.UUID
}
