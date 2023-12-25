def part1(input: str) -> int:
    against = input.split(" ")[0]
    you = input.split(" ")[1][:-1]

    score = 0

    if (against == "A" and you == "Y") or (against == "B" and you == "Z") or (against == "C" and you == "X"):
        score += 6
    elif (against == "A" and you == "X") or (against == "B" and you == "Y") or (against == "C" and you == "Z"):
        score += 3

    if you == "X":
        score += 1
    elif you == "Y":
        score += 2
    else:
        score += 3

    return score

def part2(input: str) -> int:
    against = input.split(" ")[0]
    you = input.split(" ")[1][:-1]

    score = 0

    if you == "Y":
        score += 3
        if against == "A":
            score += 1
        elif against == "B":
            score += 2
        else:
            score += 3
    elif you == "Z":
        score += 6
        if against == "A":
            score += 2
        elif against == "B":
            score += 3
        else:
            score += 1
    else:
        if against == "A":
            score += 3
        elif against == "B":
            score += 1
        else:
            score += 2

    return(score)

if __name__ == "__main__":
    with open("input.txt") as file:
        scores = []
        for line in file.readlines():
            scores.append(part2(line))
        print(sum(scores))