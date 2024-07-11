# import json
# import pandas as pd
# import matplotlib.pyplot as plt

# # List of greediness levels
# greed_levels = [0, 20, 40, 60, 80, 100]

# # Initialize a DataFrame to hold all data
# all_data = pd.DataFrame()

# # Loop through each greediness level and load the corresponding JSON file
# for greed in greed_levels:
#     with open(f'all_results_greedy{greed}.json', 'r') as file:
#         content = file.read()
#         json_objects = content.split("\n")
#         for json_str in json_objects:
#             if json_str:
#                 data = json.loads(json_str)
#                 for role, stats in data.items():
#                     all_data = all_data.append({
#                         'Greediness': greed,
#                         'Role': role.replace('_agent', '').capitalize(),  # Convert 'Bandit_Agent' to 'Bandit'
#                         'Lifespan': stats['livespan'],
#                         'Average Food': stats['average_food'],
#                         'Average Wood': stats['average_wood'],
#                         'Average Water': stats['average_water']
#                     }, ignore_index=True)

# # Calculate the average lifespan and resource levels for each role at each greediness level
# average_data = all_data.groupby(['Greediness', 'Role']).mean().reset_index()

# # Plot the average lifespan for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Lifespan'], label=f'{role}')
# plt.title('Average Lifespan per Role over Normal Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Lifespan')
# plt.legend()
# plt.show()

# # Plot the average food level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Average Food'], label=f'{role}')
# plt.title('Average Food Level per Role over Normal Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Food Level')
# plt.legend()
# plt.show()

# # Plot the average wood level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Average Wood'], label=f'{role}')
# plt.title('Average Wood Level per Role over Normal Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Wood Level')
# plt.legend()
# plt.show()

# # Plot the average water level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Average Water'], label=f'{role}')
# plt.title('Average Water Level per Role over Normal Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Water Level')
# plt.legend()
# plt.show()



# import json
# import pandas as pd
# import matplotlib.pyplot as plt

# # List of greediness levels
# greed_levels = [0, 20, 40, 60, 80, 100]

# # Initialize a DataFrame to hold all data
# all_data = pd.DataFrame()

# # Loop through each greediness level and load the corresponding JSON file
# for greed in greed_levels:
#     with open(f'all_results_exgreedy{greed}.json', 'r') as file:
#         content = file.read()
#         json_objects = content.split("\n")
#         for json_str in json_objects:
#             if json_str:
#                 data = json.loads(json_str)
#                 for role, stats in data.items():
#                     all_data = all_data.append({
#                         'Greediness': greed,
#                         'Role': role.replace('_agent', '').capitalize(),  # Convert 'Bandit_Agent' to 'Bandit'
#                         'Lifespan': stats['livespan'],
#                         'Average Food': stats['average_food'],
#                         'Average Wood': stats['average_wood'],
#                         'Average Water': stats['average_water']
#                     }, ignore_index=True)

# # Calculate the average lifespan and resource levels for each role at each greediness level
# average_data = all_data.groupby(['Greediness', 'Role']).mean().reset_index()

# # Plot the average lifespan for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Lifespan'], label=f'{role}')
# plt.title('Average Lifespan per Role over EXTREME Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Lifespan')
# plt.legend()
# plt.show()

# # Plot the average food level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Average Food'], label=f'{role}')
# plt.title('Average Food Level per Role over EXTREME Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Food Level')
# plt.legend()
# plt.show()

# # Plot the average wood level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Average Wood'], label=f'{role}')
# plt.title('Average Wood Level per Role over EXTREME Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Wood Level')
# plt.legend()
# plt.show()

# # Plot the average water level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Greediness'], role_data['Average Water'], label=f'{role}')
# plt.title('Average Water Level per Role over EXTREME Greediness Levels')
# plt.xlabel('Greediness Level')
# plt.ylabel('Average Water Level')
# plt.legend()
# plt.show()




import json
import pandas as pd
import matplotlib.pyplot as plt

# List of greediness levels
greed_levels = [0, 20, 40, 60, 80, 100]

# Initialize DataFrames to hold all data
all_data_extreme = pd.DataFrame()
all_data_normal = pd.DataFrame()

# Load data from extreme greediness JSON files
for greed in greed_levels:
    with open(f'all_results_exgreedy{greed}.json', 'r') as file:
        content = file.read()
        json_objects = content.split("\n")
        for json_str in json_objects:
            if json_str:
                data = json.loads(json_str)
                for role, stats in data.items():
                    all_data_extreme = all_data_extreme.append({
                        'Greediness': greed,
                        'Role': role.replace('_agent', '').capitalize(),
                        'Lifespan': stats['livespan'],
                        'Average Food': stats['average_food'],
                        'Average Wood': stats['average_wood'],
                        'Average Water': stats['average_water']
                    }, ignore_index=True)

# Load data from normal greediness JSON files
for greed in greed_levels:
    with open(f'all_results_greedy{greed}.json', 'r') as file:
        content = file.read()
        json_objects = content.split("\n")
        for json_str in json_objects:
            if json_str:
                data = json.loads(json_str)
                for role, stats in data.items():
                    all_data_normal = all_data_normal.append({
                        'Greediness': greed,
                        'Role': role.replace('_agent', '').capitalize(),
                        'Lifespan': stats['livespan'],
                        'Average Food': stats['average_food'],
                        'Average Wood': stats['average_wood'],
                        'Average Water': stats['average_water']
                    }, ignore_index=True)

# Calculate averages for each role at each greediness level
average_data_extreme = all_data_extreme.groupby(['Greediness', 'Role']).mean().reset_index()
average_data_normal = all_data_normal.groupby(['Greediness', 'Role']).mean().reset_index()

# Calculate differences between extreme and normal averages
difference_data = average_data_normal.set_index(['Greediness', 'Role']) - average_data_extreme.set_index(['Greediness', 'Role'])
difference_data = difference_data.reset_index()

# Plot differences for each role
for metric in ['Lifespan', 'Average Food', 'Average Wood', 'Average Water']:
    plt.figure(figsize=(10, 6))
    for role in difference_data['Role'].unique():
        role_data = difference_data[difference_data['Role'] == role]
        plt.plot(role_data['Greediness'], role_data[metric], label=f'{role}')
    plt.title(f'Difference in {metric} per Role between Normal and extreme Greediness Levels')
    plt.xlabel('Greediness Level')
    plt.ylabel(f'Difference in {metric}')
    plt.legend()
    plt.show()
