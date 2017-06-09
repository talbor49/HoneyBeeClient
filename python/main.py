from honeybee import HoneyBee


hb = HoneyBee('localhost', 8080, 'talbor49', '1234')
print(hb.use('x'))
print(hb.set('tal', 'bae'))
print(hb.get('tal'))
