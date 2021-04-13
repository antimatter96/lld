from Frame.frame_parser import Frame_Parser
from Roll.roll_parser import Roll_Parser


class Bowling_Game:
    def __init__(self, score_calc):
        self._roll_parser = Roll_Parser()
        self._frame_parser = Frame_Parser()
        self.score_calculator = score_calc
    
    def score_of_game(self, input):
        try:
            parsed_roll = self._roll_parser.parse(input)
        except Exception as e:
            print(f"{e}. Please rerun the CLI")
            quit()
        parsed_frame = self._frame_parser.parse(parsed_roll)
        score = self.score_calculator.calc_score(parsed_frame)
        return score

