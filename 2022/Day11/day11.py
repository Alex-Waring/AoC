import operator
import itertools
import math

ops = {
    "+" : operator.add,
    "-" : operator.sub,
    "*" : operator.mul,
    "**" : operator.pow,
}

class Monkey:

    def __init__(self, items, operation, test, true, false) -> None:
        self.items = items
        self.operation = operation
        self.test = test
        self.true = true
        self.false = false
        self.handled = 0

    def takeTurn(self):
        for key, item in enumerate(self.items):
            ops_function = ops[self.operation[0]]
            self.items[key] = ops_function(int(item), int(self.operation[1])) % lcm # // 3

            if self.items[key] % self.test == 0:
                monkeys[self.true].addItem(self.items[key])
            else:
                monkeys[self.false].addItem(self.items[key])

            self.handled += 1
        self.items = []

    def addItem(self, item):
        self.items.append(item)

    def __str__(self) -> str:
        return str(self.handled)

with open("input.txt") as file:

    monkeys = []

    while True:
        monkey_raw = list(itertools.islice(file, 7))
        if not monkey_raw:
            break
        else:
            list_items = [string.strip(",") for string in monkey_raw[1].split()][2:]
            operators = [monkey_raw[2].split()[-2], monkey_raw[2].split()[-1]]
            test = int(monkey_raw[3].split()[-1])
            true = int(monkey_raw[4].split()[-1])
            false = int(monkey_raw[5].split()[-1])

            monkeys.append(Monkey(
                list_items,
                operators,
                test,
                true,
                false
            ))

    lcm = 1
    for monkey in monkeys:
        lcm = math.lcm(lcm, monkey.test)
    print(lcm)

    for x in range(10000):
        if x % 1000 == 0:
            print(x)
        for monkey in monkeys:
            monkey.takeTurn()

    monkey_business = []

    for monkey in monkeys:
        monkey_business.append(monkey.handled)

    monkey_business.sort(reverse=true)
    print(monkey_business)
    print(monkey_business[0] * monkey_business[1])