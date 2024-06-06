package objects

import (
	// "SOSAMAS_FYP/internal/common/utils"

	"SOSAMAS_FYP/internal/agents"
	"SOSAMAS_FYP/internal/common/utils"
	"SOSAMAS_FYP/internal/environment"
	"SOSAMAS_FYP/internal/maps"
	"fmt"
	"math/rand"

	baseAgent "github.com/MattSScott/basePlatformSOMAS/BaseAgent"
	"github.com/google/uuid"
)

type IBaseTribe interface {
	baseAgent.IAgent[IBaseTribe]
	DecideAction()

	GetEnergy_food() float64
	UpdateEnergy_food(food_level float64)
	GetEnergy_water() float64
	UpdateEnergy_water(water_level float64)
	GetEnergy_wood() float64
	UpdateEnergy_wood(wood_level float64)
	// GetGreediness() utils.Greediness
	// GetAgentRole() string
	// GetType() string

	UpdateGameState(gameState IGameState)
}

type TribeAction int

const (
	Stay TribeAction = iota
	Move
)

type BaseTribe struct {
	*baseAgent.BaseAgent[IBaseTribe]
	AgentType           utils.AgentType
	food_level          float64
	water_level         float64
	wood_level          float64
	AgentId             uuid.UUID
	GroupID             int
	gameState           IGameState
	deadflag            bool
	Environment         environment.Landscape
	waterused_laststage float64
	foodused_laststage  float64
	CurrentArea         *maps.Area
	Agentbehaviour      agents.Agent_overall
	Greediness          utils.Greediness
}

// func (t *BaseTribe) DecideAction() TribeAction {
// 	return Stay
// }

// func (t *BaseTribe) DecideAllocation()

// initialise
func GetIBaseTribe(agentType utils.AgentType) *BaseTribe {
	var behavior agents.Agent_overall
	switch agentType {
	case utils.Forager_agent:
		behavior = &agents.Forager_agent{} // Assuming this is a concrete type implementing Agent_overall
	case utils.Farmer_agent:
		behavior = &agents.Farmer_agent{} // Similarly, for GreedyAgent
	case utils.Bandit_Agent:
		behavior = &agents.Bandit_Agent{} // And for PredatorAgent
	// Add cases for other types as necessary
	default:
		// Handle default case or error
	}
	return &BaseTribe{
		BaseAgent:           baseAgent.NewBaseAgent[IBaseTribe](),
		food_level:          utils.GenerateRandomFloat(1.0, 1.5),
		water_level:         utils.GenerateRandomFloat(1.0, 1.5),
		wood_level:          0.0,
		GroupID:             0,
		waterused_laststage: 0,
		foodused_laststage:  0,
		deadflag:            false,
		Environment:         nil,
		CurrentArea:         nil,
		AgentType:           agentType,
		Agentbehaviour:      behavior,
		Greediness:          utils.Moderate,
	}
}

func (t *BaseTribe) UpdateGameState(gameState IGameState) {
	t.gameState = gameState
}

func (t *BaseTribe) IsAlive() bool {
	return ((t.food_level > 0 && t.water_level > 0) && (!t.deadflag))
}

func (t *BaseTribe) AssignRandomEnvironment() {
	environments := map[utils.EnvironmentType]environment.Landscape{
		utils.Forest:  &environment.Forest{}, // Ensure these are non-nil and properly initialized
		utils.Desert:  &environment.Desert{}, // Same here
		utils.Pasture: &environment.Pasture{},
		// other environments...
	}

	randomEnvType := utils.EnvironmentType(rand.Intn(len(environments)))
	t.Environment = environments[randomEnvType] // Assigns the chosen environment

	if t.Environment != nil {
		t.Environment.InitializeResources() // Initializes resources of the environment
	} else {
		fmt.Println("Failed to assign an environment")
	}
}

func (t *BaseTribe) Keepenvenergyupdate(season string) (float64, float64, float64) {
	//use this at the beginning of the round
	if t.Environment != nil {
		//total resource before the increase this round
		//assume the human action first
		t.Environment.UpdateResources(season)
		waterGain, foodGain, woodGain := t.Environment.GetResources()
		return waterGain, foodGain, woodGain
	} else {
		fmt.Println("Environment is nil")
		waterGain, foodGain, woodGain := 0.0, 0.0, 0.0
		return waterGain, foodGain, woodGain
	}

}

// used in sever to update agent resources every round
func (t *BaseTribe) UpdateResources(season string, numofalive int, numberofmoderate int, water_thisround_all float64, food_thisround_all float64, wood_thisround_all float64, greedy_switch bool) {
	// check of death
	// if t.water_level < 0 || t.food_level < 0 {
	// 	t.deadflag = true
	// 	fmt.Printf("Agent ID: %v, is dead due to low energy\n", t.AgentId)
	// }
	// Decrease water and food level for agents

	//keep the final energy set at when agent dead
	if !t.deadflag {
		t.waterused_laststage = utils.GenerateRandomFloat(0.1, 0.15)
		t.foodused_laststage = utils.GenerateRandomFloat(0.1, 0.15)

		t.water_level -= t.waterused_laststage
		t.food_level -= t.foodused_laststage
		fmt.Printf("WaterUsed: %.2f, FoodUsed: %.2f", t.waterused_laststage, t.foodused_laststage)
		// Check if the agent is already dead

		// Add resources from the environment
		if t.Environment != nil {
			// //total resource before the increase this round
			// //assume the human action first
			// t.Environment.UpdateResources(season)
			// // this is rotated
			// waterGain, foodGain, woodGain := t.Environment.GetResources()

			// Determine actual amount to be consumed

			// actualWaterConsumption := min(waterGain, 0.1)
			moderate_water_consumption, moderate_food_consumption, moderate_wood_consumption := t.Agentbehaviour.Resource_takenfromenv(water_thisround_all, food_thisround_all, wood_thisround_all, numofalive)
			greedy_water_consumption, greedy_food_consumption, greedy_wood_consumption := t.Agentbehaviour.Greedy_resource_takenfromenv(water_thisround_all, food_thisround_all, wood_thisround_all, numofalive, numberofmoderate, greedy_switch)
			if t.Greediness == utils.Moderate {
				actualWaterConsumption, actualFoodConsumption, actualWoodConsumption := t.Environment.ExtractResources(moderate_water_consumption, moderate_food_consumption, moderate_wood_consumption)
				fmt.Printf("Moderate consuming")
				fmt.Printf("WaterGain: %.2f, FoodGain: %.2f, WoodGain: %.2f\n", actualWaterConsumption, actualFoodConsumption, actualWoodConsumption)
				// Update agent's resources
				t.water_level += actualWaterConsumption
				t.food_level += actualFoodConsumption
				t.wood_level += actualWoodConsumption
			} else {
				actualWaterConsumption, actualFoodConsumption, actualWoodConsumption := t.Environment.ExtractResources(greedy_water_consumption, greedy_food_consumption, greedy_wood_consumption)
				//this keep same
				fmt.Printf("Greedy consuming")
				fmt.Printf("WaterGain: %.2f, FoodGain: %.2f, WoodGain: %.2f\n", actualWaterConsumption, actualFoodConsumption, actualWoodConsumption)
				t.water_level += actualWaterConsumption
				t.food_level += actualFoodConsumption
				t.wood_level += actualWoodConsumption
			}
		}

		fmt.Printf("Agent ID: %v, Water Level: %.2f, Food Level: %.2f, Wood Level: %.2f\n", t.AgentId, t.water_level, t.food_level, t.wood_level)
	}

}

// // Helper functions
// func max(a, b float64) float64 {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b float64) float64 {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func (t *BaseTribe) GetWaterLevel() float64 {
	return t.water_level
}

func (t *BaseTribe) GetFoodLevel() float64 {
	return t.food_level
}

func (t *BaseTribe) GetWoodLevel() float64 {
	return t.wood_level
}

func (t *BaseTribe) GetType() string {
	return t.Environment.GetType()
}

func (t *BaseTribe) GetAgentRole() string {
	return t.AgentType.String()
}

func (t *BaseTribe) GetGreediness() string {
	return t.Greediness.String()
}

func (t *BaseTribe) Roleswitch(desired_role utils.AgentType) {
	switch desired_role {
	case utils.Forager_agent:
		t.Agentbehaviour = &agents.Forager_agent{}
		t.AgentType = utils.Forager_agent
		fmt.Println("Agent role switched to Forager")
	case utils.Farmer_agent:
		t.Agentbehaviour = &agents.Farmer_agent{}
		t.AgentType = utils.Farmer_agent
		fmt.Println("Agent role switched to Farmer")
	case utils.Bandit_Agent:
		t.Agentbehaviour = &agents.Bandit_Agent{}
		t.AgentType = utils.Bandit_Agent
		fmt.Println("Agent role switched to Bandit")
	}

}

func (t *BaseTribe) Relocate(targetArea *maps.Area) {
	t.CurrentArea = targetArea
	fmt.Printf("Agent ID: %v, Relocated to Area ID %v with environment: %v\n", t.AgentId, targetArea.ID, targetArea.Environment)
}

// func (t *BaseTribe) Check_possible_relocation(targetArea *Area) bool {
// 	// check area_connected
// 	check1 := t.CurrentArea.Is_area_connected(targetArea)
// }

func (t *BaseTribe) Contribute_to_woodwork(required_wood float64) float64 {
	if t.IsAlive() {
		contribution := min(t.wood_level, required_wood)
		t.wood_level -= contribution
		return contribution
	}
	return 0
}
