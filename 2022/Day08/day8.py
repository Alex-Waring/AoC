def buildMap(file):
    forest = []

    for line in file:
        line = line[:-1]
        column = 1
        for height in line:
            if len(forest) < column:
                forest.append([height])
            else:
                forest[column - 1].append(height)
            column += 1

    return forest

def tallerTrees(forest, x, y):
    height = forest[y][x]
    visable = {
        "above" : 0,
        "below" : 0,
        "left" : 0,
        "right" : 0,
    }

    for above in reversed(range(y)):
        if forest[above][x] >= height:
            visable["above"] += 1
            break
        else:
            visable["above"] += 1

    for below in range(y+1, len(forest)):
        if forest[below][x] >= height:
            visable["below"] += 1
            break
        else:
            visable["below"] += 1

    for left in reversed(range(x)):
        if forest[y][left] >= height:
            visable["left"] += 1
            break
        else:
            visable["left"] += 1

    for right in range(x+1, len(forest)):
        if forest[y][right] >= height:
            visable["right"] += 1
            break
        else:
            visable["right"] += 1

    return (visable["above"] * visable["below"] * visable["left"] * visable["right"])



with open("input.txt") as file:
    forest = buildMap(file)

    count = []

    for y in range(len(forest)):
        for x in range(len(forest[y])):
            count.append(tallerTrees(forest, x, y))

    print(count)
    print(max(count))