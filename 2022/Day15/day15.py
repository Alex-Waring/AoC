import z3
from multiprocessing import Process
import multiprocessing
import time

class Sensor:
    def __init__(self, x, y, beacon_x, beacon_y) -> None:
        self.x = x
        self.y = y
        self.beacon_x = beacon_x
        self.beacon_y = beacon_y
        self.distance = abs(x - beacon_x) + abs(y - beacon_y)

    def distanceFrom(self, x, y):
        return abs(x - self.x) + abs(y - self.y)

    def __str__(self) -> str:
        return "{}, {}".format(self.x, self.y)

def part1_loop(start, end, sensors, results):

    def calcDistance(start_x, start_y, end_x, end_y):
        abs_distance = abs(end_x - start_x) + abs(end_y - start_y)
        return abs_distance

    count = 0
    y = 2000000
    for x in range(start, end):
        poss = True
        for sensor in sensors:
            if (x,y) == (sensor.beacon_x, sensor.beacon_y):
                poss = True
                break
            if calcDistance(sensor.x, sensor.y, x, y) <= sensor.distance:
                poss = False
                break
        if not poss:
            count += 1
    results.append(count)


with open("./input.txt") as file:
    cave = list()
    sensors = []

    cave = [ ["."] *  8000000] * 4000000

    for line in file:
        values = line.strip().split()
        sensor_x = int(values[2][:-1].split("=")[1])
        sensor_y = int(values[3][:-1].split("=")[1])
        beacon_x = int(values[8][:-1].split("=")[1])
        beacon_y = int(values[9].split("=")[1])

        sensors.append(Sensor(sensor_x, sensor_y, beacon_x, beacon_y))

if __name__ == '__main__':

    start = time.perf_counter()
    sections = [
        [-1000000, 0],
        [0, 1000000],
        [1000000, 2000000],
        [2000000, 3000000],
        [3000000, 4000000],
        [4000000, 5000000],
        [5000000, 6000000],
        [6000000, 7000000],
        [7000000, 8000000],
        [8000000, 9000000],
        [9000000, 10000000],
    ]
    procs = []
    manager = multiprocessing.Manager()
    results = manager.list()

    for section in sections:
        proc = Process(target=part1_loop, args=[section[0], section[1], sensors, results])
        procs.append(proc)
        proc.start()

    for proc in procs:
        proc.join()

    finish = time.perf_counter()
    print(f'Part 1 finished in {round(finish-start,2 )} second(s): {sum(results)}')
        
    start = time.perf_counter()
    solver = z3.Solver()
    x, y = z3.Int("x"), z3.Int("y")

    solver.add(0 <= x)
    solver.add(x <= 4000000)
    solver.add(0 <= y)
    solver.add(y <= 4000000)

    def z3_abs(x):
        return z3.If(x >=0, x, -x)

    for sensor in sensors:
        sensor_x = sensor.x
        sensor_y = sensor.y
        sensor_distance = sensor.distance
        solver.add(z3_abs(sensor_x - x) + z3_abs(sensor_y - y) > sensor_distance)

    solver.check()
    model = solver.model()
    finish = time.perf_counter()
    print(f'Part 2 finished in {round(finish-start,2 )} second(s): ', model[x].as_long() * 4000000 + model[y].as_long())