cycle = 1
values = []
screen = []
x_register = 1

def noop():
    clock_cycle()

def clock_cycle():
    values.append(x_register)
    global cycle

    position_x = (cycle % 40) -1
    value_x = x_register % 40

    if position_x == value_x or position_x == value_x + 1 or position_x == value_x - 1:
        screen.append("#")
    else:
        screen.append(".")

    cycle += 1

def addx(counter):
    clock_cycle()
    clock_cycle()
    global x_register 
    x_register += counter

def spriteVisable(row, column):
    local_cycle = row * 40 + (column - 1)
    if values[local_cycle] == column or values[local_cycle] + 1 == column or values[local_cycle] - 1 == column:
        return True
    else:
        return False


def signalstrength(cycle):
    return cycle * values[cycle - 1]

with open("input.txt") as file:
    for line in file:
        line = line[:-1]
        command = line.split(" ")
        if len(command) == 1:
            noop()
        else:
            addx(int(command[1]))

    output = ""
    lines = [""]

    for x in range(1, 7):
        start = (x-1) * 40
        end = x * 40
        lines.append("")
        for y in range(start, end):
            lines[x] += screen[y]

    for line in lines:
        print(line)