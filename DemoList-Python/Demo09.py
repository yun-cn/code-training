from timeit import Timer


def test1():
    li = []
    for i in range(10000):
        li.append(i)


def test2():
    li = []
    for i in range(10000):
        li += [i]


def test3():
    li = [i for i in range(10000)]


def test4():
    li = list(range(10000))


timer1 = Timer("test()", "from __main__ import test1")
print("append", timer1.timeit(1000))

timer2 = Timer("test()", "from __main__ import test2")
print("+=", timer2.timeit(1000))

timer3 = Timer("test()", "from __main__ import test3")
print("[i for i in range]", timer3.timeit(1000))

timer4 = Timer("test()", "from __main__ import test4")
print("list(range())", timer4.timeit(1000))
