package objects

import "github.com/google/uuid"

type IGameState interface {
	GetAgent() map[uuid.UUID]IBaseTribe
	// GetMegaTribes() map[uuid.UUID]IMegaTribe
}
