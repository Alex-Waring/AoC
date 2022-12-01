import itertools
import functools

def compare(data):
    left = data[0]
    right = data[1]

    if type(left) is int and type(right) is int:
        if left < right:
            return True
        if left > right:
            return False
        else:
            return
        
    if type(left) is int:
        left = [left]
    if type(right) is int:
        right = [right]

    if len(left) == 0 and len(right) != 0:
        return True       
    elif len(left) != 0 and len(right) == 0:
        return False 
    elif len(left) == 0 and len(right) == 0:
        return 

    recurse = compare([left[0], right[0]])
    if recurse != None:
        return recurse
    else:
        return compare([left[1:],right[1:]])

def compareTransform(left, right):
    if compare([left, right]):
        return -1
    else:
        return +1
        
with open("input.txt") as file:
    
    data = []

    while True:
        data_raw = list(itertools.islice(file, 3))
        if not data_raw:
            break
        else:
            input = []
            input.append(eval(data_raw[0].strip()))
            input.append(eval(data_raw[1].strip()))
            data.append(input)

    correct_order = [
        i for i in range(1, len(data) + 1)
        if compare(data[i-1])
    ]
    print(sum(correct_order)) # Part 1

    data = []

with open("/Users/alex.waring/code/AoC2022/Day13/input.txt") as file:
    for line in file:
        if line.strip() != "":
            data.append(eval(line))
    data.append([[2]])
    data.append([[6]])

    data.sort(key=functools.cmp_to_key(compareTransform))

    print((data.index([[2]]) + 1) * (data.index([[6]]) + 1))