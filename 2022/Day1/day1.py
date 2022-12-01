import heapq

def getElves(input: str) -> list[int]:
    return[
        sum(map(int, lines.split()))
        for lines in input.split("\n\n")
    ]

def part1(elves: list[int]) -> int:
    return max(elves)

def part2(elves: list[int]) -> int:
    return sum(heapq.nlargest(3, elves))

if __name__ == "__main__":
    with open("input.txt") as file:
        print(part1(getElves(file.read())))
    with open("input.txt") as file:
        print(part2(getElves(file.read())))