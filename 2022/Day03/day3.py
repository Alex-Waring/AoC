from itertools import islice

def part1(items: str) -> int:
    string = items[:-1]
    n = len(string)//2
    output = 0

    for letter in set(string[0:n]):
        if letter in string[n:]:
            output = getPriorities(letter)

    return(
        output
    )

def getPriorities(letter: str) -> int:
    if letter.isupper():
        return(ord(letter) - 38)
    else:
        return(ord(letter) - 96)

def part2(line1: str, line2, line3):
    for letter in set(line1[:-1]):
        if (letter in line2) and (letter in line3):
            return(getPriorities(letter))

if __name__ == "__main__":
    with open("input.txt") as file:
        priorities = []
        for line in file.readlines():
            priorities.append(part1(line))
        print(sum(priorities))

    with open("input.txt") as file:
        priorities = []

        lines = []
        for line in file:
            lines.append(line)
            if len(lines) >= 3:
                priorities.append(part2(lines[0], lines[1], lines[2]))
                lines = []
        print(sum(priorities))