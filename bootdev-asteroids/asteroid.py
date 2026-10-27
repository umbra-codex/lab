import pygame
import random
from circle_shape import CircleShape
from logger import log_event
from constants import ASTEROID_MIN_RADIUS, LINE_WIDTH

class Asteroid(CircleShape):
  def __init__(self, x, y, radius):
    super().__init__(x, y, radius)

  def draw(self, screen):
    pygame.draw.circle(screen, "white", self.position, self.radius, LINE_WIDTH)

  def update(self, dt):
    self.position += self.velocity * dt

  def split(self):
    self.kill()

    if self.radius <= ASTEROID_MIN_RADIUS:
      return
    
    log_event("asteroid_split")

    angle = random.uniform(20, 50)
    left_velocity = self.velocity.rotate(angle)
    right_velocity = self.velocity.rotate(-angle)
    new_radius = self.radius - ASTEROID_MIN_RADIUS

    first_asteroid_split = Asteroid(self.position.x, self.position.y, new_radius)
    first_asteroid_split.velocity = left_velocity * 1.2
    second_asteroid_split = Asteroid(self.position.x, self.position.y, new_radius)
    second_asteroid_split.velocity = right_velocity * 1.2

