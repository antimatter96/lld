from Frame.frame_factory import Frame_Factory

class Frame_Parser:
    def __init__(self):
        self._frame_factory = Frame_Factory()

    def parse(self, rolls):
        no_of_roll = 0
        frames = []

        while no_of_roll < len(rolls):

            curr_frame = self._frame_factory.get_frame(rolls, no_of_roll)
            if curr_frame.is_strike():
                no_of_roll += 1
            else:
                no_of_roll += 2
            frames.append(curr_frame)
        
        return frames