import pygame
import sys
from Visual.ui_manager import UIManager
from Visual.GameState import GameState
from Visual.event_handler import EventHandler

def main():
    pygame.init()
    screen = pygame.display.set_mode((800, 600))
    pygame.display.set_caption("Simulation Visualizer")
    
    running = True
    clock = pygame.time.Clock()

    game_state = GameState()
    ui_manager = UIManager()
    event_handler = EventHandler()

    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            event_handler.handle_event(event, game_state, ui_manager)
        
        screen.fill((255, 255, 255))  # Clear the screen with white
        game_state.update()
        game_state.draw(screen)
        ui_manager.draw(screen)

        pygame.display.flip()  # Update the display
        clock.tick(60)  # Maintain 60 frames per second

    pygame.quit()
    sys.exit()

if __name__ == "__main__":
    main()
