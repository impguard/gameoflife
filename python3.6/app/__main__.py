from time import sleep

from app.game import state_to_str, random_state, transition


state = random_state(50, 50)

while(True):
    state = transition(state)
    print('\033[H')
    print('\033[2J')
    print(state_to_str(state))
    sleep(0.05)

    if not any(state):
        print('The game is over!')
        break
