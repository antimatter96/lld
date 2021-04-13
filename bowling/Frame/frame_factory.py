from Roll.roll import Roll
from Frame.frame import Frame, Spare, Strike


class Frame_Factory:
    def get_frame(self, rolls, idx):
        first_roll = rolls[idx]

        if first_roll.no_of_knocked_pins() == 10:
            return Strike(rolls[idx])
        
        if (idx + 1 < len(rolls)):
            second_roll = rolls[idx +1]
            if (first_roll.no_of_knocked_pins() + second_roll.no_of_knocked_pins()) == 10:
                return Spare(rolls[idx], rolls[idx + 1])
        else:
            second_roll = Roll()
        
        return Frame(first_roll, second_roll)