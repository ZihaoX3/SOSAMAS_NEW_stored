package agents

//energy used per round : same range 0.-0.2
// engergy foraging: 0.3

// import (
// 	"fmt"
// 	"math/rand"
// )

// we need to know the total resource available in the environments --->

//
type Forager_agent struct {
	waterconsumption float64
	foodconsumption  float64
	woodconsumption  float64
	// attacksuccessrate float64
	// defendsuccessrate float64
}

func (m *Forager_agent) GetType() string {
	return "Forager"
}

func (m *Forager_agent) Resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int) (float64, float64, float64) {
	if available_water > float64(num_alive_agent)*moderate_max_consumption {
		m.waterconsumption = moderate_max_consumption
	} else {
		m.waterconsumption = available_water / float64(num_alive_agent)
	}

	if available_food > float64(num_alive_agent)*moderate_max_consumption {
		m.foodconsumption = moderate_max_consumption
	} else {
		m.foodconsumption = available_food / float64(num_alive_agent)
	}

	if available_wood > float64(num_alive_agent)*moderate_max_consumption {
		m.woodconsumption = moderate_max_consumption
	} else {
		m.woodconsumption = available_wood / float64(num_alive_agent)
	}

	return m.waterconsumption, m.foodconsumption, m.woodconsumption
}

func (m *Forager_agent) Greedy_resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int, num_moderate_agent int, greedy_mode_switch bool) (float64, float64, float64) {
	//take everything at once
	//normal switch
	averagetaken_water := available_water / float64(num_alive_agent)
	averagetaken_food := available_food / float64(num_alive_agent)
	averagetaken_wood := available_wood / float64(num_alive_agent)
	// extreme switch
	restofwateronland := available_water - (moderate_max_consumption * float64(num_alive_agent))
	restoffoodonland := available_food - (moderate_max_consumption * float64(num_alive_agent))
	restofwoodonland := available_wood - (moderate_max_consumption * float64(num_alive_agent))
	extremwater := restofwateronland / float64(num_alive_agent-num_moderate_agent)
	extremfood := restoffoodonland / float64(num_alive_agent-num_moderate_agent)
	extremwood := restofwoodonland / float64(num_alive_agent-num_moderate_agent)
	//baseline
	baselinewater := available_water / float64(num_alive_agent)
	baselinefood := available_food / float64(num_alive_agent)
	baselinewood := available_wood / float64(num_alive_agent)

	if greedy_mode_switch {
		if averagetaken_water > baselinewater {
			m.waterconsumption = averagetaken_water

		} else {
			m.waterconsumption = baselinewater
		}

		if averagetaken_food > baselinefood {
			m.foodconsumption = averagetaken_food
		} else {
			m.foodconsumption = baselinefood
		}

		if averagetaken_wood > baselinewood {
			m.woodconsumption = averagetaken_wood
		} else {
			m.woodconsumption = baselinewood
		}
	} else {
		if extremwater > baselinewater {
			m.waterconsumption = extremwater
		} else {
			m.waterconsumption = baselinewater
		}

		if extremfood > baselinefood {
			m.foodconsumption = extremfood
		} else {
			m.foodconsumption = baselinefood
		}

		if extremwood > baselinewood {
			m.woodconsumption = extremwood
		} else {
			m.woodconsumption = baselinewood
		}

	}
	return m.waterconsumption, m.foodconsumption, m.woodconsumption
}
