# a + b + c = 1000 a*a + b*b = c*c

# first
# import time
#
# start_time = time.time()
#
# for a in range(0, 1001):
#     for b in range(0, 1001):
#         for c in range(0, 1001):
#             if a + b + c == 1000 and a ** a + b ** b == c ** c:
#                 print("a, b, c: %d, %d, %d" % (a, b, c))
#
# end_time = time.time()
# print("time: %d", end_time - start_time)
# print("finished")
#

# second
import time

start_time = time.time()
for a in range(0, 1000):
    for b in range(0, 1000):
        c = 1000 - a -b
        if a**2 + b**2 == c**2:
            print("a, b, c = %d, %d, %d" % (a, b, c))

end_time = time.time()
print('Cost: %f' % (end_time - start_time))
print('Finished!')