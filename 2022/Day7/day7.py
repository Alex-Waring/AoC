from collections import defaultdict
from itertools import accumulate

directories = defaultdict(int)

with open("input.txt") as file:

    for line in file:
        match line.split():
            case '$', 'cd', '/': currentDirectory = ['']
            case '$', 'cd', '..': currentDirectory.pop()
            case '$', 'cd', x: currentDirectory.append(x+'/')
            case '$', 'ls': pass
            case 'dir', _: pass
            case size, _:
                for p in accumulate(currentDirectory):
                    directories[p] += int(size)

    print(sum(size for size in directories.values() if size <= 100000))
    print(min(size for size in directories.values() if size >= directories[''] - 40000000))