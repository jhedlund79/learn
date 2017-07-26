"""runs a sequence of programs
Each program except the first recieves the std output of the previous
program on its std input, by default.
"""

def _launch():
    print('pipeline launched')

if __name__ == '__main__':
    _launch()
