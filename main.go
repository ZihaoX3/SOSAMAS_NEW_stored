package main

import (
	"SOSAMAS_FYP/internal/server"
	"fmt"
)

func main() {
	// Parameters for initialization
	numTribes := 10                    // Example: Number of tribes
	tribeSize := 5                     // Example: Size of each tribe
	porofGreedy := 1.0                 // Example: Proportion of greedy members
	mapSize := 5                       // Example: Size of the map
	bandit_original_attack_rate := 0.4 // Example: Bandit original attack rate
	origional_attack_rate := 0.2
	origional_defence_rate := 0.2
	rate_of_energy_taken := 0.8
	origional_death_rate := 0.1
	origional_birth_rate := 0.1

	// Initialize the server with the specified parameters
	//(numTribes int, tribeSize int, porofgreedy float64, map_size int, bandit_original_attack_rate float64, origional_attack_rate float64, origional_defence_rate float64, rate_of_energy_taken float64)
	s := server.Initialize(numTribes, tribeSize, porofGreedy, mapSize, bandit_original_attack_rate, origional_attack_rate, origional_defence_rate, rate_of_energy_taken)

	if err := s.ExportAreaStates(); err != nil {
		fmt.Println("Failed to export area states:", err)
		return
	}

	// Run the simulation
	// Parameters: number of iterations, and whether the greedy switch is enabled
	//RunSimulation(iterations int, greedySwitch bool, origional_defence_rate float64, origional_death_rate float64, origional_birth_rate float64)
	s.RunSimulation(100, false, origional_defence_rate, origional_death_rate, origional_birth_rate) // Adjust the 100 to however many iterations you need
}
