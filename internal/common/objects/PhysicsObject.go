package objects

import (
	utils "SOSAMAS_FYP/internal/common/utils"

	"github.com/google/uuid"
)

type PhysicsObjectInterface interface {
	GetId() uuid.UUID
	GetPosition() utils.Coordinates
}

type PhysicsObject struct {
	id          uuid.UUID
	coordinates utils.Coordinates
}

func (po *PhysicsObject) GetId() uuid.UUID {
	return po.id

}

func (po *PhysicsObject) GetPosition() utils.Coordinates {
	return po.coordinates
}

// func GetPhysicsObject() *PhysicsObject {
// 	return &PhysicsObject{
// 		id:          uuid.New(),
// 		coordinates: utils.GenerateRandomCoordinates(),
// 		amount:      amount,
// 	}
// }
