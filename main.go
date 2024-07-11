package main

import (
	"SOSAMAS_FYP/internal/server"
	"fmt"
)

func main() {
	// Parameters
	numTribes := 10    // Number of tribes
	tribeSize := 5     // Size of each tribe
	porofGreedy := 0.2 // Proportion of greedy members
	mapSize := 5       // Size of the map
	// bandit_original_attack_rate := 0.4 // Bandit original attack rate
	bandit_original_attack_rate := 0.6 // Enhance Bandit original attack rate
	origional_attack_rate := 0.2
	origional_defence_rate := 0.2
	// rate_of_energy_taken := 0.8
	rate_of_energy_taken := 0.9 // Enhance rate of energy taken
	origional_death_rate := 0.1
	origional_birth_rate := 0.1

	// Initialization
	s := server.Initialize(numTribes, tribeSize, porofGreedy, mapSize, bandit_original_attack_rate, origional_attack_rate, origional_defence_rate, rate_of_energy_taken)
	if err := s.ExportAreaStates(); err != nil {
		fmt.Println("Failed to export area states:", err)
		return
	}
	s.RunSimulation(100, true, origional_defence_rate, origional_death_rate, origional_birth_rate) // Adjust the 100 to however many iterations you need
}
