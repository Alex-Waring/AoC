crate_locations = [
    "",
    "QMGCL",
    "RRLCTFHG",
    "VJFNMTWR",
    "JFDVQQ",
    "NFMSLBT",
    "RNVHCDP",
    "HCT",
    "GSJVZNHP",
    "ZFHG"
]

def moveCrates(quantity, start, end):
    for x in range(quantity):
        crate_locations[end] = crate_locations[end] + crate_locations[start][-1]
        crate_locations[start] = crate_locations[start][:-1]

def moveCratePart2(quantity, start, end):
    volume = quantity * -1
    crate_locations[end] = crate_locations[end] + crate_locations[start][volume:]
    crate_locations[start] = crate_locations[start][:volume]

def parser(line):
    input = line.split("\n")[0]
    quantity = input.split(" ")[1]
    start = input.split(" ")[3]
    end = input.split(" ")[5]

    return quantity, start, end


if __name__ == "__main__":
    with open("input.txt") as file:
        for line in file:
            instruction = parser(line)
            moveCratePart2(int(instruction[0]), int(instruction[1]), int(instruction[2]))
    print(crate_locations)