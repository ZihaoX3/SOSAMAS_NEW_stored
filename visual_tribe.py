import pygame
import json
import math

# Initialize Pygame
pygame.init()

# Set up the display
screen_width, screen_height = 1100, 700  # Increased window size
screen = pygame.display.set_mode((screen_width, screen_height))
pygame.display.set_caption("Simulation Visualizer")

# Load JSON data for simulation states and area configurations
with open('simulation_data.json', 'r') as file:
    simulation_data = json.load(file)

with open('area_states.json', 'r') as file:
    area_data = json.load(file)

# Define some colors
background_color = pygame.Color('white')
tribe_colors = {
    'Bandit_Agent': pygame.Color('red'),
    'Farmer_agent': pygame.Color('yellow'),
    'Forager_agent': pygame.Color('cyan'),
}
forest_color = pygame.Color('green')
pasture_color = pygame.Color('blue')
desert_color = pygame.Color('sandybrown')
text_color = pygame.Color('black')
line_color = pygame.Color('gray')
success_color = pygame.Color('green')
failure_color = pygame.Color('pink')
relocation_color = pygame.Color('orange')
boat_color = pygame.Color('black')

# Set up the clock for FPS control
clock = pygame.time.Clock()
map_radius = min(screen_width, screen_height) * 0.3

def draw_arrow(screen, color, start, end, width=3, arrow_size=8):
    """Draw an arrow from start to end with the given color and width."""
    pygame.draw.line(screen, color, start, end, width)
    rotation = math.degrees(math.atan2(start[1] - end[1], end[0] - start[0])) + 90
    pygame.draw.polygon(screen, color, (
        (end[0] + arrow_size * math.sin(math.radians(rotation)), end[1] + arrow_size * math.cos(math.radians(rotation))),
        (end[0] + arrow_size * math.sin(math.radians(rotation - 120)), end[1] + arrow_size * math.cos(math.radians(rotation - 120))),
        (end[0] + arrow_size * math.sin(math.radians(rotation + 120)), end[1] + arrow_size * math.cos(math.radians(rotation + 120)))
    ))

def draw_marker(screen, color, position, success=True, shape='check'):
    """Draw a marker at the given position to indicate success, failure, or relocation."""
    if shape == 'check':
        if success:
            # Draw a green checkmark
            pygame.draw.line(screen, color, (position[0] - 5, position[1]), (position[0], position[1] + 5), 3)
            pygame.draw.line(screen, color, (position[0], position[1] + 5), (position[0] + 10, position[1] - 5), 3)
        else:
            # Draw a red 'X'
            pygame.draw.line(screen, color, (position[0] - 5, position[1] - 5), (position[0] + 5, position[1] + 5), 3)
            pygame.draw.line(screen, color, (position[0] + 5, position[1] - 5), (position[0] - 5, position[1] + 5), 3)
    elif shape == 'relocation':
        # Draw an orange circle for relocation
        pygame.draw.circle(screen, color, position, 10, 2)


def draw_boat_marker(screen, position):
    """Draw a black mark under the tribe's circle to indicate a boat."""
    pygame.draw.line(screen, boat_color, (position[0] - 5, position[1] + 8), (position[0] + 5, position[1] + 8), 3)

def draw_shelter_markers(screen, position, shelter_level):
    """Draw small marks next to the tribe's circle to indicate shelter levels."""
    for i in range(shelter_level):
        pygame.draw.line(screen, pygame.Color('brown'), (position[0] + 8 + i * 5, position[1] - 5), (position[0] + 8 + i * 5, position[1] + 5), 2)


def draw_map():
    font = pygame.font.Font(None, 20)
    center_x, center_y = screen_width // 2, screen_height // 2
    num_areas = len(area_data)
    angle_between_areas = 2 * math.pi / num_areas  # Angle between each area

    # Draw connections first
    for area in area_data:
        for neighbor_id in area['neighbors']:
            neighbor = next((a for a in area_data if a['id'] == neighbor_id), None)
            if neighbor:
                start_angle = angle_between_areas * area['id']
                end_angle = angle_between_areas * neighbor_id
                start_pos = (center_x + map_radius * math.cos(start_angle), center_y + map_radius * math.sin(start_angle))
                end_pos = (center_x + map_radius * math.cos(end_angle), center_y + map_radius * math.sin(end_angle))
                pygame.draw.line(screen, line_color, start_pos, end_pos, 2)

    # Draw areas
    for area in area_data:
        angle = angle_between_areas * area['id']
        x = center_x + map_radius * math.cos(angle)
        y = center_y + map_radius * math.sin(angle)
        color = forest_color if area['environment'] == 'Forest' else pasture_color if area['environment'] == 'Pasture' else desert_color
        pygame.draw.circle(screen, color, (int(x), int(y)), 60, 2)  # Draw border only
        text_surface = font.render(f"Area {area['id']}", True, text_color)
        screen.blit(text_surface, (int(x) - text_surface.get_width() // 2, int(y) - text_surface.get_height() // 2))

def draw_tribes(screen, iteration, area_data, tribe_colors):
    font = pygame.font.Font(None, 20)
    center_x, center_y = screen_width // 2, screen_height // 2
    map_radius = 0.3 * min(screen_width, screen_height)
    area_radius = 30  # Radius of the area circle
    member_radius = 5  # Radius for each tribe member for better visibility
    angle_between_areas = 2 * math.pi / len(area_data)

    for area in area_data:
        base_angle = angle_between_areas * area['id']
        area_center_x = center_x + map_radius * math.cos(base_angle)
        area_center_y = center_y + map_radius * math.sin(base_angle)

        tribes_in_area = [t for t in iteration['tribes'] if t['area_id'][0] == area['id']]
        num_tribes_in_area = len(tribes_in_area)
        angle_increment = 2 * math.pi / max(1, num_tribes_in_area)

        for index, tribe_in_area in enumerate(tribes_in_area):
            tribe_angle = base_angle + index * angle_increment
            x = area_center_x + area_radius * math.cos(tribe_angle)
            y = area_center_y + area_radius * math.sin(tribe_angle)
            tribe_color = tribe_colors.get(tribe_in_area['role'][0], pygame.Color('white'))
            pygame.draw.circle(screen, tribe_color, (int(x), int(y)), member_radius)

            if tribe_in_area['is_attacking'] and 'target_id' in tribe_in_area:
                target_id = int(tribe_in_area['target_id'])
                target_tribe = next((t for t in iteration['tribes'] if t['id'] == str(target_id)), None)
                if target_tribe and target_tribe['area_id']:
                    target_area_id = target_tribe['area_id'][0]
                    target_area = next((a for a in area_data if a['id'] == target_area_id), None)
                    if target_area:
                        target_angle = angle_between_areas * target_area_id
                        target_x = center_x + map_radius * math.cos(target_angle)
                        target_y = center_y + map_radius * math.sin(target_angle)
                        draw_arrow(screen, pygame.Color('red'), (int(x), int(y)), (int(target_x), int(target_y)), arrow_size=10)
                        
                        # Calculate midpoint of the arrow
                        mid_x = (x + target_x) // 2
                        mid_y = (y + target_y) // 2
                        
                        success = tribe_in_area.get('attack_success_check', False)
                        draw_marker(screen, success_color if success else failure_color, (int(mid_x), int(mid_y)), success=success)

            if tribe_in_area.get('relocated'):
                prev_area_id = tribe_in_area.get('previous_location', None)
                if prev_area_id is not None:
                    prev_area = next((a for a in area_data if a['id'] == prev_area_id), None)
                    if prev_area:
                        prev_angle = angle_between_areas * prev_area_id
                        prev_x = center_x + map_radius * math.cos(prev_angle)
                        prev_y = center_y + map_radius * math.sin(prev_angle)
                        draw_arrow(screen, relocation_color, (int(prev_x), int(prev_y)), (int(x), int(y)), arrow_size=5)
            if tribe_in_area['has_boat']:
                draw_boat_marker(screen, (int(x), int(y)))

            if tribe_in_area['shelter_level'] > 0:
                draw_shelter_markers(screen, (int(x), int(y)), tribe_in_area['shelter_level'])

                
def draw_text(surface, text, position, font_size=20, text_color=pygame.Color('black')):
    font = pygame.font.Font(None, font_size)
    text_surface = font.render(text, True, text_color)
    rect = text_surface.get_rect(topleft=position)
    surface.blit(text_surface, rect)

def draw_tribe_info(screen, iteration, tribe_colors):
    font = pygame.font.Font(None, 15)
    y_offset = 10
    for tribe in iteration['tribes']:
        if tribe['members'] > 0:
            tribe_color = tribe_colors.get(tribe['role'][0], text_color)
            if tribe['has_boat']:
                if tribe['shelter_level'] > 0:
                    tribe_info = f"ID {tribe['id']}: area: {tribe['area_id']}, num: {tribe['members']}, Food: {tribe['food_level']:.2f}, Wood: {tribe['wood_level']:.2f}, Water: {tribe['water_level']:.2f}, {', '.join(tribe['role'])}, Shelter:{tribe['shelter_level']:.2f}, Boat: Yes"
                else:
                    tribe_info = f"ID {tribe['id']}: area: {tribe['area_id']}, num: {tribe['members']}, Food: {tribe['food_level']:.2f}, Wood: {tribe['wood_level']:.2f}, Water: {tribe['water_level']:.2f}, {', '.join(tribe['role'])}, Boat: Yes"
            else:
                if tribe['shelter_level'] > 0:
                    tribe_info = f"ID {tribe['id']}: area: {tribe['area_id']}, num: {tribe['members']}, Food: {tribe['food_level']:.2f}, Wood: {tribe['wood_level']:.2f}, Water: {tribe['water_level']:.2f}, {', '.join(tribe['role'])}, Shelter:{tribe['shelter_level']:.2f}"
                else:
                    tribe_info = f"ID {tribe['id']}: area: {tribe['area_id']}, num: {tribe['members']}, Food: {tribe['food_level']:.2f}, Wood: {tribe['wood_level']:.2f}, Water: {tribe['water_level']:.2f}, {', '.join(tribe['role'])}"
            text_surface = font.render(tribe_info, True, tribe_color)
            screen.blit(text_surface, (20, y_offset))
            y_offset += 10

def draw_events(screen, events, font_size=15, text_color=pygame.Color('black')):
    font = pygame.font.Font(None, font_size)
    y_offset = screen_height - 500  # Start drawing from the bottom
    for event in reversed(events):  # Draw from the newest to the oldest
        text_surface = font.render(event, True, text_color)
        screen.blit(text_surface, (screen_width - 280, y_offset))
        y_offset -= 10

def main_game_loop():
    running = True
    iteration_index = 0
    total_iterations = len(simulation_data)
    events = []
    added_events = set()  # Track added events to prevent duplicates
    paused = False  # Flag to control pause/resume

    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            elif event.type == pygame.USEREVENT + 1 and not paused:
                iteration_index = (iteration_index + 1) % total_iterations
            elif event.type == pygame.KEYDOWN:
                if event.key == pygame.K_SPACE:
                    paused = not paused  # Toggle pause/resume

        screen.fill(background_color)
        draw_map()
        draw_tribes(screen, simulation_data[iteration_index], area_data, tribe_colors)
        draw_text(screen, f"Iteration: {iteration_index + 1} / {total_iterations}", (50, 200))

        draw_tribe_info(screen, simulation_data[iteration_index], tribe_colors)

        current_iteration = simulation_data[iteration_index]
        round_events = []  # Use a list to store unique events for the current round
        if not paused:
            for tribe in current_iteration['tribes']:
                if tribe.get('boat_building_time') == iteration_index:
                    event_text = f"Round {iteration_index + 1}: Tribe {tribe['id']} built a boat."
                    if event_text not in added_events:
                        round_events.append(event_text)
                        added_events.add(event_text)
                if tribe.get('relocated'):
                    event_text = f"Round {iteration_index + 1}: Tribe {tribe['id']} relocated to {tribe['area_id']}."
                    if event_text not in added_events:
                        round_events.append(event_text)
                        added_events.add(event_text)
                if tribe.get('shelter_building'):
                    event_text = f"Round {iteration_index + 1}: Tribe {tribe['id']} built a shelter."
                    if event_text not in added_events:
                        round_events.append(event_text)
                        added_events.add(event_text)
                if tribe['is_attacking']:
                    target_id = tribe['target_id']
                    if tribe.get('attack_success_check', False):
                        event_text = f"Round {iteration_index + 1}: Tribe {tribe['id']} successfully attacked Tribe {target_id}."
                    else:
                        event_text = f"Round {iteration_index + 1}: Tribe {tribe['id']} failed to attack Tribe {target_id}."
                    if event_text not in added_events:
                        round_events.append(event_text)
                        added_events.add(event_text)

            if round_events:
                events.extend(round_events)
                events = events[-40:]  # Keep only the last 40 events for display

            # Debugging output to check for duplicates
            print(f"Iteration {iteration_index + 1}, events: {len(events)}")
            for e in events:
                print(e)

        draw_events(screen, events)

        pygame.display.flip()
        clock.tick(60)

    pygame.quit()

if __name__ == "__main__":
    pygame.time.set_timer(pygame.USEREVENT + 1, 1500)
    main_game_loop()
