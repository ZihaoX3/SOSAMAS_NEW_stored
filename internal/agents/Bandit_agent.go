package agents

import (
	"math/rand"
)

type Bandit_Agent struct {
	waterconsumption float64
	foodconsumption  float64
	woodconsumption  float64
	// attacksuccessrate float64
	// defendsuccessrate float64
}

func (p *Bandit_Agent) GetType() string {
	return "Bandit"
}

func (p *Bandit_Agent) Resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int) (float64, float64, float64) {
	if available_water > float64(num_alive_agent)*predator_min_consumption {
		p.waterconsumption = predator_min_consumption
	} else {
		p.waterconsumption = available_water / float64(num_alive_agent)
	}

	if available_food > float64(num_alive_agent)*predator_min_consumption {
		p.foodconsumption = predator_min_consumption
	} else {
		p.foodconsumption = available_food / float64(num_alive_agent)
	}

	if available_wood > float64(num_alive_agent)*predator_min_consumption {
		p.woodconsumption = predator_min_consumption
	} else {
		p.woodconsumption = available_wood / float64(num_alive_agent)
	}

	return p.waterconsumption, p.foodconsumption, p.woodconsumption
}

func (p *Bandit_Agent) Greedy_resource_takenfromenv(available_water, available_food, available_wood float64, num_alive_agent int, num_moderate_agent int, greedy_mode_switch bool) (float64, float64, float64) {
	if available_water > float64(num_alive_agent)*predator_max_consumption {
		p.waterconsumption = predator_max_consumption
	} else {
		p.waterconsumption = available_water / float64(num_alive_agent)
	}

	if available_food > float64(num_alive_agent)*predator_max_consumption {
		p.foodconsumption = predator_max_consumption
	} else {
		p.foodconsumption = available_food / float64(num_alive_agent)
	}

	if available_wood > float64(num_alive_agent)*predator_max_consumption {
		p.woodconsumption = predator_max_consumption
	} else {
		p.woodconsumption = available_wood / float64(num_alive_agent)
	}

	return p.waterconsumption, p.foodconsumption, p.woodconsumption
}

// func (p *Bandit_Agent) attack() {
// 	//how to attck? ---> randomly take one neighbor down?
// 	//Take all energy?

// 	// if that land is a predator  what to do? attack?
// 	// random a number, larger than 0.3 (70% chance to succeed) then take the resource---> every agents on that land
// 	available_attack_area:= p.CurrentArea.Neighbors
// }

// func (b *Bandit_Agent) Attack(targets []*BaseTribe) {
// 	successRate := rand.Float64() // Random success rate

// 	if successRate > 0.5 { // Assuming success if > 0.5
// 		targets := b.CurrentArea.GetAllTribes() // Get all potential target tribes
// 		totalEnergyStolen := 0.0

// 		for _, target := range targets {
// 			if target.Type != "Bandit" { // Avoid attacking other bandits
// 				energyStolen := target.Energy * 0.9
// 				target.Energy -= energyStolen
// 				totalEnergyStolen += energyStolen
// 			}
// 		}

// 		// Redistribute stolen energy among Bandit tribe members
// 		// Assuming a method to get all bandit members: b.GetBanditMembers()
// 		members := b.GetBanditMembers()
// 		energyPerMember := totalEnergyStolen / float64(len(members))
// 		for _, member := range members {
// 			member.Energy += energyPerMember
// 		}
// 	}
// }

func CheckSuccessfulAttack(attack_rate float64, target_defence float64) bool {
	porpotion := attack_rate / (attack_rate + target_defence)
	return (rand.Float64() < (1.2 * porpotion))
}

// func (p *Bandit_Agent) Should_Attack(neighbors []*Tribe) bool {

// }
