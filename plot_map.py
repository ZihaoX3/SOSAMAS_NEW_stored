# import matplotlib.pyplot as plt
# import networkx as nx

# # Defining the nodes (areas) and their attributes
# nodes = {
#     0: {"type": "Forest", "water": 0.45, "food": 0.46, "wood": 0.33},
#     1: {"type": "Desert", "water": 0.01, "food": 0.01, "wood": 0.04},
#     2: {"type": "Forest", "water": 0.49, "food": 0.46, "wood": 0.18},
#     3: {"type": "Forest", "water": 0.49, "food": 0.42, "wood": 0.28},
#     4: {"type": "Forest", "water": 0.47, "food": 0.46, "wood": 0.22},
#     5: {"type": "Forest", "water": 0.39, "food": 0.42, "wood": 0.27},
#     6: {"type": "Pasture", "water": 0.26, "food": 0.20, "wood": 0.11},
#     7: {"type": "Forest", "water": 0.42, "food": 0.48, "wood": 0.24},
#     8: {"type": "Desert", "water": 0.02, "food": 0.02, "wood": 0.10},
#     9: {"type": "Forest", "water": 0.40, "food": 0.48, "wood": 0.22}
# }

# # Defining the edges (connections between areas)
# edges = [
#     (0, 3), (0, 9), (0, 6), (0, 4),
#     (1, 8), (1, 7), (1, 3),
#     (2, 8), (2, 4), (2, 5),
#     (3, 1), (3, 6), (3, 4), (3, 7),
#     (4, 2), (4, 3), (4, 0), (4, 9),
#     (5, 2), (5, 8), (5, 6), (5, 7),
#     (6, 0), (6, 3), (6, 5), (6, 9),
#     (7, 1), (7, 5), (7, 3),
#     (8, 1), (8, 2), (8, 5),
#     (9, 0), (9, 4), (9, 6)
# ]

# # Creating the graph
# G = nx.Graph()
# G.add_nodes_from(nodes)
# G.add_edges_from(edges)

# # Assigning colors based on environment types
# color_map = []
# for node in G:
#     if nodes[node]["type"] == "Forest":
#         color_map.append("green")
#     elif nodes[node]["type"] == "Desert":
#         color_map.append("sandybrown")
#     elif nodes[node]["type"] == "Pasture":
#         color_map.append("lightgreen")
#     else:
#         color_map.append("gray")

# # Drawing the graph with legends
# plt.figure(figsize=(12, 8))
# nx.draw(G, with_labels=True, node_color=color_map, node_size=2000, font_size=10, font_weight='bold')

# # Creating a legend for the graph
# forest_patch = plt.Line2D([0], [0], marker='o', color='w', label='Forest',
#                           markersize=10, markerfacecolor='green')
# desert_patch = plt.Line2D([0], [0], marker='o', color='w', label='Desert',
#                           markersize=10, markerfacecolor='sandybrown')
# pasture_patch = plt.Line2D([0], [0], marker='o', color='w', label='Pasture',
#                            markersize=10, markerfacecolor='lightgreen')

# plt.legend(handles=[forest_patch, desert_patch, pasture_patch], title="Environment Types")
# plt.title("Simulation Area Map")
# plt.show()


import matplotlib.pyplot as plt
import networkx as nx

# Defining the nodes (areas) and their attributes, without nodes 8 and 9
nodes = {
    0: {"type": "Forest", "water": 0.45, "food": 0.46, "wood": 0.33},
    1: {"type": "Desert", "water": 0.01, "food": 0.01, "wood": 0.04},
    2: {"type": "Forest", "water": 0.49, "food": 0.46, "wood": 0.18},
    3: {"type": "Forest", "water": 0.49, "food": 0.42, "wood": 0.28},
    4: {"type": "Forest", "water": 0.47, "food": 0.46, "wood": 0.22},
    5: {"type": "Forest", "water": 0.39, "food": 0.42, "wood": 0.27},
    6: {"type": "Pasture", "water": 0.26, "food": 0.20, "wood": 0.11},
    7: {"type": "Forest", "water": 0.42, "food": 0.48, "wood": 0.24}
}

# Defining the edges (connections between areas), without connections involving nodes 8 and 9
edges = [
    (0, 3), (0, 6), (0, 4),
    (1, 7), (1, 3),
    (2, 4), (2, 5),
    (3, 1), (3, 6), (3, 4), (3, 7),
    (4, 2), (4, 3), (4, 0),
    (5, 2), (5, 6), (5, 7),
    (6, 0), (6, 3), (6, 5),
    (7, 1), (7, 5), (7, 3)
]

# Creating the graph
G = nx.Graph()
G.add_nodes_from(nodes.items())
G.add_edges_from(edges)

# Assigning colors based on environment types
color_map = [nodes[node]["type"] == "Forest" and "green" or
             nodes[node]["type"] == "Desert" and "sandybrown" or
             nodes[node]["type"] == "Pasture" and "lightgreen" for node in G]

# Drawing the graph with legends
plt.figure(figsize=(12, 8))
pos = nx.spring_layout(G, seed=42)  # For consistent layout
nx.draw(G, pos, with_labels=True, node_color=color_map, node_size=2000, font_size=15, font_weight='bold')

# Creating a legend for the graph
forest_patch = plt.Line2D([0], [0], marker='o', color='w', label='Forest',
                          markersize=12, markerfacecolor='green')
desert_patch = plt.Line2D([0], [0], marker='o', color='w', label='Desert',
                          markersize=12, markerfacecolor='sandybrown')
pasture_patch = plt.Line2D([0], [0], marker='o', color='w', label='Pasture',
                           markersize=12, markerfacecolor='lightgreen')

plt.legend(handles=[forest_patch, desert_patch, pasture_patch], title="Environment Types", loc='upper left')
plt.title("Simulation Area Map")
plt.show()
