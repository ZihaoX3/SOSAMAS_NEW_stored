# SOSAMAS_FYP
Multi-Agent Simulation for Self-Organised Social Arrangements
## Project Overview

This project simulates the adaptive social structures and behaviors of North American indigenous tribes using Self-Organizing Multi-Agent Systems (SOMAS). The aim is to explore decentralized decision-making processes, resource management strategies, and social adaptability within a dynamic environment.

## Features

- Simulation of various geographical landscapes (forest, desert, pasture)
- Autonomous agent behaviors including foraging, farming, and banditry
- Resource management and optimization
- Adaptive decision-making based on historical satisfaction
- Relocation and shelter-building strategies
- Birth and death rate simulation to mimic real-life population dynamics

## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/somas-simulation.git
    cd somas-simulation
    ```

2. **Install dependencies:**
    Ensure you have Go and Python installed on your system. Then, install the required Python packages:
    ```sh
    pip install -r requirements.txt
    ```

## Usage

1. **Run the simulation:**
    ```sh
    go run main.go
    ```

2. **Analyze simulation data:**
    ```sh
    python analysis_json.py
    ```

3. **Visualization:**
    The simulation results can be visualized using the provided scripts. Ensure you have the necessary plotting libraries installed (e.g., matplotlib).

## Configuration

You can adjust various simulation parameters in the `config.json` file to customize the simulation environment and agent behaviors.

## Example

Here's an example of running a simulation and analyzing the results:
```sh
go run main.go
python analysis_json.py

