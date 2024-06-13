# import json
# import pandas as pd

# # Load data from normal.json
# with open('all_results_normal.json', 'r') as file:
#     data_normal = pd.DataFrame([json.loads(line) for line in file])

# # Load data from all_desert.json
# with open('alL_results_all_desert.json', 'r') as file:
#     data_all_desert = pd.DataFrame([json.loads(line) for line in file])

# # Calculate averages for each role
# average_data_normal = data_normal.groupby('Role').mean().reset_index()
# average_data_all_desert = data_all_desert.groupby('Role').mean().reset_index()

# import numpy as np
# import matplotlib.pyplot as plt

# # Define bar width
# bar_width = 0.35

# # Define positions of bars for normal and all desert environments
# r1 = np.arange(len(average_data_normal))
# r2 = [x + bar_width for x in r1]

# # Plot averages for each role
# for metric in ['Lifespan', 'Average Food', 'Average Wood', 'Average Water']:
#     plt.figure(figsize=(10, 6))

#     plt.bar(r1, average_data_normal[metric], color='b', width=bar_width, edgecolor='grey', label='Normal')
#     plt.bar(r2, average_data_all_desert[metric], color='g', width=bar_width, edgecolor='grey', label='All Desert')

#     plt.title(f'Average {metric} per Role for Normal and All Desert Environments')
#     plt.xlabel('Role')
#     plt.ylabel(f'Average {metric}')
#     plt.xticks([r + bar_width / 2 for r in range(len(average_data_normal))], average_data_normal['Role'])
#     plt.legend()

#     plt.show()


# import json
# import pandas as pd
# import matplotlib.pyplot as plt

# # Initialize DataFrames to hold all data
# data_tribenormal = pd.DataFrame()
# data_tribealldesert = pd.DataFrame()

# # Load data from tribesize10.json
# with open('all_results_normal.json', 'r') as file:
#     content = file.read()
#     json_objects = content.split("\n")
#     for json_str in json_objects:
#         if json_str:
#             data = json.loads(json_str)
#             for role, stats in data.items():
#                 data_tribenormal = data_tribenormal.append({
#                     'Role': role.replace('_agent', '').capitalize(),
#                     'Lifespan': stats['livespan'],
#                     'Average Food': stats['average_food'],
#                     'Average Wood': stats['average_wood'],
#                     'Average Water': stats['average_water']
#                 }, ignore_index=True)

# # Load data from tribenum20.json
# with open('all_results_all_desert.json', 'r') as file:
#     content = file.read()
#     json_objects = content.split("\n")
#     for json_str in json_objects:
#         if json_str:
#             data = json.loads(json_str)
#             for role, stats in data.items():
#                 data_tribealldesert = data_tribealldesert.append({
#                     'Role': role.replace('_agent', '').capitalize(),
#                     'Lifespan': stats['livespan'],
#                     'Average Food': stats['average_food'],
#                     'Average Wood': stats['average_wood'],
#                     'Average Water': stats['average_water']
#                 }, ignore_index=True)

# import numpy as np

# average_data_tribenormal = data_tribenormal.groupby('Role').mean().reset_index()
# average_data_tribealldesert = data_tribealldesert.groupby('Role').mean().reset_index()
# # Define bar width
# bar_width = 0.35

# # Define positions of bars for tribe size 10 and tribe number 20
# r1 = np.arange(len(average_data_tribenormal))
# r2 = [x + bar_width for x in r1]

# # Plot averages for each role
# plt.figure(figsize=(10, 6))

# plt.bar(r1, average_data_tribenormal['Lifespan'], color='b', width=bar_width, edgecolor='grey', label='Normal')
# plt.bar(r2, average_data_tribealldesert['Lifespan'], color='g', width=bar_width, edgecolor='grey', label='All Desert')

# plt.title('Average Lifespan per Role for Normal and All Desert Environments')
# plt.xlabel('Role')
# plt.ylabel('Average Lifespan')
# plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribenormal))], average_data_tribenormal['Role'])
# plt.legend()
# plt.show()

# # Plot the average food level for each role
# plt.figure(figsize=(10, 6))

# plt.bar(r1, average_data_tribenormal['Average Food'], color='b', width=bar_width, edgecolor='grey', label='Normal')
# plt.bar(r2, average_data_tribealldesert['Average Food'], color='g', width=bar_width, edgecolor='grey', label='All Desert')

# plt.title('Average Food Level per Role for Normal and All Desert Environments')
# plt.xlabel('Role')
# plt.ylabel('Average Food Level')
# plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribenormal))], average_data_tribenormal['Role'])
# plt.legend()
# plt.show()


# # Plot the average wood level for each role
# plt.figure(figsize=(10, 6))

# plt.bar(r1, average_data_tribenormal['Average Wood'], color='b', width=bar_width, edgecolor='grey', label='Normal')
# plt.bar(r2, average_data_tribealldesert['Average Wood'], color='g', width=bar_width, edgecolor='grey', label='All Desert')

# plt.title('Average Wood Level per Role for Normal and All Desert Environments')
# plt.xlabel('Role')
# plt.ylabel('Average Wood Level')
# plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribenormal))], average_data_tribenormal['Role'])
# plt.legend()
# plt.show()


# # Plot the average water level for each role
# plt.figure(figsize=(10, 6))

# plt.bar(r1, average_data_tribenormal['Average Water'], color='b', width=bar_width, edgecolor='grey', label='Normal')
# plt.bar(r2, average_data_tribealldesert['Average Water'], color='g', width=bar_width, edgecolor='grey', label='All Desert')

# plt.title('Average Water Level per Role for Normal and All Desert Environments')
# plt.xlabel('Role')
# plt.ylabel('Average Water Level')
# plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribenormal))], average_data_tribenormal['Role'])
# plt.legend()
# plt.show()


import json
import pandas as pd
import matplotlib.pyplot as plt

# Initialize DataFrames to hold all data
data_tribenormal = pd.DataFrame()
data_tribeallpasture = pd.DataFrame()
data_tribeallforest = pd.DataFrame()

# Load data from tribesize10.json
with open('all_results_normal.json', 'r') as file:
    content = file.read()
    json_objects = content.split("\n")
    for json_str in json_objects:
        if json_str:
            data = json.loads(json_str)
            for role, stats in data.items():
                data_tribenormal = data_tribenormal.append({
                    'Role': role.replace('_agent', '').capitalize(),
                    'Lifespan': stats['livespan'],
                    'Average Food': stats['average_food'],
                    'Average Wood': stats['average_wood'],
                    'Average Water': stats['average_water']
                }, ignore_index=True)

# Load data from tribenum20.json
with open('all_results_all_pasture.json', 'r') as file:
    content = file.read()
    json_objects = content.split("\n")
    for json_str in json_objects:
        if json_str:
            data = json.loads(json_str)
            for role, stats in data.items():
                data_tribeallpasture = data_tribeallpasture.append({
                    'Role': role.replace('_agent', '').capitalize(),
                    'Lifespan': stats['livespan'],
                    'Average Food': stats['average_food'],
                    'Average Wood': stats['average_wood'],
                    'Average Water': stats['average_water']
                }, ignore_index=True)

with open('all_results_all_forest.json', 'r') as file:
    content = file.read()
    json_objects = content.split("\n")
    for json_str in json_objects:
        if json_str:
            data = json.loads(json_str)
            for role, stats in data.items():
                data_tribeallforest = data_tribeallforest.append({
                    'Role': role.replace('_agent', '').capitalize(),
                    'Lifespan': stats['livespan'],
                    'Average Food': stats['average_food'],
                    'Average Wood': stats['average_wood'],
                    'Average Water': stats['average_water']
                }, ignore_index=True)

import numpy as np

average_data_tribenormal = data_tribenormal.groupby('Role').mean().reset_index()
average_data_tribeallpasture = data_tribeallpasture.groupby('Role').mean().reset_index()
average_data_tribeallforest = data_tribeallforest.groupby('Role').mean().reset_index()
# Define bar width
bar_width = 0.35

# Define positions of bars for tribe size 10 and tribe number 20
r1 = np.arange(len(average_data_tribenormal))
r2 = [x + bar_width for x in r1]
r3 = [x + bar_width for x in r2]

# Plot averages for each role
plt.figure(figsize=(10, 6))

def plot_data(title, y_label, data_key):
    plt.figure(figsize=(10, 6))

    plt.bar(r1, average_data_tribenormal[data_key], color='b', width=bar_width, edgecolor='grey', label='Normal')
    plt.bar(r2, average_data_tribeallpasture[data_key], color='g', width=bar_width, edgecolor='grey', label='All Pasture')
    plt.bar(r3, average_data_tribeallforest[data_key], color='r', width=bar_width, edgecolor='grey', label='All Forest')

    plt.title(title)
    plt.xlabel('Role')
    plt.ylabel(y_label)
    plt.xticks([r + bar_width / 2 for r in range(len(average_data_tribenormal))], average_data_tribenormal['Role'])
    plt.legend()
    plt.show()

# Now you can simply call the function with the appropriate parameters
plot_data('Average Lifespan per Role for Normal and All pasture and forest Environments', 'Average Lifespan', 'Lifespan')
plot_data('Average Food Level per Role for Normal and All Pasture and Forest Environments', 'Average Food Level', 'Average Food')
plot_data('Average Wood Level per Role for Normal and All Pasture and Forest Environments', 'Average Wood Level', 'Average Wood')
plot_data('Average Water Level per Role for Normal and All Pasture and Forest Environments', 'Average Water Level', 'Average Water')