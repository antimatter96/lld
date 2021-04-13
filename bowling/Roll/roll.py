class Roll:
    def __init__(self, no_of_knocked_pins = 0):
        self._no_of_knocked_pins = no_of_knocked_pins
    
    def reset_roll(self):
        self._no_of_knocked_pins = 0
    
    def no_of_knocked_pins(self):
        return self._no_of_knocked_pins
