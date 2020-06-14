class Student:
    classroom = '101'
    address = 'beijing'

    def __init__(self, name, age):
        self.name = name
        self.age  = age

    def print_age(self):
        print('%s %s' % (self.name, self.age))