import numpy

def addToCave(input):
    for key in range(len(input) - 1):

        start_x = int(input[key].split(",")[0])
        start_y = int(input[key].split(",")[1])
        y_values.append(start_y)
        end_x = int(input[key + 1].split(",")[0])
        end_y = int(input[key + 1].split(",")[1])

        if start_x == end_x:
            if start_y < end_y:
                for y in range(start_y, end_y + 1):
                    cave[y][start_x] = "#"
            else:
                for y in range(end_y, start_y + 1):
                    cave[y][start_x] = "#"
        if start_y == end_y:
            if start_x < end_x:
                for x in range(start_x, end_x + 1):
                    cave[start_y][x] = "#"
            else:
                for x in range(end_x, start_x + 1):
                    cave[start_y][x] = "#"
    y_values.append(int(input[-1].split(",")[1]))
    return

def sandFall(max_y):

    y = 0
    x = 500

    while True:
        cave[y][x] = "0"
        if cave[y+1][x] == ".":
            cave[y][x] = "."
            cave[y+1][x] = "0"
        elif cave[y+1][x-1] == ".":
            cave[y][x] = "."
            cave[y+1][x-1] = "0"
            x -= 1
        elif cave[y+1][x+1] == ".":
            cave[y][x] = "."
            cave[y+1][x+1] = "0"
            x +=1
        else:
            return False
        y += 1

def checkBlocked(cave):
    if cave[0][500] == "0":
        return True
    else:
        return False


with open("input.txt") as file:
    cave = list()
    y_values = []

    for i in range(200):
        cave.append([])
        for x in range(800):
            cave[i].append(".")

    for line in file:
        input = line.strip().split(" -> ")
        addToCave(input)
file.close()

for y in range(len(cave)):
    for x in range(len(cave[y])):
        if y == max(y_values) + 2:
            cave[y][x] = "#"

sand = 0

while not checkBlocked(cave):
    sandFall(max(y_values))
    sand += 1

with open("output.txt", "w") as file:
    for row in cave:
        output = ""
        for index in row:
            output += index
        file.write(output + "\n")

print(sand)