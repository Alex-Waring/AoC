import networkx

heightmap = []

def createGraph(heightmap):
    graph = networkx.DiGraph()
    for y in range(len(heightmap)):
        for x in range(len(heightmap[y])):

            graph.add_node((x, y))
            adjacent = [(x - 1, y), (x + 1, y), (x, y + 1), (x, y - 1)]

            for adj_x, adj_y in adjacent:
                if 0 <= adj_x < len(heightmap[y]) and 0 <= adj_y < len(heightmap):
                    if heightmap[adj_y][adj_x] - heightmap[y][x] <= 1:
                        graph.add_edge((x, y), (adj_x, adj_y))

    return graph


def SPF(start, end):
    graph = createGraph(heightmap)
    path = networkx.shortest_path(graph, source=start, target=end)
    return len(path) - 1


with open("input.txt") as file:
    for line in file:
        line = line.strip()
        row = []
        for letter in line:
            row.append(ord(letter) - 96)
        heightmap.append(row)

for y, row in enumerate(heightmap):
    for x, value in enumerate(row):
        if value == (ord("S") - 96):
            heightmap[y][x] = 0
            start = (x, y)
        elif value == (ord("E") - 96):
            heightmap[y][x] = 26
            end = (x, y)

path = SPF(start, end)
print(path)
paths = [path]

for y, row in enumerate(heightmap):
    for x, value in enumerate(row):
        if value == (ord("a") - 96):
            start = (x, y)
            try:
                paths.append(SPF(start, end))
            except:
                print("Nope!! You're lost")

print(min(paths))