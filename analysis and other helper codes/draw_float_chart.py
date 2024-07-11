from graphviz import Digraph

dot = Digraph()

# Define the nodes and their styles
dot.node('A', 'Start Simulation Cycle')
dot.node('B', 'Update Resource Levels (Seasonal)')
dot.node('C', 'Health Check (Each Member)')
dot.node('D', 'Decision Making')
dot.node('E', 'Execute Actions')
dot.node('F', 'State Update')
dot.node('G', 'Data Collection')
dot.node('H', 'End of Cycle')

# Add edges between the nodes to show the flow
dot.edge('A', 'B', 'Initialize resources')
dot.edge('B', 'C', 'Gather resources')
dot.edge('C', 'D', 'Check survival status')
dot.edge('D', 'E', 'Determine actions')
dot.edge('E', 'F', 'Perform actions')
dot.edge('F', 'G', 'Update states')
dot.edge('G', 'H', 'Collect data')
dot.edge('H', 'A', 'Cycle complete, start new')

# Optionally, add sub-processes or conditions if needed
dot.edge('D', 'B', 'Update seasonal effects', constraint='false')  # Example of a feedback loop

# Render and view the graph
dot.render('simulation_cycle_flowchart', view=True)
