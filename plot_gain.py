import pandas as pd
import matplotlib.pyplot as plt

# Assuming the data is stored in a CSV file named 'simulation_data.csv'
# and it has columns: 'Round', 'EnvironmentalWater', 'Food', 'Wood'

# Read the CSV file
df = pd.read_csv('simulation_gains.csv')

# Plotting all three resources on the same graph
plt.figure(figsize=(12, 8))
plt.plot(df['Round'], df['WaterGain'], label='Environmental Water gain', color='blue')
plt.plot(df['Round'], df['FoodGain'], label='Environmental Food gain', color='green')
plt.plot(df['Round'], df['WoodGain'], label='Environmental Wood gain', color='brown')

# Adding labels and title
plt.xlabel('Round')
plt.ylabel('Resource Level')
plt.title('Forest Resource Levels Over Time')
plt.legend()

# Display the plot
plt.show()

