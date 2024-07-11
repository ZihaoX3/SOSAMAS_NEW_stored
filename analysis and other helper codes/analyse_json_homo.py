import json
from statistics import mean

# Load the JSON file
with open('simulation_data.json') as f:
    data = json.load(f)

iterations = {}
food_levels = {}
wood_levels = {}
water_levels = {}

for iteration_data in data:
    for tribe in iteration_data['tribes']:
        tribe_role = tribe['role'][0]  # Get the first role of the tribe
        if tribe_role == 'Farmer_agent':  # Only process data for the 'Bandit_Agent' role
            if tribe_role not in iterations:
                iterations[tribe_role] = []
                food_levels[tribe_role] = []
                wood_levels[tribe_role] = []
                water_levels[tribe_role] = []
            iterations[tribe_role].append(iteration_data['iteration'])
            food_levels[tribe_role].append(tribe['food_level'])
            wood_levels[tribe_role].append(tribe['wood_level'])
            water_levels[tribe_role].append(tribe['water_level'])

results = {}
role = 'Farmer_agent'  # Define the role you want to analyze

if role in iterations.keys():
    results[role] = {
        'livespan': mean(iterations[role]),
        'average_food': mean(food_levels[role]),
        'average_wood': mean(wood_levels[role]),
        'average_water': mean(water_levels[role])
    }

# Print the results for the bandit role
print(results)

# Append the results to a file
with open('all_results_onlyfarmer.json', 'a') as f:
    json.dump(results, f, indent=4)