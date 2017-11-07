from time import sleep

from app.game import clear_state_str, state_to_str, random_state, transition


state = random_state(10, 10)

while(True):
    state = transition(state)
    print(clear_state_str(state))
    print(state_to_str(state))
    sleep(1)

    if not any(state):
        print('The game is over!')
        break
