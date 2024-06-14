package environment

type Landscape interface {
	GetType() string
	InitializeResources()
	UpdateResources(season string)
	GetResources() (float64, float64, float64) // returns water, food, wood levels
	ExtractResources(water float64, food float64, wood float64) (float64, float64, float64)
	GetWaterLevel() float64
	GetFoodLevel() float64
	GetWoodLevel() float64
}

// Forest
const forest_wood_midgenrate = 1.5
const forest_wood_mingenrate = 1.0
const forest_water_maxgenrate = 2.5
const forest_water_midgenrate = 2.0
const forest_water_mingenrate = 1.5
const forest_food_maxgenrate = 2.0
const forest_food_midgenrate = 1.5
const forest_food_mingenrate = 1.0
const forest_generatediff = 0.5

// Desert
const desert_wood_genrate = 0.01
const desert_water_genrate = 0.01
const desert_food_genrate = 0.01
const desert_generatediff = 0.005

// Pasture
const pasture_wood_genrate = 0.3
const pasture_water_genrate = 0.3
const pasture_food_genrate = 0.2
const pasture_generatediff = 0.1
