from constants import MISS_SYMBOL, SPARE_SYMBOL, STRIKE_SYMBOL
from Roll.roll import Roll

class Roll_Parser:
    def parse(self, string_roll):
        is_valid = self.validate_input(string_roll)
        if not is_valid:
            raise Exception('Input Length is not Valid')
        rolls = []
        for idx in range(len(string_roll)):
            rolls.append(self.create_roll(string_roll, idx))
        return rolls

    def validate_input(self, string_roll):
        no_of_strikes = string_roll.count(STRIKE_SYMBOL)
        no_of_chars = len(string_roll) + no_of_strikes

        if string_roll[-3] == STRIKE_SYMBOL:

            if string_roll[-2:] == STRIKE_SYMBOL + STRIKE_SYMBOL:
                return no_of_chars ==  24

            elif string_roll[-2] == STRIKE_SYMBOL  or string_roll[-1] == STRIKE_SYMBOL:
                return no_of_chars == 23

            else:
                return no_of_chars == 22 # In Last Frame scored STRIKE , and subsequent were not STRIKE

        elif  string_roll[-2] == SPARE_SYMBOL:
            return no_of_chars == 21

        else:
            return no_of_chars == 20 # Since we have 2 rolls per frame
        
    def create_roll(self, string_roll, idx):
        no_of_knocked_pins = self.get_knocked_pins_of(string_roll, idx)
        return Roll(no_of_knocked_pins)
    
    def get_knocked_pins_of(self,string_roll, idx):
        char_to_parse = string_roll[idx]
        if char_to_parse == MISS_SYMBOL:
            return 0
        elif char_to_parse == SPARE_SYMBOL:
            return 10 - self.get_knocked_pins_of(string_roll, idx - 1)
        elif char_to_parse == STRIKE_SYMBOL:
            return 10
        
        return int(char_to_parse)

        