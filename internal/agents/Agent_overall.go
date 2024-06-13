package agents

type Agent_overall interface {
	GetType() string
	Resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int) (float64, float64, float64)
	//possible actions?
	// CheckSuccessfulAttack(attackrate float64, target_defence_rate float64) bool
	Greedy_resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int, num_moderate_agent int, greedy_mode_switch bool) (float64, float64, float64)
}

const moderate_max_consumption = 0.2  // any need for random?
const predator_max_consumption = 0.05 // any need for random?
const predator_min_consumption = 0.03 // any need for random?

// const predator_max_consumption = 0.1  // any need for random?
// const predator_min_consumption = 0.08 // any need for random?

const farmer_max_consumption = 0.4 // any need for random?
