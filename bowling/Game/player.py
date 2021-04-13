class Player:
    def __init__(self, name):
        self._name = name
        self._final_score = 0
    
    def set_final_score(self, score):
        self._final_score = score
    
    def get_final_score(self):
        return self._final_score
    
    def get_name(self):
        return self._name

    def __str__(self):
        return f"{self._name} scored {self._final_score}"