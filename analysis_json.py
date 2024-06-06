import json
from statistics import mean

# Load the JSON file
with open('simulation_data.json') as f:
    data = json.load(f)

# Initialize variables
iterations = {}
food_levels = {}
wood_levels = {}
water_levels = {}

# Loop through the data
for iteration_data in data:
    for tribe in iteration_data['tribes']:
        role = tribe['role'][0]
        if role not in iterations:
            iterations[role] = []
            food_levels[role] = []
            wood_levels[role] = []
            water_levels[role] = []
        iterations[role].append(iteration_data['iteration'])
        food_levels[role].append(tribe['food_level'])
        wood_levels[role].append(tribe['wood_level'])
        water_levels[role].append(tribe['water_level'])

# Calculate averages and append to results file
results = {}
# do in bandit, forager and farmer in order
for role in ['Bandit_Agent', 'Farmer_agent', 'Forager_agent']:
    results[role] = {
        'livespan': mean(iterations[role]),
        'average_food': mean(food_levels[role]),
        'average_wood': mean(wood_levels[role]),
        'average_water': mean(water_levels[role])
    }

# Append the results to a file
with open('all_results_tribesize15.json', 'a') as f:
    json.dump(results, f)
    f.write('\n')
