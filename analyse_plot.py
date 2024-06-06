import json
import matplotlib.pyplot as plt

# Initialize lists to store the average iterations for each role
forager_iterations = []
farmer_iterations = []
bandit_iterations = []

# Read the results from the file
with open('all_results.json') as f:
    for line in f:
        result = json.loads(line)
        forager_iterations.append(result.get("Forager_agent", {}).get("livespan", 0))
        farmer_iterations.append(result.get("Farmer_agent", {}).get("livespan", 0))
        bandit_iterations.append(result.get("Bandit_Agent", {}).get("livespan", 0))

# Plotting function
def plot_iterations(forager, farmer, bandit):
    plt.figure(figsize=(10, 6))
    plt.plot(forager, label='Forager')
    plt.plot(farmer, label='Farmer')
    plt.plot(bandit, label='Bandit')
    plt.xlabel('Simulation Run')
    plt.ylabel('Average Lifespan')
    plt.title('Average Iterations (Lifespan) Over 150 Simulations')
    plt.legend()
    plt.grid(True)
    plt.show()

# Plot the average iterations for each role
plot_iterations(forager_iterations, farmer_iterations, bandit_iterations)

# plot average food, water and wood for bandit tribe 
bandit_food = []
bandit_water = []
bandit_wood = []

# plot average food, water and wood for forager tribe
forager_food = []
forager_water = []
forager_wood = []

farmer_wood = []
farmer_food = []
farmer_water = []

# Read the results from the file
with open('all_results.json') as f:
    for line in f:
        result = json.loads(line)
        bandit_food.append(result.get("Bandit_Agent", {}).get("average_food", 0))
        bandit_water.append(result.get("Bandit_Agent", {}).get("average_water", 0))
        bandit_wood.append(result.get("Bandit_Agent", {}).get("average_wood", 0))
        forager_food.append(result.get("Forager_agent", {}).get("average_food", 0))
        forager_water.append(result.get("Forager_agent", {}).get("average_water", 0))
        forager_wood.append(result.get("Forager_agent", {}).get("average_wood", 0))
        farmer_food.append(result.get("Farmer_agent", {}).get("average_food", 0))
        farmer_water.append(result.get("Farmer_agent", {}).get("average_water", 0))
        farmer_wood.append(result.get("Farmer_agent", {}).get("average_wood", 0))


# Plotting function
def plot_resources_bandit(food, water, wood):
    plt.figure(figsize=(10, 6))
    plt.plot(food, label='Food')
    plt.plot(water, label='Water')
    plt.plot(wood, label='Wood')
    plt.xlabel('Simulation Run')
    plt.ylabel('Average Resources')
    plt.title('Average Resources Over 150 Simulations for bandit tribe')
    plt.legend()
    plt.grid(True)
    plt.show()

# Plot the average resources for the bandit tribe
plot_resources_bandit(bandit_food, bandit_water, bandit_wood)

def plot_resources_forager(food, water, wood):
    plt.figure(figsize=(10, 6))
    plt.plot(food, label='Food')
    plt.plot(water, label='Water')
    plt.plot(wood, label='Wood')
    plt.xlabel('Simulation Run')
    plt.ylabel('Average Resources')
    plt.title('Average Resources Over 150 Simulations for forager tribe')
    plt.legend()
    plt.grid(True)
    plt.show()

# Plot the average resources for the forager tribe
plot_resources_forager(forager_food, forager_water, forager_wood)

def plot_resources_farmer(food, water, wood):
    plt.figure(figsize=(10, 6))
    plt.plot(food, label='Food')
    plt.plot(water, label='Water')
    plt.plot(wood, label='Wood')
    plt.xlabel('Simulation Run')
    plt.ylabel('Average Resources')
    plt.title('Average Resources Over 150 Simulations for farmer tribe')
    plt.legend()
    plt.grid(True)
    plt.show()

# Plot the average resources for the farmer tribe
plot_resources_farmer(farmer_food, farmer_water, farmer_wood)



def plot_avg_food_forager_farmer_bandit(forager, farmer, bandit):
    plt.figure(figsize=(10, 6))
    plt.plot(forager, label='Forager')
    plt.plot(farmer, label='Farmer')
    plt.plot(bandit, label='Bandit')
    plt.xlabel('Simulation Run')
    plt.ylabel('Average Food Level')
    plt.title('Average Food Level Over 150 Simulations')
    plt.legend()
    plt.grid(True)
    plt.show()


# Plot the average food level for each role
plot_avg_food_forager_farmer_bandit(forager_food, farmer_food, bandit_food)

def plot_avg_water_forager_farmer_bandit(forager, farmer, bandit):
    plt.figure(figsize=(10, 6))
    plt.plot(forager, label='Forager')
    plt.plot(farmer, label='Farmer')
    plt.plot(bandit, label='Bandit')
    plt.xlabel('Simulation Run')
    plt.ylabel('Average Water Level')
    plt.title('Average Water Level Over 150 Simulations')
    plt.legend()
    plt.grid(True)
    plt.show()


# Plot the average water level for each role
plot_avg_water_forager_farmer_bandit(forager_water, farmer_water, bandit_water)


def plot_avg_wood_forager_farmer_bandit(forager, farmer, bandit):
    plt.figure(figsize=(10, 6))
    plt.plot(forager, label='Forager')
    plt.plot(farmer, label='Farmer')
    plt.plot(bandit, label='Bandit')
    plt.xlabel('Simulation Run')
    plt.ylabel('Average Wood Level')
    plt.title('Average Wood Level Over 150 Simulations')
    plt.legend()
    plt.grid(True)
    plt.show()
    
# Plot the average wood level for each role
plot_avg_wood_forager_farmer_bandit(forager_wood, farmer_wood, bandit_wood)



