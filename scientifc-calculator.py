import math

class ScientificCalculator:
    def __init__(self):
        pass

    def add(self, a, b):
        return a + b

    def subtract(self, a, b):
        return a - b

    def multiply(self, a, b):
        return a * b

    def divide(self, a, b):
        return a / b

    def power(self, a, b):
        return a ** b

    def square(self, a):
        return a * a

    def cube(self, a):
        return a * a * a

    def root(self, a, b):
        return a ** (1 / b)

    def log(self, a, b):
        return math.log(a, b)

    def sin(self, a):
        return math.sin(a)

    def cos(self, a):
        return math.cos(a)

    def tan(self, a):
        return math.tan(a)

    def asin(self, a):
        return math.asin(a)

    def acos(self, a):
        return math.acos(a)

    def atan(self, a):
        return math.atan(a)

if __name__ == "__main__":
    calc = ScientificCalculator()

    result = calc.add(1, 2)
    print(result)

    result = calc.subtract(5, 3)
    print(result)

    result = calc.multiply(4, 5)
    print(result)

    result = calc.divide(10, 2)
    print(result)

    result = calc.power(2, 3)
    print(result)

    result = calc.square(5)
    print(result)

    result = calc.cube(3)
    print(result)

    result = calc.root(8, 3)
    print(result)

    result = calc.log(10, 100)
    print(result)

    result = calc.sin(math.pi / 2)
    print(result)

    result = calc.cos(math.pi)
    print(result)

    result = calc.tan(math.pi / 4)
    print(result)

    result = calc.asin(1)
    print(result)

    result = calc.acos(0)
    print(result)

    result = calc.atan(1)
    print(result)
