# import json
# import pandas as pd
# import matplotlib.pyplot as plt

# # Load the JSON data
# with open('all_results_onlybandit.json', 'r') as f:
#     content = f.read()
#     content = '[' + content.replace('}{', '},{') + ']'
#     data_bandit = json.loads(content)

# with open('all_results_greedy0.json') as f:
#     data_normal = [json.loads(line) for line in f]


# # Convert the data to pandas DataFrames
# df_normal = pd.DataFrame(data_normal)
# df_bandit = pd.DataFrame(data_bandit)

# # Calculate the averages for the 'Bandit_Agent' role
# averages_normal = df_normal['Bandit_Agent'].apply(pd.Series).mean()
# averages_bandit = df_bandit['Bandit_Agent'].apply(pd.Series).mean()

# # Create a new DataFrame with the averages
# df_averages = pd.DataFrame({'Normal': averages_normal, 'All Bandit': averages_bandit})

# # Create a bar chart for 'livespan'
# df_averages.loc['livespan'].transpose().plot(kind='bar', color=['blue', 'green'], title='livespan')
# plt.xticks(rotation=0)
# plt.title('Difference of lifespan for Bandit between Normal and All Bandit simulation')
# plt.ylabel('Average Lifespan')
# plt.show()

# # Create a bar chart for the other properties
# df_averages.loc[['average_food', 'average_water', 'average_wood']].transpose().plot(kind='bar', color=['blue', 'green', 'red'])
# plt.xticks(rotation=0)
# plt.title('Resources level for Bandit between Normal and All Bandit simulation')
# plt.ylabel('Average Resource level')
# plt.show()


# import json
# import pandas as pd
# import matplotlib.pyplot as plt

# # Load the JSON data
# with open('all_results_onlyforager.json', 'r') as f:
#     content = f.read()
#     content = '[' + content.replace('}{', '},{') + ']'
#     data_forager = json.loads(content)

# with open('all_results_greedy0.json') as f:
#     data_normal = [json.loads(line) for line in f]


# # Convert the data to pandas DataFrames
# df_normal = pd.DataFrame(data_normal)
# df_bandit = pd.DataFrame(data_forager)

# # Calculate the averages for the 'Bandit_Agent' role
# averages_normal = df_normal['Forager_agent'].apply(pd.Series).mean()
# averages_forager = df_bandit['Forager_agent'].apply(pd.Series).mean()

# # Create a new DataFrame with the averages
# df_averages = pd.DataFrame({'Normal': averages_normal, 'All Forager': averages_forager})

# # Create a bar chart for 'livespan'
# df_averages.loc['livespan'].transpose().plot(kind='bar', color=['blue', 'green'], title='livespan')
# plt.xticks(rotation=0)
# plt.title('Difference of lifespan for Foragers between Normal and All Foragers simulation')
# plt.ylabel('Average Lifespan')
# plt.show()

# # Create a bar chart for the other properties
# df_averages.loc[['average_food', 'average_water', 'average_wood']].transpose().plot(kind='bar', color=['blue', 'green', 'red'])
# plt.xticks(rotation=0)
# plt.title('Resources level for Foragers between Normal and All Foragers simulation')
# plt.ylabel('Average Resource level')
# plt.show()

import json
import pandas as pd
import matplotlib.pyplot as plt

# Load the JSON data
with open('all_results_onlyfarmer.json', 'r') as f:
    content = f.read()
    content = '[' + content.replace('}{', '},{') + ']'
    data_farmer = json.loads(content)

with open('all_results_greedy0.json') as f:
    data_normal = [json.loads(line) for line in f]


# Convert the data to pandas DataFrames
df_normal = pd.DataFrame(data_normal)
df_bandit = pd.DataFrame(data_farmer)

# Calculate the averages for the 'Bandit_Agent' role
averages_normal = df_normal['Farmer_agent'].apply(pd.Series).mean()
averages_farmer = df_bandit['Farmer_agent'].apply(pd.Series).mean()

# Create a new DataFrame with the averages
df_averages = pd.DataFrame({'Normal': averages_normal, 'All Farmers': averages_farmer})

# Create a bar chart for 'livespan'
df_averages.loc['livespan'].transpose().plot(kind='bar', color=['blue', 'green'], title='livespan')
plt.xticks(rotation=0)
plt.title('Difference of lifespan for Farmers between Normal and All Farmers simulation')
plt.ylabel('Average Lifespan')
plt.show()

# Create a bar chart for the other properties
df_averages.loc[['average_food', 'average_water', 'average_wood']].transpose().plot(kind='bar', color=['blue', 'green', 'red'])
plt.xticks(rotation=0)
plt.title('Resources level for Farmers between Normal and All Farmers simulation')
plt.ylabel('Average Resource level')
plt.show()