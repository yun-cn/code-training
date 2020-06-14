import string

class MyTemplate(string.Template):
    delimiter = '%'
    idpattern = '[a-z]+_[a-z]+'

template_text = '''
    Delimiter    : %%
    Replaced     : %with_underscore
    Ignored      : %notunderscored
'''

d = {
    'with_underscore': 'replaced',
    'notunderscored': 'not replace'
}

t = MyTemplate(template_text)
print('Modified Id pattern: ')
print(t.safe_substitute(d))