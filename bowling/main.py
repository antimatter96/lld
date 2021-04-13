from Game.player import Player
from Game.game import Bowling_Game
from Game.score_calc import Score_Calculator
from constants import STRIKE_SYMBOL, SPARE_SYMBOL , MISS_SYMBOL

def declare_winner(winner):
    print(f'{winner.get_name()} has won the game with {winner.get_final_score()}')

def display_scorecard(players):
    print()
    print('-- ScoreCard ---')
    print()
    for idx, player in enumerate(players):
        print( f" {idx + 1} - {player.get_name()}  scored  {player.get_final_score()}")

if __name__ == '__main__':
    score_calc = Score_Calculator()
    bowling_game = Bowling_Game(score_calc)
    
    print(f"Welcome to the Bowling Alley !!")

    try:
        no_of_players = int(input('Enter the number of Players: '))
    except:
        print('No of players must be an integer. Please rerun the CLI')
        quit()
    players = []

    for idx in range(no_of_players):
        player_name = input('Enter player name: ')
        players.append(Player(player_name))
    
    print()
    print(f"** Use {STRIKE_SYMBOL} for STRIKE, {SPARE_SYMBOL} for SPARE , {MISS_SYMBOL} for a MISS **")
    print()
    for player in players:
        player_frame = input(f"Enter {player.get_name()} bowling frame: ")
        if not player_frame:
            print('Blank Player Frame is not Allowed')
            quit()
        player_score = bowling_game.score_of_game(player_frame)
        player.set_final_score(player_score)
    
    players.sort(key=lambda player: player.get_final_score(), reverse=True)
    
    declare_winner(players[0])
    display_scorecard(players)

