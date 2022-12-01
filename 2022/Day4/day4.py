def parseInput(sections):
    elf_1 = sections.split(",")[0]
    elf_2 = sections.split(",")[1]

    output = [[], []]

    output[0] = elf_1.split("-")
    output[1] = elf_2.split("-")

    return output

def part1(input):
    output = 0
    for list in input:
        elf_1 = set(range(int(list[0][0]), int(list[0][1])+1))
        elf_2 = set(range(int(list[1][0]), int(list[1][1])+1))
        if elf_1.issubset(elf_2) or elf_2.issubset(elf_1):
            output += 1
    return output

def part2(input):
    output = 0
    for list in input:
        elf_1 = set(range(int(list[0][0]), int(list[0][1])+1))
        elf_2 = set(range(int(list[1][0]), int(list[1][1])+1))
        if len(elf_1 & elf_2) > 0:
            output += 1
    return output

if __name__ == "__main__":
    with open("input.txt") as file:
        sections = []
        for line in file:
            sections.append(parseInput(line[:-1]))
        print(part1(sections))
        print(part2(sections))