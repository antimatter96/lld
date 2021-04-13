from constants import MAX_FRAMES

class Score_Calculator:
    def __init__(self):
        self._max_frames = MAX_FRAMES
    
    def calc_score(self, parsed_frames):
        score = 0
        if parsed_frames:
            for idx in range(self._max_frames):
                if parsed_frames[idx].is_strike():
                    score += parsed_frames[idx].no_of_knocked_pins() + self.bonus_strike(parsed_frames, idx)
                elif parsed_frames[idx].is_spare():
                    score += parsed_frames[idx].no_of_knocked_pins() + self.bonus_spare(parsed_frames, idx)
                else:
                    score += parsed_frames[idx].no_of_knocked_pins()
        return score


    def bonus_strike(self, parsed_frames, idx):
        bonus = 0
        if(parsed_frames[idx + 1].is_strike()):
            bonus += parsed_frames[idx + 1].no_of_knocked_pins() + parsed_frames[idx + 2].no_of_pins_knocked_in_first_roll()
        else:
            bonus += parsed_frames[idx + 1].no_of_knocked_pins()
        return bonus

    def bonus_spare(self, parsed_frame, idx):
        return parsed_frame[idx + 1].no_of_pins_knocked_in_first_roll()
        