package environment

import (
	"SOSAMAS_FYP/internal/common/utils"
)

type Desert struct {
	water, food, wood float64
}

func (d *Desert) GetType() string {
	return "Desert"
}

func (d *Desert) InitializeResources() {
	d.water = utils.GenerateRandomFloat(1, 1.5)
	d.food = utils.GenerateRandomFloat(1, 1.5)
	d.wood = utils.GenerateRandomFloat(0.5, 1)
}

func (d *Desert) UpdateResources(season string) {
	if season == utils.Spring.String() {
		d.water += utils.GenerateRandomFloat(min(desert_water_genrate-desert_generatediff, 0), desert_water_genrate+desert_generatediff)
		d.food += utils.GenerateRandomFloat(min(desert_food_genrate-desert_generatediff, 0), desert_food_genrate+desert_generatediff)
		d.wood += utils.GenerateRandomFloat(min(desert_wood_genrate-desert_generatediff, 0), desert_wood_genrate+desert_generatediff)
	}
	if season == utils.Summer.String() {
		d.water += utils.GenerateRandomFloat(min(desert_water_genrate-desert_generatediff, 0), desert_water_genrate+desert_generatediff)
		d.food += utils.GenerateRandomFloat(min(desert_food_genrate-desert_generatediff, 0), desert_food_genrate+desert_generatediff)
		d.wood += utils.GenerateRandomFloat(min(desert_wood_genrate-desert_generatediff, 0), desert_wood_genrate+desert_generatediff)
	}
	if season == utils.Autumn.String() {
		d.water += utils.GenerateRandomFloat(min(desert_water_genrate-desert_generatediff, 0), desert_water_genrate+desert_generatediff)
		d.food += utils.GenerateRandomFloat(min(desert_food_genrate-desert_generatediff, 0), desert_food_genrate+desert_generatediff)
		d.wood += utils.GenerateRandomFloat(min(desert_wood_genrate-desert_generatediff, 0), desert_wood_genrate+desert_generatediff)
	}
	if season == utils.Winter.String() {
		d.water += utils.GenerateRandomFloat(min(desert_water_genrate-desert_generatediff, 0), desert_water_genrate+desert_generatediff)
		d.food += utils.GenerateRandomFloat(min(desert_food_genrate-desert_generatediff, 0), desert_food_genrate+desert_generatediff)
		d.wood += utils.GenerateRandomFloat(min(desert_wood_genrate-desert_generatediff, 0), desert_wood_genrate+desert_generatediff)
	}
}

func (d *Desert) GetResources() (float64, float64, float64) {
	return d.water, d.food, d.wood
}

func (d *Desert) ExtractResources(waterReq, foodReq, woodReq float64) (float64, float64, float64) {
	// find the max can get from desert
	waterExtracted := min(waterReq, d.water)
	foodExtracted := min(foodReq, d.food)
	woodExtracted := min(woodReq, d.wood)

	d.water -= waterExtracted
	d.food -= foodExtracted
	d.wood -= woodExtracted

	return waterExtracted, foodExtracted, woodExtracted
}
func (d *Desert) GetWaterLevel() float64 {
	return d.water
}

func (d *Desert) GetFoodLevel() float64 {
	return d.food
}

func (d *Desert) GetWoodLevel() float64 {
	return d.wood
}
