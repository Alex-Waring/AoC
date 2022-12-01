import numpy

headPosition = {
    "x" : 1,
    "y" : 1
}

bodyPositions = [
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
    { "x" : 1,"y" : 1},
]

totalPositions = {
    "1.1" : "s"
}

def moveHead(direction, steps):

    for step in range(1, steps + 1):
        match direction:
            case "R":
                bodyPositions[0]["x"] += 1
            case "L":
                bodyPositions[0]["x"] -= 1
            case "U":
                bodyPositions[0]["y"] += 1
            case "D":
                bodyPositions[0]["y"] -= 1

        for part in range(len(bodyPositions) - 1):
                    if calculateDistance(bodyPositions[part], bodyPositions[part + 1]):
                        moveBody(part, part + 1)

        location = str(bodyPositions[-1]["x"]) + "." + str(bodyPositions[-1]["y"])
        totalPositions[location] = "#"

def moveBody(first, second):

    x_dif = bodyPositions[first]["x"] - bodyPositions[second]["x"]
    y_dif = bodyPositions[first]["y"] - bodyPositions[second]["y"]

    if y_dif == 0:
        bodyPositions[second]["x"] += 1 * numpy.sign(x_dif)
    elif x_dif == 0:
        bodyPositions[second]["y"] += 1 * numpy.sign(y_dif)
    else:
        bodyPositions[second]["x"] += 1 * numpy.sign(x_dif)
        bodyPositions[second]["y"] += 1 * numpy.sign(y_dif)

def calculateDistance(first, second):

    x_dif = abs(first["x"] - second["x"])
    y_dif = abs(first["y"] - second["y"])

    if x_dif > 1 or y_dif > 1:
        return True
    else:
        return False

with open("input.txt") as file:
    for line in file:
        line = line[:-1]
        command = line.split(" ")
        moveHead(command[0], int(command[1]))
    print(len(totalPositions))