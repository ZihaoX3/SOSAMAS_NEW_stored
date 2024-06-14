package server

import (
	"SOSAMAS_FYP/internal/common/objects"
	"SOSAMAS_FYP/internal/common/utils"
	"SOSAMAS_FYP/internal/maps"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Server struct {
	Tribes         []*objects.Tribe
	Areas          []*maps.Area
	PreviousStates map[string]TribeState
}

type TribeState struct {
	ID                 string   `json:"id"`
	AreaID             []int    `json:"area_id"`
	Alive_Members      int      `json:"members"`
	FoodLevel          float64  `json:"food_level"`
	WoodLevel          float64  `json:"wood_level"`
	WaterLevel         float64  `json:"water_level"`
	Role               []string `json:"role"`
	IsAttacking        bool     `json:"is_attacking"`
	Attacksuccesscheck bool     `json:"attack_success_check,omitempty"`
	TargetID           string   `json:"target_id,omitempty"`
	Haveboat           bool     `json:"has_boat"`
	ShelterLevel       int      `json:"shelter_level"`
	Relocated          bool     `json:"relocated,omitempty"`
	Previous_location  int      `json:"previous_location,omitempty"`
	BoatBuildingTime   int      `json:"boat_building_time,omitempty"`
	ShelterBuilding    bool     `json:"shelter_building,omitempty"`
}

type SimulationState struct {
	Iteration int          `json:"iteration"`
	Tribes    []TribeState `json:"tribes"`
	Agents    []AgentState `json:"agents"`
}

type AreaState struct {
	ID          int     `json:"id"`
	Environment string  `json:"environment"`
	Neighbors   []int   `json:"neighbors"`
	WaterLevel  float64 `json:"water_level"`
	FoodLevel   float64 `json:"food_level"`
	WoodLevel   float64 `json:"wood_level"`
}

type AgentState struct {
	ID           string  `json:"id"`
	TribeID      string  `json:"tribe_id"`
	Role         string  `json:"role"`
	WoodLevel    float64 `json:"wood_level"`
	FoodLevel    float64 `json:"food_level"`
	WaterLevel   float64 `json:"water_level"`
	CurrentArea  int     `json:"current_area"`
	HaveBoat     bool    `json:"has_boat"`
	ShelterLevel int     `json:"shelter_level"`
}

type EnvironmentState struct {
	AreaID     int     `json:"area_id"`
	WaterLevel float64 `json:"water_level"`
	FoodLevel  float64 `json:"food_level"`
	WoodLevel  float64 `json:"wood_level"`
}

type RoundEnvironmentState struct {
	Round             int                `json:"round"`
	EnvironmentStates []EnvironmentState `json:"environments"`
}

func (s *Server) ExportEnvironmentStates(round int) error {
	var roundState RoundEnvironmentState
	roundState.Round = round
	roundState.EnvironmentStates = make([]EnvironmentState, len(s.Areas))

	for i, area := range s.Areas {
		roundState.EnvironmentStates[i] = EnvironmentState{
			AreaID:     area.ID,
			WaterLevel: area.Environment.GetWaterLevel(),
			FoodLevel:  area.Environment.GetFoodLevel(),
			WoodLevel:  area.Environment.GetWoodLevel(),
		}
	}

	// Append to an existing file or create a new one if it doesn't exist
	filePath := "environment_states.json"
	var file *os.File
	var err error
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
	} else {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(roundState); err != nil {
		return fmt.Errorf("failed to encode environment states: %w", err)
	}

	return nil
}

func Initialize(numTribes int, tribeSize int, porofgreedy float64, map_size int, bandit_original_attack_rate float64, origional_attack_rate float64, origional_defence_rate float64, rate_of_energy_taken float64) *Server {
	s := &Server{
		Tribes:         make([]*objects.Tribe, numTribes),
		Areas:          maps.InitializeAreas(map_size),
		PreviousStates: make(map[string]TribeState),
	}

	for i := range s.Tribes {
		// s.Tribes[i].porofgreedy = porofgreedy
		s.Tribes[i] = objects.NewTribe(tribeSize, porofgreedy, bandit_original_attack_rate, origional_attack_rate, origional_defence_rate, rate_of_energy_taken)

		randomAreaIndex := rand.Intn(len(s.Areas))
		// Assign each tribe member to the same area
		for _, member := range s.Tribes[i].Memebers {
			member.CurrentArea = s.Areas[randomAreaIndex]
			member.Environment = member.CurrentArea.Environment
		}
	}
	return s
}

func (s *Server) RunSimulation(iterations int, greedySwitch bool, origional_defence_rate float64, origional_death_rate float64, origional_birth_rate float64) {
	var allStates []SimulationState

	seasonCounter := 0
	currentSeason := utils.Spring
	for _, tribe := range s.Tribes {
		tribe.InitializeTribe(origional_death_rate, origional_birth_rate)
	}

	for i := 0; i < iterations; i++ {
		currentState := SimulationState{
			Iteration: i,
			Tribes:    []TribeState{},
			Agents:    []AgentState{},
		}

		if i > 0 && i%3 == 0 {
			seasonCounter++
			if seasonCounter >= 4 {
				seasonCounter = 0
			}
			currentSeason = utils.SeasonType(seasonCounter)
		}

		for _, tribe := range s.Tribes {
			tribe.SimulateCycle(currentSeason.String(), greedySwitch, s.Tribes, origional_defence_rate, origional_death_rate)
		}

		for _, tribe := range s.Tribes {
			if tribe.Getnumberofalive() > 0 {
				avgWater, avgFood, avgWood := tribe.CalculateAverageResources()
				areaIDs := make([]int, 0)
				for _, member := range tribe.GetAliveMembers() {
					if !utils.Contains(areaIDs, member.CurrentArea.ID) {
						areaIDs = append(areaIDs, member.CurrentArea.ID)
					}
				}

				unique_roles_list := make([]string, 0)
				for _, member := range tribe.GetAliveMembers() {
					if !utils.ContainsString(unique_roles_list, member.GetAgentRole()) {
						unique_roles_list = append(unique_roles_list, member.GetAgentRole())
					}
				}

				tribeState := TribeState{
					ID:                 strconv.Itoa(tribe.TribeID),
					AreaID:             areaIDs,
					Alive_Members:      int(tribe.Getnumberofalive()),
					FoodLevel:          avgFood,
					WoodLevel:          avgWood,
					WaterLevel:         avgWater,
					Role:               unique_roles_list,
					IsAttacking:        tribe.Isattacking,
					Attacksuccesscheck: tribe.Lastattacksuccess,
					TargetID:           strconv.Itoa(tribe.TargetID),
					Haveboat:           tribe.Haveboat,
					ShelterLevel:       tribe.Shettle_level,
				}

				// Check for relocation
				prevState, exists := s.PreviousStates[tribeState.ID]
				if exists {
					if !equalAreaIDs(prevState.AreaID, tribeState.AreaID) {
						tribeState.Relocated = true
						tribeState.Previous_location = prevState.AreaID[0]
					}

					// Check for boat building time
					if !prevState.Haveboat && tribeState.Haveboat {
						tribeState.BoatBuildingTime = i
					}

					// Check for shelter building
					if prevState.ShelterLevel < tribeState.ShelterLevel {
						tribeState.ShelterBuilding = true
					}
				}

				// Update previous state
				s.PreviousStates[tribeState.ID] = tribeState

				currentState.Tribes = append(currentState.Tribes, tribeState)

				for _, member := range tribe.Memebers {
					haveboat := tribe.Haveboat
					shelterlevel := tribe.Shettle_level
					if member.IsAlive() {
						agentState := AgentState{
							ID:           member.AgentId.String(),
							TribeID:      strconv.Itoa(tribe.TribeID),
							Role:         member.GetAgentRole(),
							WoodLevel:    member.GetWoodLevel(),
							FoodLevel:    member.GetFoodLevel(),
							WaterLevel:   member.GetWaterLevel(),
							CurrentArea:  member.CurrentArea.ID,
							HaveBoat:     haveboat,
							ShelterLevel: shelterlevel,
						}
						currentState.Agents = append(currentState.Agents, agentState)
					}
				}
			}
		}
		allStates = append(allStates, currentState)
	}
	jsonData, err := json.MarshalIndent(allStates, "", "    ")
	if err != nil {
		fmt.Println("Error serializing simulation states:", err)
		return
	}

	if err := os.WriteFile("simulation_data.json", jsonData, 0644); err != nil {
		fmt.Println("Error writing simulation states to file:", err)
	}

	fmt.Println("Simulation ended")
}

func equalAreaIDs(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (s *Server) ExportAreaStates() error {
	areaStates := make([]AreaState, len(s.Areas))
	for i, area := range s.Areas {
		neighbors := make([]int, len(area.Neighbors))
		for j, neighbor := range area.Neighbors {
			neighbors[j] = neighbor.ID
		}
		areaStates[i] = AreaState{
			ID:          area.ID,
			Environment: area.Environment.GetType(),
			Neighbors:   neighbors,
			WaterLevel:  area.Environment.GetWaterLevel(),
			FoodLevel:   area.Environment.GetFoodLevel(),
			WoodLevel:   area.Environment.GetWoodLevel(),
		}
	}

	jsonData, err := json.MarshalIndent(areaStates, "", "    ")
	if err != nil {
		return fmt.Errorf("error serializing area states: %w", err)
	}
	if err := os.WriteFile("area_states.json", jsonData, 0644); err != nil {
		return fmt.Errorf("error writing area states to file: %w", err)
	}
	return nil
}
