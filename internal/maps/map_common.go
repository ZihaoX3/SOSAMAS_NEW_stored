package maps

import (
	"SOSAMAS_FYP/internal/environment"
	"math/rand"
	"time"
)

// const map_size = 10
const max_connections = 2

type LandscapeType int

type Area struct {
	ID          int
	Environment environment.Landscape // Use the interface directly
	Neighbors   []*Area
}

func NewArea(id int, env environment.Landscape) *Area {
	return &Area{
		ID:          id,
		Environment: env, // Set the environment of the area
		Neighbors:   []*Area{},
	}
}

func (a *Area) AddNeighbor(neighbor *Area) {
	a.Neighbors = append(a.Neighbors, neighbor)
}

func getRandomLandscape(rnd *rand.Rand) environment.Landscape {
	var landscape environment.Landscape
	switch rnd.Intn(3) {
	case 0:
		landscape = &environment.Forest{}
	case 1:
		landscape = &environment.Pasture{}
	default:
		landscape = &environment.Desert{}
	}

	landscape.InitializeResources()
	return landscape
}

func InitializeAreas(map_size int) []*Area {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	areas := make([]*Area, map_size)

	// Initialize areas
	for i := 0; i < map_size; i++ {
		landscape := getRandomLandscape(rnd) // Assuming this function returns a random Landscape
		areas[i] = NewArea(i, landscape)

	}

	// Randomly establish connections
	for _, area := range areas {
		setRandomConnections(area, areas, rnd, map_size)
	}

	return areas
}

func setRandomConnections(area *Area, areas []*Area, rnd *rand.Rand, map_size int) {
	for len(area.Neighbors) < max_connections {
		potentialNeighbor := areas[rnd.Intn(map_size)]
		if potentialNeighbor != area && !isAlreadyConnected(area, potentialNeighbor) {
			area.AddNeighbor(potentialNeighbor)
			potentialNeighbor.AddNeighbor(area) // For bidirectional connection
		}
	}
}

func isAlreadyConnected(area1, area2 *Area) bool {
	for _, neighbor := range area1.Neighbors {
		if neighbor == area2 {
			return true
		}
	}
	return false
}

func Is_area_connected(areawein []*Area, areawecheck *Area) bool {
	for _, a := range areawein {
		for _, neighbor := range a.Neighbors {
			if neighbor == areawecheck {
				return true
			}
		}
	}
	return false
}
