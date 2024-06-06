package objects

import (
	"SOSAMAS_FYP/internal/common/utils"
	"SOSAMAS_FYP/internal/maps"
	"fmt"
	"math"
	"math/rand"
)

type ResourceMetrics struct {
	Water       float64
	Food        float64
	Wood        float64
	AttackCount int // Add this field to track the number of successful attacks
}

type Tribe struct {
	Memebers                     []*BaseTribe
	deathrate                    float64
	birthrate                    float64
	numofdeath                   float64
	numofbirth                   float64
	tribetype                    utils.AgentType
	porofgreedy                  float64
	attackrate                   float64
	defendrate                   float64
	energytakenrate              float64
	TribeID                      int
	totalneeds_ofwood_forboat    float64
	totalneeds_ofwood_forshettle float64
	Haveboat                     bool
	Shettle_level                int
	Isattacking                  bool
	IsAttacked                   bool
	TargetID                     int
	SatisfactionHistory          []float64
	CurrentSatisfaction          float64
	Ranking                      map[int]float64
	CurrentResources             ResourceMetrics
	PreviousResources            ResourceMetrics
	boat_making_signal           bool
	shettle_signal               bool
	attack_signal                bool
	relocate_signal              bool
	Area_last_stage              *maps.Area
	Lastattacksuccess            bool
	PotentialBirths              float64
	PotentialDeaths              float64
	SuccessfulAttackGains        map[int]ResourceMetrics
	// attackrate  float64
	// defendrate  float64
}

func RandomAgentType() utils.AgentType {
	return utils.AgentType(rand.Intn(3))
}

func NewTribe(size int, pofgreedy float64, bandit_original_attack_rate float64, origional_attack_rate float64, origional_defence_rate float64, rate_of_energy_taken float64) *Tribe {
	tribe := &Tribe{Memebers: make([]*BaseTribe, size),
		SatisfactionHistory:   make([]float64, 0),
		CurrentSatisfaction:   0,
		SuccessfulAttackGains: make(map[int]ResourceMetrics)}
	randomTribeType := utils.AgentType(rand.Intn(3))
	tribe.tribetype = randomTribeType
	tribe.porofgreedy = pofgreedy
	numofgreedy := int(pofgreedy * float64(size))
	//keep everyagent same type in one tribe
	for i := 0; i < size; i++ {
		tribe.Memebers[i] = GetIBaseTribe(randomTribeType)
		tribe.Memebers[i].Greediness = utils.Moderate
	}
	for i := 0; i < numofgreedy; i++ {
		tribe.Memebers[i].Greediness = utils.Greedy
	}
	tribe.Ranking = make(map[int]float64)

	if randomTribeType == utils.Bandit_Agent {
		tribe.attackrate = bandit_original_attack_rate
		tribe.defendrate = origional_defence_rate
		tribe.energytakenrate = rate_of_energy_taken

	} else {
		tribe.attackrate = origional_attack_rate
		tribe.defendrate = origional_defence_rate

	}
	tribe.TribeID = rand.Intn(10000)
	tribe.Lastattacksuccess = false
	return tribe
}

func (t *Tribe) InitializeTribe(origional_death_rate float64, origional_birth_rate float64) {
	t.deathrate = origional_death_rate
	t.birthrate = origional_birth_rate
	t.numofdeath = 0
	t.numofbirth = 0
	t.totalneeds_ofwood_forboat = 3
	t.totalneeds_ofwood_forshettle = 2
	t.Haveboat = false
	t.Shettle_level = 0
	t.Isattacking = false
	t.IsAttacked = false
	t.SatisfactionHistory = append(t.SatisfactionHistory, 0.0)
	t.CurrentSatisfaction = 0.0
	t.boat_making_signal = false
	t.shettle_signal = false
	t.attack_signal = false
	t.relocate_signal = false
	t.Lastattacksuccess = false
	t.PotentialBirths = 0
	t.PotentialDeaths = 0
}

func (t *Tribe) SimulateCycle(currentSeason string, greedy_switch bool, alltribes []*Tribe, origional_defence_rate float64, origional_death_rate float64) {
	t.PreviousResources = t.CurrentResources
	t.Lastattacksuccess = false
	// Update resources for each member
	water_this_round, food_this_round, wood_this_round := t.Memebers[0].Keepenvenergyupdate(currentSeason)
	fmt.Println("food_this_round", food_this_round)
	fmt.Println("water_this_round", water_this_round)
	fmt.Println("wood_this_round", wood_this_round)

	fmt.Println("simulation started for tribe ID: ", t.TribeID)
	for _, member := range t.Memebers {
		// keep alive
		if member.IsAlive() {
			member.UpdateResources(currentSeason, t.Getnumberofalive(), t.Getnumberofmoderate(), water_this_round, food_this_round, wood_this_round, greedy_switch)
			if member.water_level <= 0 || member.food_level <= 0 {
				member.deadflag = true
				fmt.Printf("Agent ID: %v, is dead due to low energy\n", member.AgentId)
			}
			fmt.Printf("Agent ID: %v, Role: %s\n", member.AgentId, member.GetAgentRole())
			// fmt.Printf("its environment: %s\n", member.Environment.GetType())
			fmt.Printf("Agent greedyiness: %s\n", member.Greediness.String())
		} else {
			// fmt.Printf("Agent ID: %v, is dead with water: %.2f, food: %.2f, wood: %.2f\n", member.AgentId, member.water_level, member.food_level, member.wood_level)
		}

	}
	t.CurrentResources = t.CalculateCurrentResources()
	Action := t.DecideActionsBasedOnSatisfaction()
	fmt.Println("Action: ", Action)
	// Check tribe type
	t.Check_tribe_type()

	//get neighbor info
	current_areas := make(map[int]*maps.Area)
	for _, member := range t.Memebers {
		// if member.CurrentArea not in current_areas, append it
		if _, ok := current_areas[member.CurrentArea.ID]; !ok {
			current_areas[member.CurrentArea.ID] = member.CurrentArea
		}
	}
	var areasSlice []*maps.Area
	for _, area := range current_areas {
		areasSlice = append(areasSlice, area)
	}

	// Attack other tribes
	if t.boat_making_signal && !t.Haveboat {
		if t.CurrentResources.Wood > (t.totalneeds_ofwood_forboat / float64(len(t.GetAliveMembers()))) {
			t.Make_a_boat()
			t.boat_making_signal = false
		}
	}
	t.Isattacking = false
	if t.attack_signal {
		// Find possible targets
		if t.CurrentResources.Food > 0.05 && t.CurrentResources.Water > 0.05 {
			possibleTargets := t.FindTargets(areasSlice, alltribes)
			if len(possibleTargets) > 0 { // Ensure there are targets available
				target := t.ChooseTarget(possibleTargets)
				attack_res := t.Attack(target)
				t.Isattacking = true
				t.TargetID = target.TribeID
				if attack_res {
					t.Lastattacksuccess = true
				}
				fmt.Printf("Bandit tribe ID: %v in area: %v, attacked role %s tribe ID: %v in area: %v\n", t.TribeID, t.Memebers[0].CurrentArea.ID, target.tribetype, target.TribeID, target.Memebers[0].CurrentArea.ID)
			} else {
				fmt.Println("No available targets for attack.")
			}
		}
		t.attack_signal = false
		// t.Memebers[0].Roleswitch(utils.Farmer_agent)
	}
	if t.shettle_signal {
		if t.CurrentResources.Wood > (t.totalneeds_ofwood_forshettle / float64(len(t.GetAliveMembers()))) {
			t.Make_a_shettle()
			t.shettle_signal = false
		}

	}
	if t.relocate_signal && t.Haveboat {
		// Consider shelter level as a factor in deciding to relocate
		// The higher the shelter level, the less likely they are to relocate
		shelterRelocationInfluence := math.Exp(-float64(t.Shettle_level) * 0.4) // Exponential decay to decrease relocation likelihood

		// Only proceed with relocation if random chance (influenced by shelter level) allows
		if rand.Float64() < shelterRelocationInfluence {
			potentialRelocationAreas := make(map[int]*maps.Area)

			// Populate potential relocation areas with direct neighbors
			for _, member := range t.Memebers {
				for _, neighbor := range member.CurrentArea.Neighbors {
					potentialRelocationAreas[neighbor.ID] = neighbor
				}
			}

			var bestArea *maps.Area
			var highestAvgGain float64

			if t.tribetype == utils.Bandit_Agent && len(t.SuccessfulAttackGains) > 0 {
				for areaID, area := range potentialRelocationAreas {
					if gain, exists := t.SuccessfulAttackGains[areaID]; exists && gain.AttackCount > 0 {
						avgGain := (gain.Water + gain.Food + gain.Wood) / float64(gain.AttackCount)
						if avgGain > highestAvgGain {
							highestAvgGain = avgGain
							bestArea = area
						}
					}
				}
			}

			if bestArea == nil {
				keys := make([]int, 0, len(potentialRelocationAreas))
				for k := range potentialRelocationAreas {
					keys = append(keys, k)
				}
				if len(keys) > 0 {
					randomKey := keys[rand.Intn(len(keys))]
					bestArea = potentialRelocationAreas[randomKey]
				}
			}

			if bestArea != nil {
				fmt.Println("Tribe ID: ", t.TribeID, "relocating to area ID: ", bestArea.ID, "with shelter level", t.Shettle_level, "which will be reset")
				for _, member := range t.GetAliveMembers() {
					member.Relocate(bestArea)
				}
				t.Shettle_level = 0 // Reset shelter level after relocation
			} else {
				fmt.Println("No available areas for relocation or decision made to stay due to high shelter investment.")
			}
		} else {
			fmt.Println("Tribe chooses to stay due to high shelter level.")
		}

		t.relocate_signal = false
	}

	// update defence rate, t.shettle_level is the addition 0.05 of defence rate
	t.defendrate = origional_defence_rate + float64(t.Shettle_level)*0.05
	t.deathrate = origional_death_rate - 0.02*float64(t.Shettle_level)

	// // Apply death rate
	// totalDeaths := int(t.numofdeath) + int(t.deathrate*float64(len(t.getAliveMembers())))
	// for i := 0; i < (totalDeaths); i++ {
	// 	t.killRandomAgent()
	// 	t.numofdeath -= 1
	// }

	// // Apply birth rate
	// if len(t.getAliveMembers()) != 1 {
	// 	totalBirths := int(t.numofbirth) + int(t.birthrate*float64(len(t.getAliveMembers())))
	// 	for i := 0; i < totalBirths; i++ {
	// 		t.addNewAgent()
	// 		t.numofbirth -= 1
	// 	}
	// }

	t.applyDeathRate()
	t.applyBirthRate()
	// Update attack & defence rate

	// _, _, average_wood := t.CalculateAverageResources()
	// if !t.Haveboat {
	// 	if average_wood >= (t.totalneeds_ofwood_forboat / float64(len(t.getAliveMembers()))) {
	// 		t.Make_a_boat()
	// 	}
	// } else {
	// 	if average_wood >= (t.totalneeds_ofwood_forshettle / float64(len(t.getAliveMembers()))) {
	// 		t.Make_a_shettle()
	// 	}

	// }

	t.UpdateSatisfaction()
	fmt.Printf("Total members number: %d, Alive members number: %d\n", len(t.Memebers), t.Getnumberofalive())
	fmt.Println("simulation ended")
}

func (t *Tribe) applyDeathRate() {
	for _, member := range t.Memebers {
		if rand.Float64() < t.deathrate { // Assuming deathrate is the probability of dying each cycle
			member.deadflag = true
		}
	}
}

func (t *Tribe) applyBirthRate() {
	if len(t.GetAliveMembers()) != 0 {
		t.PotentialBirths += t.birthrate * float64(len(t.GetAliveMembers())) // Calculate expected births
		for i := 0; i < int(t.PotentialBirths); i++ {
			t.addNewAgent()
			t.PotentialBirths -= 1
		}
	}
}

// func (t *Tribe) killRandomAgent() {
// 	aliveMembers := t.getAliveMembers()
// 	if len(aliveMembers) == 0 {
// 		return
// 	}
// 	randomIndex := rand.Intn(len(aliveMembers))
// 	aliveMembers[randomIndex].deadflag = true
// 	fmt.Println("Agent ID: ", aliveMembers[randomIndex].AgentId, "is killed")
// }

func (t *Tribe) GetAliveMembers() []*BaseTribe {
	var aliveMembers []*BaseTribe
	for _, member := range t.Memebers {
		if member.IsAlive() {
			aliveMembers = append(aliveMembers, member)
		}
	}
	return aliveMembers
}

func (t *Tribe) Getnumberofalive() int {
	aliveCount := 0
	for _, member := range t.Memebers {
		if member.IsAlive() {
			aliveCount++
		}
	}
	return aliveCount
}

func (t *Tribe) addNewAgent() {
	thistribetype := t.tribetype
	newAgent := GetIBaseTribe(thistribetype)
	//random get number from 1 to 10
	if t.porofgreedy > rand.Float64() {
		newAgent.Greediness = utils.Greedy
	}
	newAgent.Environment = t.Memebers[0].Environment
	newAgent.CurrentArea = t.Memebers[0].CurrentArea
	t.Memebers = append(t.Memebers, newAgent)
	fmt.Println("New agent ID: ", newAgent.AgentId, "is born")
}

// checking

func (t *Tribe) CalculateAverageResources() (float64, float64, float64) {
	var totalWater, totalFood, totalWood float64
	fmt.Println("Calculation started")
	aliveCount := 0

	for _, member := range t.Memebers {
		if member.IsAlive() {
			totalWater += member.GetWaterLevel()
			totalFood += member.GetFoodLevel()
			totalWood += member.GetWoodLevel()
			aliveCount++
		}
	}

	if aliveCount > 0 {
		fmt.Println("Calculation average")
		return totalWater / float64(aliveCount), totalFood / float64(aliveCount), totalWood / float64(aliveCount)
	}
	return 0, 0, 0
}

func (t *Tribe) CalculateTotalWood() (totalWood float64) {
	for _, member := range t.Memebers {
		totalWood += member.wood_level
	}
	return totalWood
}

func (t *Tribe) CalcuateTotalFood() (totalFood float64) {
	for _, member := range t.Memebers {
		totalFood += member.food_level
	}
	return totalFood
}

func (t *Tribe) CalcuateTotalWater() (totalWater float64) {
	for _, member := range t.Memebers {
		totalWater += member.water_level
	}
	return totalWater
}

func (t *Tribe) GetAllAgents() []*BaseTribe {
	return t.Memebers
}

func (t *Tribe) Getnumberofmoderate() int {
	count := 0
	for _, member := range t.Memebers {
		if member.Greediness == utils.Moderate {
			count++
		}
	}
	return count
}

func CheckSuccessfulAttack(attack_rate float64, target_defence float64) bool {
	porpotion := attack_rate / (attack_rate + target_defence)
	return (rand.Float64() < (1.2 * porpotion))
}

// func (t *Tribe) Attack
func (t *Tribe) Attack(target *Tribe) bool {
	success := CheckSuccessfulAttack(t.attackrate, target.defendrate)
	if success {
		waterStolenTotal, foodStolenTotal, woodStolenTotal := Beingattack(target, t.energytakenrate)

		// Retrieve the gains for the target area or initialize if it does not exist
		areaGains, exists := t.SuccessfulAttackGains[target.Memebers[0].CurrentArea.ID]
		if !exists {
			areaGains = ResourceMetrics{
				Water:       0,
				Food:        0,
				Wood:        0,
				AttackCount: 0,
			}
		}

		// Update the resource gains and the attack count
		areaGains.Water += waterStolenTotal
		areaGains.Food += foodStolenTotal
		areaGains.Wood += woodStolenTotal
		areaGains.AttackCount += 1 // Increment attack count
		t.SuccessfulAttackGains[target.Memebers[0].CurrentArea.ID] = areaGains

		// Logging the updated gains
		fmt.Printf("Updated cumulative gains for area ID %d: Water: %.2f, Food: %.2f, Wood: %.2f, Attacks: %d\n",
			target.Memebers[0].CurrentArea.ID, areaGains.Water, areaGains.Food, areaGains.Wood, areaGains.AttackCount)

		t.UpdateRanking(target, success)
	}

	// Deduct a small amount of resources after each attack attempt (successful or not)
	for _, member := range t.Memebers {
		member.water_level -= 0.05
		member.food_level -= 0.05
	}

	return success
}

func (t *Tribe) RecordAttackSuccess(targetAreaID int, waterStolen float64, foodStolen float64, woodStolen float64) {
	if gain, exists := t.SuccessfulAttackGains[targetAreaID]; exists {
		gain.Water += waterStolen
		gain.Food += foodStolen
		gain.Wood += woodStolen
		gain.AttackCount++
		t.SuccessfulAttackGains[targetAreaID] = gain
	} else {
		t.SuccessfulAttackGains[targetAreaID] = ResourceMetrics{
			Water:       waterStolen,
			Food:        foodStolen,
			Wood:        woodStolen,
			AttackCount: 1,
		}
	}
}

func Beingattack(target *Tribe, energytakenrate float64) (float64, float64, float64) {
	waterstolentotal := 0.0
	foodstolentotal := 0.0
	woodstolentotal := 0.0

	for _, member := range target.Memebers {
		if member.IsAlive() {
			waterStolen := member.water_level * energytakenrate
			foodStolen := member.food_level * energytakenrate
			woodStolen := member.wood_level * energytakenrate

			member.water_level -= waterStolen
			member.food_level -= foodStolen
			member.wood_level -= woodStolen

			waterstolentotal += waterStolen
			foodstolentotal += foodStolen
			woodstolentotal += woodStolen
		}
	}

	return waterstolentotal, foodstolentotal, woodstolentotal
}

func (t *Tribe) Check_tribe_type() {
	//check which role is the most in the tribe
	//if the most role is farmer, then the tribe is farmer
	//if the most role is forager, then the tribe is forager
	//if the most role is bandit, then the tribe is bandit
	num_of_farmer := 0
	num_of_forager := 0
	num_of_bandit := 0
	for _, member := range t.Memebers {
		if member.AgentType == utils.Bandit_Agent {
			num_of_bandit++
		}
		if member.AgentType == utils.Farmer_agent {
			num_of_farmer++
		}
		if member.AgentType == utils.Forager_agent {
			num_of_forager++
		}
	}

	if num_of_farmer > num_of_forager && num_of_farmer > num_of_bandit {
		t.tribetype = utils.Farmer_agent
	}
	if num_of_forager > num_of_farmer && num_of_forager > num_of_bandit {
		t.tribetype = utils.Forager_agent
	}
	if num_of_bandit >= num_of_farmer && num_of_bandit >= num_of_forager {
		t.tribetype = utils.Bandit_Agent
	}
}

func (t *Tribe) Make_a_boat() {
	totalWoodNeeded := t.totalneeds_ofwood_forboat
	totalWoodAvailable := 0.0
	for _, member := range t.Memebers {
		totalWoodAvailable += member.wood_level
	}
	contributionRate := totalWoodNeeded / totalWoodAvailable

	fmt.Println("Tribe ID: ", t.TribeID, "starts making a boat with contribution rate:", contributionRate)

	for _, member := range t.Memebers {
		woodContribution := member.wood_level * contributionRate
		member.Contribute_to_woodwork(woodContribution)
		if woodContribution < 0.01 {
			continue // Skip this member's contribution
		}
		fmt.Printf("Member ID: %v contributes %.2f wood to the boat\n", member.AgentId, woodContribution)
	}
	fmt.Println("Tribe ID: ", t.TribeID, "has successfully made a boat")
	t.Haveboat = true
}

func (t *Tribe) Make_a_shettle() {
	totalWoodNeeded := t.totalneeds_ofwood_forshettle
	totalWoodAvailable := 0.0
	for _, member := range t.Memebers {
		totalWoodAvailable += member.wood_level
	}
	contributionRate := totalWoodNeeded / totalWoodAvailable

	fmt.Println("Tribe ID: ", t.TribeID, "starts making shettle with contribution rate:", contributionRate)

	for _, member := range t.Memebers {
		woodContribution := member.wood_level * contributionRate
		// Check if contribution is too small, skip or adjust here
		member.Contribute_to_woodwork(woodContribution)
		if woodContribution < 0.01 {
			continue // Skip this member's contribution
		}
		member.wood_level -= woodContribution
		fmt.Printf("Member ID: %v contributes %.2f wood to the shelter\n", member.AgentId, woodContribution)
	}

	fmt.Println("Tribe ID: ", t.TribeID, "has successfully made a shettle")
	t.Shettle_level++
}

func (b *Tribe) FindTargets(areas []*maps.Area, allTribes []*Tribe) []*Tribe {
	var targets []*Tribe
	connectedAreas := make(map[int]*maps.Area)

	for _, area := range areas {
		connectedAreas[area.ID] = area
		for _, neighbor := range area.Neighbors {
			connectedAreas[neighbor.ID] = neighbor
		}
	}

	for _, tribe := range allTribes {
		if tribe == b || tribe.Getnumberofalive() == 0 {
			continue
		}

		tribeAreaID := tribe.Memebers[0].CurrentArea.ID
		if _, ok := connectedAreas[tribeAreaID]; ok {
			if _, exists := b.Ranking[tribe.TribeID]; !exists {
				// Initialize ranking for tribes not yet encountered
				b.Ranking[tribe.TribeID] = 0.0
			}
			targets = append(targets, tribe)
			fmt.Printf("the ranking for tribe ID: %v is %v\n", tribe.TribeID, b.Ranking[tribe.TribeID])
		}
	}

	return targets
}

func (b *Tribe) ChooseTarget(tribes []*Tribe) *Tribe {
	var bestTarget *Tribe
	highestRank := -math.MaxFloat64 // Initial low value to ensure any real value is higher
	for _, tribe := range tribes {
		if _, exists := b.Ranking[tribe.TribeID]; !exists {
			b.Ranking[tribe.TribeID] = 0.0 // Ensure the ranking exists
		}
		if tribe.TribeID != b.TribeID && b.Ranking[tribe.TribeID] > highestRank {
			highestRank = b.Ranking[tribe.TribeID]
			bestTarget = tribe
		}
	}
	return bestTarget
}

func (b *Tribe) UpdateRanking(target *Tribe, success bool) {
	if _, exists := b.Ranking[target.TribeID]; !exists {
		b.Ranking[target.TribeID] = 0.0
	}
	if success {
		b.Ranking[target.TribeID] -= 1 // Successful attack should increases perceived weakness but consider attack next round means lack of resource?  Debateable
	} else {
		b.Ranking[target.TribeID] -= 5 // Failed attack decreases perceived weakness
	}
}

func (t *Tribe) CalculateCurrentResources() ResourceMetrics {
	totalWater, totalFood, totalWood := t.CalculateAverageResources()
	return ResourceMetrics{
		Water: totalWater,
		Food:  totalFood,
		Wood:  totalWood,
	}
}

func (t *Tribe) UpdateSatisfaction() {
	satisfactionChange := t.CalculateSatisfactionChange(t.PreviousResources, t.CurrentResources)
	t.CurrentSatisfaction += satisfactionChange
	t.SatisfactionHistory = append(t.SatisfactionHistory, t.CurrentSatisfaction)
	//keep it short
	if len(t.SatisfactionHistory) > 6 {
		t.SatisfactionHistory = t.SatisfactionHistory[len(t.SatisfactionHistory)-6:]
	}
}

func (t *Tribe) CalculateSatisfactionChange(previous, current ResourceMetrics) float64 {
	deltaFood := current.Food - previous.Food
	deltaWater := current.Water - previous.Water
	deltaWood := current.Wood - previous.Wood

	satisfactionDelta := (deltaFood*0.4 + deltaWater*0.4 + deltaWood*0.2)
	return satisfactionDelta
}

func (t *Tribe) DecideActionsBasedOnSatisfaction() string {
	if len(t.SatisfactionHistory) >= 4 {
		// Check the last 5 rounds for long term reaction
		if t.longestDecliningStreak(6) >= 3 { // Consider making a boat to relocate if possible
			t.relocate_signal = true
			if !t.Haveboat {
				t.boat_making_signal = true
				fmt.Println("Making boat")
			}

			fmt.Println("Have boat signal is , ", t.Haveboat)
			return "Relocate to new area"
		}

	}

	if len(t.SatisfactionHistory) >= 3 {
		// Check the last 3 rounds for immediate reaction
		if t.longestDecliningStreak(6) >= 2 {
			if t.tribetype == utils.Bandit_Agent {
				t.attack_signal = true
				return "Prepare Attack"
			} else {
				t.shettle_signal = true
				return "Defense with increase shettler level" // For non-bandit tribes, consider making a shelter
			}
		}
	}
	return "Continue Normal Activities"
}

func (t *Tribe) longestDecliningStreak(rounds int) int {
	if len(t.SatisfactionHistory) < 2 {
		return 0
	}

	start := 0
	if rounds < len(t.SatisfactionHistory) {
		start = len(t.SatisfactionHistory) - rounds
	}
	recentHistory := t.SatisfactionHistory[start:]
	currentStreak := 0
	longestStreak := 0
	for i := 1; i < len(recentHistory); i++ {
		if recentHistory[i] < recentHistory[i-1] {
			currentStreak++
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}
	}
	return longestStreak
}
