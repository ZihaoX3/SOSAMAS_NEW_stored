# import json
# import pandas as pd
# import matplotlib.pyplot as plt

# # List of greediness levels
# size_levels = [5,10,15]

# # Initialize a DataFrame to hold all data
# all_data = pd.DataFrame()

# # Loop through each greediness level and load the corresponding JSON file
# for size in size_levels:
#     with open(f'all_results_tribesize{size}.json', 'r') as file:
#         content = file.read()
#         json_objects = content.split("\n")
#         for json_str in json_objects:
#             if json_str:
#                 data = json.loads(json_str)
#                 for role, stats in data.items():
#                     all_data = all_data.append({
#                         'Size': size,
#                         'Role': role.replace('_agent', '').capitalize(),  # Convert 'Bandit_Agent' to 'Bandit'
#                         'Lifespan': stats['livespan'],
#                         'Average Food': stats['average_food'],
#                         'Average Wood': stats['average_wood'],
#                         'Average Water': stats['average_water']
#                     }, ignore_index=True)

# # Calculate the average lifespan and resource levels for each role at each greediness level
# average_data = all_data.groupby(['Size', 'Role']).mean().reset_index()

# # Plot the average lifespan for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Size'], role_data['Lifespan'], label=f'{role}')
# plt.title('Average Lifespan per Role over Tribe Size Levels')
# plt.xlabel('Size Level')
# plt.ylabel('Average Lifespan')
# plt.legend()
# plt.show()

# # Plot the average food level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Size'], role_data['Average Food'], label=f'{role}')
# plt.title('Average Food Level per Role over Tribe Size Levels')
# plt.xlabel('Size Level')
# plt.ylabel('Average Food Level')
# plt.legend()
# plt.show()

# # Plot the average wood level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Size'], role_data['Average Wood'], label=f'{role}')
# plt.title('Average Wood Level per Role over Tribe Size Levels')
# plt.xlabel('Size Level')
# plt.ylabel('Average Wood Level')
# plt.legend()
# plt.show()

# # Plot the average water level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Size'], role_data['Average Water'], label=f'{role}')
# plt.title('Average Water Level per Role over Tribe Size Levels')
# plt.xlabel('Size Level')
# plt.ylabel('Average Water Level')
# plt.legend()
# plt.show()

# import json
# import pandas as pd
# import matplotlib.pyplot as plt

# # List of greediness levels
# num_levels = [10,15,20]

# # Initialize a DataFrame to hold all data
# all_data = pd.DataFrame()

# # Loop through each greediness level and load the corresponding JSON file
# for num in num_levels:
#     with open(f'all_results_tribenum{num}.json', 'r') as file:
#         content = file.read()
#         json_objects = content.split("\n")
#         for json_str in json_objects:
#             if json_str:
#                 data = json.loads(json_str)
#                 for role, stats in data.items():
#                     all_data = all_data.append({
#                         'Number': num,
#                         'Role': role.replace('_agent', '').capitalize(),  # Convert 'Bandit_Agent' to 'Bandit'
#                         'Lifespan': stats['livespan'],
#                         'Average Food': stats['average_food'],
#                         'Average Wood': stats['average_wood'],
#                         'Average Water': stats['average_water']
#                     }, ignore_index=True)

# # Calculate the average lifespan and resource levels for each role at each greediness level
# average_data = all_data.groupby(['Number', 'Role']).mean().reset_index()

# # Plot the average lifespan for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Number'], role_data['Lifespan'], label=f'{role}')
# plt.title('Average Lifespan per Role over Tribe Number Levels')
# plt.xlabel('Tribe Number')
# plt.ylabel('Average Lifespan')
# plt.legend()
# plt.show()

# # Plot the average food level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Number'], role_data['Average Food'], label=f'{role}')
# plt.title('Average Food Level per Role over Tribe Number Levels')
# plt.xlabel('Tribe Number')
# plt.ylabel('Average Food Level')
# plt.legend()
# plt.show()

# # Plot the average wood level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Number'], role_data['Average Wood'], label=f'{role}')
# plt.title('Average Wood Level per Role over Tribe Number Levels')
# plt.xlabel('Tribe Number')
# plt.ylabel('Average Wood Level')
# plt.legend()
# plt.show()

# # Plot the average water level for each role
# plt.figure(figsize=(10, 6))
# for role in average_data['Role'].unique():
#     role_data = average_data[average_data['Role'] == role]
#     plt.plot(role_data['Number'], role_data['Average Water'], label=f'{role}')
# plt.title('Average Water Level per Role over Tribe Number Levels')
# plt.xlabel('Tribe Number')
# plt.ylabel('Average Water Level')
# plt.legend()
# plt.show()


# Compare  10* 10 and 20*5  ---> tribe size 10 and tribe number 20
import json
import pandas as pd
import matplotlib.pyplot as plt

# Initialize DataFrames to hold all data
data_tribesize10 = pd.DataFrame()
data_tribenum20 = pd.DataFrame()

# Load data from tribesize10.json
with open('all_results_tribesize10.json', 'r') as file:
    content = file.read()
    json_objects = content.split("\n")
    for json_str in json_objects:
        if json_str:
            data = json.loads(json_str)
            for role, stats in data.items():
                data_tribesize10 = data_tribesize10.append({
                    'Role': role.replace('_agent', '').capitalize(),
                    'Lifespan': stats['livespan'],
                    'Average Food': stats['average_food'],
                    'Average Wood': stats['average_wood'],
                    'Average Water': stats['average_water']
                }, ignore_index=True)

# Load data from tribenum20.json
with open('all_results_tribenum20.json', 'r') as file:
    content = file.read()
    json_objects = content.split("\n")
    for json_str in json_objects:
        if json_str:
            data = json.loads(json_str)
            for role, stats in data.items():
                data_tribenum20 = data_tribenum20.append({
                    'Role': role.replace('_agent', '').capitalize(),
                    'Lifespan': stats['livespan'],
                    'Average Food': stats['average_food'],
                    'Average Wood': stats['average_wood'],
                    'Average Water': stats['average_water']
                }, ignore_index=True)

import numpy as np

average_data_tribesize10 = data_tribesize10.groupby('Role').mean().reset_index()
average_data_tribenum20 = data_tribenum20.groupby('Role').mean().reset_index()
# Define bar width
bar_width = 0.35

# Define positions of bars for tribe size 10 and tribe number 20
r1 = np.arange(len(average_data_tribesize10))
r2 = [x + bar_width for x in r1]

# Plot averages for each role
plt.figure(figsize=(10, 6))

plt.bar(r1, average_data_tribesize10['Lifespan'], color='b', width=bar_width, edgecolor='grey', label='Tribe Size 10 Number 10')
plt.bar(r2, average_data_tribenum20['Lifespan'], color='g', width=bar_width, edgecolor='grey', label='Tribe Number 20 Number 5')

plt.title('Average Lifespan per Role for Different Tribe Sizes and Numbers')
plt.xlabel('Role')
plt.ylabel('Average Lifespan')
plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribesize10))], average_data_tribesize10['Role'])
plt.legend()
plt.show()

# Plot the average food level for each role
plt.figure(figsize=(10, 6))

plt.bar(r1, average_data_tribesize10['Average Food'], color='b', width=bar_width, edgecolor='grey', label='Tribe Size 10  Number 10')
plt.bar(r2, average_data_tribenum20['Average Food'], color='g', width=bar_width, edgecolor='grey', label='Tribe Number 20  Number 5')

plt.title('Average Food Level per Role for Different Tribe Sizes and Numbers')
plt.xlabel('Role')
plt.ylabel('Average Food Level')
plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribesize10))], average_data_tribesize10['Role'])
plt.legend()
plt.show()


# Plot the average wood level for each role
plt.figure(figsize=(10, 6))

plt.bar(r1, average_data_tribesize10['Average Wood'], color='b', width=bar_width, edgecolor='grey', label='Tribe Size 10  Number 10')
plt.bar(r2, average_data_tribenum20['Average Wood'], color='g', width=bar_width, edgecolor='grey', label='Tribe Number 20  Number 5')

plt.title('Average Wood Level per Role for Different Tribe Sizes and Numbers')
plt.xlabel('Role')
plt.ylabel('Average Wood Level')
plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribesize10))], average_data_tribesize10['Role'])
plt.legend()
plt.show()


# Plot the average water level for each role
plt.figure(figsize=(10, 6))

plt.bar(r1, average_data_tribesize10['Average Water'], color='b', width=bar_width, edgecolor='grey', label='Tribe Size 10  Number 10')
plt.bar(r2, average_data_tribenum20['Average Water'], color='g', width=bar_width, edgecolor='grey', label='Tribe Number 20  Number 5')

plt.title('Average Water Level per Role for Different Tribe Sizes and Numbers')
plt.xlabel('Role')
plt.ylabel('Average Water Level')
plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribesize10))], average_data_tribesize10['Role'])
plt.legend()
plt.show()

