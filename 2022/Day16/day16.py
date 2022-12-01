import collections

class Valve:
    def __init__(self, name, flow_rate, index) -> None:
        self.name = name
        self.flow_rate = flow_rate
        self.index = index
        self.neighbours = []

    def __repr__(self):
        return self.name

    def addNeighbour(self, neighbour):
        self.neighbours.append(neighbour)

    def searchSpaceIndex(self, index):
        self.SSIndex = index
        self.bitmap = 1 << index

def returnValve(name, flow_rate=0):
    global valves

    if any(valve.name == name for valve in valves.values()):
        return valves[name]
    else:
        valves[name] = Valve(name, flow_rate, len(valves))
        return valves[name]

def run(minutes):
    queue = collections.deque()
    best = collections.defaultdict(lambda: -1)

    def addToQueue(node, bitmap, flow, time):
        if time >= 0 and (best[(node, bitmap, time)] < flow):
            best[(node, bitmap, time)] = flow
            queue.append((node, bitmap, flow, time))

    addToQueue(valves["AA"], 0, 0, minutes)

    while queue:
        node, bitmap, flow, time = queue.popleft()

        # Added is a map of all nodes visited that we are searching, so after visiting the first of 5 it would look like 00001
        # We identify if we are at that node and we have not visited it before we add the total flow here to the queue and mark the node as visited

        # To mark the node as visited we put 1 in the location for the index (index 1 > 1, index three > 00100) and | that with the visited nodes
        # To identify that the node has not been visited before, we check that there is no collision between the binary index of the node and the added map

        if (bitmap & node.bitmap) == 0 and time >= 1:                                        
            flow_here = (time - 1) * node.flow_rate
            addToQueue(node, bitmap | node.bitmap, flow + flow_here, time - 1)

        for node_dest in search_space:
            time_move = distance[node.index][node_dest.index]
            if time_move <= time:
                addToQueue(node_dest, bitmap, flow, time - time_move)

    return best

with open("/Users/alex.waring/code/AoC2022/Day16/input.txt") as file:

    valves = {}
    max_nodes = 60
    search_space = []
    distance = [[max_nodes + 10] * max_nodes for _ in range(max_nodes)]
    for i in range(max_nodes):
        distance[i][i] = 0
    
    for line in file:
        values = line.strip().split()
        valve = values[1]
        flow_rate = int(values[4][:-1].split("=")[1])
        neighbours = values[9:]

        valve_obj = returnValve(valve, flow_rate)

        if flow_rate > 0 or valve == "AA":
            valve_obj.searchSpaceIndex(len(search_space))
            search_space.append(returnValve(valve))
        for neighbour in neighbours:
            valve_obj.addNeighbour(neighbour.strip(","))

    for valve in valves.values():
        for neighbor in valve.neighbours:
            distance[valve.index][valves[neighbor].index] = 1

    for valve in valves.values():
        for start in range(len(valves)):
            for end in range(len(valves)):
                distance[start][end] = min(distance[start][end], distance[start][valve.index] + distance[valve.index][end])

    best = run(30)
    print(max(best.values()))
    
    best_part2 = run(26)

    possible_combinations = 1 << len(search_space)

    # Create a table the length of all possible combinations of visited nodes, then loop through the results from 26 minutes
    # For each bitmap, not result, find the maximum flow for that entry

    table = [0] * possible_combinations
    for (node, bitmap, time), max_flow in best_part2.items():
        if max_flow > table[bitmap]:
            table[bitmap] = max_flow

    part2 = 0

    # for each possible bitmap, create a map (map3) that is the opposite of it
    # Take the max flow for those two entries
    # Loop down, removing the leading 1 each time (111011, 11011, 1011, 11, 1)

    for bitmap in range(possible_combinations):
        map3 = ((possible_combinations - 1) ^ bitmap)
        part2 = max(part2, table[map3])
        map2 = bitmap
        while map2 > 0:
            part2 = max(part2, table[map3] + table[map2])
            map2 = (map2 - 1) & bitmap


    print(part2)