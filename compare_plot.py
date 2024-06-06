# import matplotlib.pyplot as plt
# import json

# def load_results(filename):
#     with open(filename, 'r') as f:
#         data = [json.loads(line) for line in f]
#     return data

# def extract_metrics(data, role):
#     metrics = {'lifespan': [], 'average_food': [], 'average_wood': [], 'average_water': []}
#     for entry in data:
#         if role in entry:
#             metrics['lifespan'].append(entry[role]['livespan'])
#             metrics['average_food'].append(entry[role]['average_food'])
#             metrics['average_wood'].append(entry[role]['average_wood'])
#             metrics['average_water'].append(entry[role]['average_water'])
#     return metrics

# def plot_metrics(original_metrics, modified_metrics, role, color1, color2):
#     fig, axs = plt.subplots(4, 1, figsize=(10, 20), sharex=True)
#     fig.suptitle(f'Comparison of {role} Role Across Simulations', fontsize=16)
    
#     titles = ['Lifespan', 'Average Food Level', 'Average Wood Level', 'Average Water Level']
#     metrics_keys = ['lifespan', 'average_food', 'average_wood', 'average_water']
    
#     for i, ax in enumerate(axs.flatten()):
#         ax.plot(original_metrics[metrics_keys[i]], color=color1, label='Original')
#         ax.plot(modified_metrics[metrics_keys[i]], color=color2, label='Modified')
#         ax.set_title(titles[i])
#         ax.set_ylabel(titles[i])
#         ax.grid(True)
#         ax.legend()
    
#     plt.xlabel('Simulation Run')
#     plt.tight_layout(rect=[0, 0.03, 1, 0.95])
#     plt.show()

# # Load the results from both datasets
# original_data = load_results('all_results_normal.json')
# modified_data = load_results('all_results_tribenum20.json')

# # Specify the roles you want to analyze
# roles = ['Bandit_Agent', 'Farmer_agent', 'Forager_agent']
# colors = [('blue', 'red'), ('green', 'orange'), ('purple', 'yellow')]

# # Loop through each role and plot the metrics
# for role, (color1, color2) in zip(roles, colors):
#     original_metrics = extract_metrics(original_data, role)
#     modified_metrics = extract_metrics(modified_data, role)
#     plot_metrics(original_metrics, modified_metrics, role, color1, color2)


import json
import matplotlib.pyplot as plt
import numpy as np

# Function to plot data
def plot_with_difference(original, modified, title, ylabel):
    # Calculate the difference
    difference = np.array(modified) - np.array(original)
    
    # Apply a simple moving average for smoothing
    window_size = 10  # Size of the moving average window
    smooth_difference = np.convolve(difference, np.ones(window_size)/window_size, mode='valid')
    
    # Create figure and plot space
    plt.figure(figsize=(12, 6))
    
    # Plot original and modified data
    plt.plot(original, label='Original', color='blue')
    plt.plot(modified, label='Modified', color='red')
    
    # Plot smoothed difference
    # Adjust the x-values for the smoothed line for proper alignment
    plt.plot(range(window_size // 2, len(smooth_difference) + window_size // 2), smooth_difference, label='Difference (smoothed)', color='green', linestyle='--')
    
    # Adding titles and labels
    plt.title(title)
    plt.xlabel('Simulation Run')
    plt.ylabel(ylabel)
    plt.legend()
    plt.grid(True)
    plt.show()

# Load the original and modified results
with open('all_results_greedy40.json') as f_original, open('all_results_nogreedy.json') as f_modified:
    original_data = [json.loads(line) for line in f_original]
    modified_data = [json.loads(line) for line in f_modified]

# Prepare lists for different metrics
original_lifespan = [data['Forager_agent']['livespan'] for data in original_data]
modified_lifespan = [data['Forager_agent']['livespan'] for data in modified_data]

original_food = [data['Forager_agent']['average_food'] for data in original_data]
modified_food = [data['Forager_agent']['average_food'] for data in modified_data]

original_water = [data['Forager_agent']['average_water'] for data in original_data]
modified_water = [data['Forager_agent']['average_water'] for data in modified_data]

original_wood = [data['Forager_agent']['average_wood'] for data in original_data]
modified_wood = [data['Forager_agent']['average_wood'] for data in modified_data]

# Plot each metric
plot_with_difference(original_lifespan, modified_lifespan, 'Comparison of Lifespan Across Simulations for Bandit', 'Lifespan')
plot_with_difference(original_food, modified_food, 'Comparison of Average Food Level Across Simulations for Bandit', 'Average Food Level')
plot_with_difference(original_water, modified_water, 'Comparison of Average Water Level Across Simulations for Bandit', 'Average Water Level')
plot_with_difference(original_wood, modified_wood, 'Comparison of Average Wood Level Across Simulations for Bandit', 'Average Wood Level')
