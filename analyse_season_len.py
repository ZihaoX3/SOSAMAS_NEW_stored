import json
import pandas as pd
import matplotlib.pyplot as plt

# List of greediness levels
length = [3,5,7]

# Initialize a DataFrame to hold all data
all_data = pd.DataFrame()

# Loop through each greediness level and load the corresponding JSON file
for size in length:
    with open(f'all_results_seasonlen{size}.json', 'r') as file:
        content = file.read()
        json_objects = content.split("\n")
        for json_str in json_objects:
            if json_str:
                data = json.loads(json_str)
                for role, stats in data.items():
                    all_data = all_data.append({
                        'Season Length': size,
                        'Role': role.replace('_agent', '').capitalize(),  # Convert 'Bandit_Agent' to 'Bandit'
                        'Lifespan': stats['livespan'],
                        'Average Food': stats['average_food'],
                        'Average Wood': stats['average_wood'],
                        'Average Water': stats['average_water']
                    }, ignore_index=True)

# Calculate the average lifespan and resource levels for each role at each greediness level
average_data = all_data.groupby(['Season Length', 'Role']).mean().reset_index()

# Plot the average lifespan for each role
plt.figure(figsize=(10, 6))
for role in average_data['Role'].unique():
    role_data = average_data[average_data['Role'] == role]
    plt.plot(role_data['Season Length'], role_data['Lifespan'], label=f'{role}')
plt.title('Average Lifespan per Role over Map size Levels')
plt.xlabel('Season Length')
plt.ylabel('Average Lifespan')
plt.legend()
plt.show()

# Plot the average food level for each role
plt.figure(figsize=(10, 6))
for role in average_data['Role'].unique():
    role_data = average_data[average_data['Role'] == role]
    plt.plot(role_data['Season Length'], role_data['Average Food'], label=f'{role}')
plt.title('Average Food Level per Role over Season Length Levels')
plt.xlabel('Season Length')
plt.ylabel('Average Food Level')
plt.legend()
plt.show()

# Plot the average wood level for each role
plt.figure(figsize=(10, 6))
for role in average_data['Role'].unique():
    role_data = average_data[average_data['Role'] == role]
    plt.plot(role_data['Season Length'], role_data['Average Wood'], label=f'{role}')
plt.title('Average Wood Level per Role over Season Length Levels')
plt.xlabel('Season Length')
plt.ylabel('Average Wood Level')
plt.legend()
plt.show()

# Plot the average water level for each role
plt.figure(figsize=(10, 6))
for role in average_data['Role'].unique():
    role_data = average_data[average_data['Role'] == role]
    plt.plot(role_data['Season Length'], role_data['Average Water'], label=f'{role}')
plt.title('Average Water Level per Role over Season Length')
plt.xlabel('Season Length')
plt.ylabel('Average Water Level')
plt.legend()
plt.show()