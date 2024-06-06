package environment

import (
	"SOSAMAS_FYP/internal/common/utils"
)

type Forest struct {
	water, food, wood float64
}

func (f *Forest) GetType() string {
	return "Forest"
}

func (f *Forest) InitializeResources() {
	f.water = utils.GenerateRandomFloat(3, 4)
	f.food = utils.GenerateRandomFloat(3, 4)
	f.wood = utils.GenerateRandomFloat(8, 10)
}

func (f *Forest) UpdateResources(season string) {
	switch season {
	case utils.Spring.String():
		f.water += utils.GenerateRandomFloat(forest_water_midgenrate-forest_generatediff, forest_water_midgenrate+forest_generatediff)
		f.food += utils.GenerateRandomFloat(forest_food_midgenrate-forest_generatediff, forest_food_midgenrate+forest_generatediff)
		f.wood += utils.GenerateRandomFloat(forest_wood_midgenrate-forest_generatediff, forest_wood_midgenrate+forest_generatediff)
	case utils.Summer.String():
		f.water += utils.GenerateRandomFloat(forest_water_maxgenrate-forest_generatediff, forest_water_maxgenrate+forest_generatediff)
		f.food += utils.GenerateRandomFloat(forest_food_midgenrate-forest_generatediff, forest_food_midgenrate+forest_generatediff)
		f.wood += utils.GenerateRandomFloat(forest_wood_midgenrate-forest_generatediff, forest_wood_midgenrate+forest_generatediff)
	case utils.Autumn.String():
		f.water += utils.GenerateRandomFloat(forest_water_midgenrate-forest_generatediff, forest_water_midgenrate+forest_generatediff)
		f.food += utils.GenerateRandomFloat(forest_food_maxgenrate-forest_generatediff, forest_food_maxgenrate+forest_generatediff)
		f.wood += utils.GenerateRandomFloat(forest_wood_midgenrate-forest_generatediff, forest_wood_midgenrate+forest_generatediff)
	case utils.Winter.String():
		f.water += utils.GenerateRandomFloat(forest_water_mingenrate-forest_generatediff, forest_water_mingenrate+forest_generatediff)
		f.food += utils.GenerateRandomFloat(forest_food_mingenrate-forest_generatediff, forest_food_mingenrate+forest_generatediff)
		f.wood += utils.GenerateRandomFloat(forest_wood_mingenrate-forest_generatediff, forest_wood_mingenrate+forest_generatediff)
	}
}

func (f *Forest) GetResources() (float64, float64, float64) {
	return f.water, f.food, f.wood
}

func (f *Forest) GetWaterLevel() float64 {
	return f.water
}

func (f *Forest) GetFoodLevel() float64 {
	return f.food

}

func (f *Forest) GetWoodLevel() float64 {
	return f.wood
}

// environment reduction
func (f *Forest) ExtractResources(waterReq, foodReq, woodReq float64) (float64, float64, float64) {
	// find the max can get from pasture
	waterExtracted := min(waterReq, f.water)
	foodExtracted := min(foodReq, f.food)
	woodExtracted := min(woodReq, f.wood)

	f.water -= waterExtracted
	f.food -= foodExtracted
	f.wood -= woodExtracted

	return waterExtracted, foodExtracted, woodExtracted
}
