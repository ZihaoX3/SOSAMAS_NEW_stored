package environment

import (
	"SOSAMAS_FYP/internal/common/utils"
)

type Pasture struct {
	water, food, wood float64
}

func (p *Pasture) GetType() string {
	return "Pasture"
}

func (p *Pasture) InitializeResources() {
	p.water = utils.GenerateRandomFloat(3, 4)
	p.food = utils.GenerateRandomFloat(3, 4)
	p.wood = utils.GenerateRandomFloat(2, 3)
}

func (p *Pasture) UpdateResources(season string) {
	if season == utils.Spring.String() {
		p.water += utils.GenerateRandomFloat(min(pasture_water_genrate-pasture_generatediff, 0), pasture_water_genrate+pasture_generatediff)
		p.food += utils.GenerateRandomFloat(min(pasture_food_genrate-pasture_generatediff, 0), pasture_food_genrate+pasture_generatediff)
		p.wood += utils.GenerateRandomFloat(min(pasture_wood_genrate-pasture_generatediff, 0), pasture_wood_genrate+pasture_generatediff)
	}
	if season == utils.Summer.String() {
		p.water += utils.GenerateRandomFloat(min(pasture_water_genrate-pasture_generatediff, 0), pasture_water_genrate+pasture_generatediff)
		p.food += utils.GenerateRandomFloat(min(pasture_food_genrate-pasture_generatediff, 0), pasture_food_genrate+pasture_generatediff)
		p.wood += utils.GenerateRandomFloat(min(pasture_wood_genrate-pasture_generatediff, 0), pasture_wood_genrate+pasture_generatediff)
	}
	if season == utils.Autumn.String() {
		p.water += utils.GenerateRandomFloat(min(pasture_water_genrate-pasture_generatediff, 0), pasture_water_genrate+pasture_generatediff)
		p.food += utils.GenerateRandomFloat(min(pasture_food_genrate-pasture_generatediff, 0), pasture_food_genrate+pasture_generatediff)
		p.wood += utils.GenerateRandomFloat(min(pasture_wood_genrate-pasture_generatediff, 0), pasture_wood_genrate+pasture_generatediff)
	}
	if season == utils.Winter.String() {
		p.water += utils.GenerateRandomFloat(min(pasture_water_genrate-pasture_generatediff, 0), pasture_water_genrate+pasture_generatediff)
		p.food += utils.GenerateRandomFloat(min(pasture_food_genrate-pasture_generatediff, 0), pasture_food_genrate+pasture_generatediff)
		p.wood += utils.GenerateRandomFloat(min(pasture_wood_genrate-pasture_generatediff, 0), pasture_wood_genrate+pasture_generatediff)
	}
}

func (p *Pasture) GetResources() (float64, float64, float64) {
	return p.water, p.food, p.wood
}

func (p *Pasture) ExtractResources(waterReq, foodReq, woodReq float64) (float64, float64, float64) {
	// find the max can get from pasture
	waterExtracted := min(waterReq, p.water)
	foodExtracted := min(foodReq, p.food)
	woodExtracted := min(woodReq, p.wood)

	p.water -= waterExtracted
	p.food -= foodExtracted
	p.wood -= woodExtracted

	return waterExtracted, foodExtracted, woodExtracted
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func (p *Pasture) GetWaterLevel() float64 {
	return p.water
}

func (p *Pasture) GetFoodLevel() float64 {
	return p.food
}

func (p *Pasture) GetWoodLevel() float64 {
	return p.wood
}
