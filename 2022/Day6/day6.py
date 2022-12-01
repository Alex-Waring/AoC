def identifyPairs(buffer):
    with open("input.txt") as file:
        chars = file.read()

        check = False
        index = 0

        while not check:
            if len(set(chars[index:index + buffer])) == buffer:
                print(index + buffer)
                check = True
            else:
                index += 1


if __name__ == "__main__":
    identifyPairs(4)
    identifyPairs(14)