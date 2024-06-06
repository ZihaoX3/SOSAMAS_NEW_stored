package agents

type Farmer_agent struct {
	waterconsumption float64
	foodconsumption  float64
	woodconsumption  float64
	// attacksuccessrate float64
	// defendsuccessrate float64
}

func (g *Farmer_agent) GetType() string {
	return "Farmer"
}

func (g *Farmer_agent) Resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int) (float64, float64, float64) {

	// //take everything at once
	// g.waterconsumption = available_water / float64(num_alive_agent)
	// g.foodconsumption = available_food / float64(num_alive_agent)
	// g.woodconsumption = available_wood / float64(num_alive_agent)

	if available_water > float64(num_alive_agent)*moderate_max_consumption {
		g.waterconsumption = farmer_max_consumption
	} else {
		g.waterconsumption = available_water / float64(num_alive_agent)
	}

	if available_food > float64(num_alive_agent)*moderate_max_consumption {
		g.foodconsumption = farmer_max_consumption
	} else {
		g.foodconsumption = available_food / float64(num_alive_agent)
	}

	if available_wood > float64(num_alive_agent)*moderate_max_consumption {
		g.woodconsumption = farmer_max_consumption
	} else {
		g.woodconsumption = available_wood / float64(num_alive_agent)
	}

	return g.waterconsumption, g.foodconsumption, g.woodconsumption
}

func (g *Farmer_agent) Greedy_resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int, num_moderate_agent int, greedy_mode_switch bool) (float64, float64, float64) {
	//take everything at once
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
			g.waterconsumption = averagetaken_water
		} else {
			g.waterconsumption = baselinewater
		}
		if averagetaken_food > baselinefood {
			g.foodconsumption = averagetaken_food
		} else {
			g.foodconsumption = baselinefood
		}
		if averagetaken_wood > baselinewood {
			g.woodconsumption = averagetaken_wood
		} else {
			g.woodconsumption = baselinewood
		}
	} else {
		if extremwater > baselinewater {
			g.waterconsumption = extremwater
		} else {
			g.waterconsumption = baselinewater
		}

		if extremfood > baselinefood {
			g.foodconsumption = extremfood
		} else {
			g.foodconsumption = baselinefood
		}

		if extremwood > baselinewood {
			g.woodconsumption = extremwood
		} else {
			g.woodconsumption = baselinewood
		}
	}
	return g.waterconsumption, g.foodconsumption, g.woodconsumption
}

// func (g *Farmer_agent) CheckSuccessfulAttack(attackrate float64, target_defence_rate float64) bool {
// 	return false
// }
