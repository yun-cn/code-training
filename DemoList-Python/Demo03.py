import string
value = {'var': 'foo'}

t = string.Template("$var is here but $missing is not provided")

try:
    print('substitute()   :', t.substitute(value))
except KeyError as err:
    print('Error: ', str(err))

print('safe_substitute():', t.safe_substitute(value))