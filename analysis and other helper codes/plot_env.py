import pandas as pd
import matplotlib.pyplot as plt

# Assuming the data is stored in a CSV file named 'simulation_data.csv'
# and it has columns: 'Round', 'EnvironmentalWater', 'Food', 'Wood'

# Read the CSV file
df = pd.read_csv('simulation_data.csv')

# Plotting all three resources on the same graph
plt.figure(figsize=(12, 8))
plt.plot(df['Round'], df['EnvironmentalWater'], label='Environmental Water', color='blue')
plt.plot(df['Round'], df['Food'], label='Food', color='green')
plt.plot(df['Round'], df['Wood'], label='Wood', color='brown')

# Adding labels and title
plt.xlabel('Round')
plt.ylabel('Resource Level')
plt.title('Desert Resource Levels Over Time')
plt.legend()

# Display the plot
plt.show()

