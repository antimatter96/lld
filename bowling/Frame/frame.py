from Roll.roll import Roll

class Frame:
    def __init__(self, first_roll, second_roll = Roll()):
        self._first_roll = first_roll
        self._second_roll = second_roll
    
    def no_of_knocked_pins(self):
        return self._first_roll.no_of_knocked_pins() + self._second_roll.no_of_knocked_pins()
    
    def no_of_pins_knocked_in_first_roll(self):
        return self._first_roll.no_of_knocked_pins()

    def is_strike(self):
        return False

    def is_spare(self):
        return False


class Strike(Frame):
    def __init__(self, first_roll):
        super().__init__(first_roll)
    
    def is_strike(self):
        return True

class Spare(Frame):
    def __init__(self, first_roll, second_roll):
        super().__init__(first_roll, second_roll)
    
    def is_spare(self):
        return True
    

