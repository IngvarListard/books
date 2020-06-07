"""
Blah blah blah
"""
from math import hypot
import collections

Card = collections.namedtuple('Card', ['rank', 'suit'])

class FrenchDeck:
    ranks = [str(n) for n in range(2, 11) + list('JQKA')]
    suits = 'spades diamonds clubs hearts'.split()

    def __init__(self):
        self._cards = [Card(rank, suit) for suit in self.suits
                                        for rank in self.ranks]

        def __len__(self):
            return len(self._cards)

        def __getitem__(self, position):
            """Метод вызываемый при попытке получения объекта по
            индексу

            """
            return self._cards[position]


class Vector:

    def __init__(self, x=0, y=0):
        self.x = x
        self.y = y

    def __repr__(self):
        """Строковое представление данных. Желательно возвращать данные
        в виде в котором можно воспроизвести сам класс"""
        return 'Vector(%r, %r)' % (self.x, self.y)

    def __abs__(self):
        """Возврат абсолютного значения функцие abs()"""
        return hypot(self.x, self.y)

    def __bool__(self):
        """Возврат булевого занчения"""
        return bool(abs(self))

    def __add__(self, other):
        """Перегрузка оператора сложения"""
        x = self.x + other.x
        y = self.y + other.y
        return Vector(x, y)

    def __mul__(self, scalar):
        """Перегрузка оператора умножения (стр 36)"""
        return Vector(self.x * scalar, self.y * scalar)


"""Распаковка с помощью *"""
a, b, *rest = range(5)
a, b, *rest = range(3)
a, b, *rest = range(2)
a, *body, c, d = range(5)
*head, b, c, d = range(5)

"""Именованый кортеж"""
from collections import namedtuple

City = namedtuple('City', 'name country population coordinates')
City._fields  # возвращает поля
City._make(*rest)  # вызывает конструктор класса
City._asdict()  # возвращает ordered_dict

"""Срезы"""
s = 'bicycle'
s[::3]  # срез с шагом
s[::-1]
s[::-2]

"""Изменение массивов 'на лету'"""
arr = range(1, 10)
arr[2:5] = [20, 30]
del arr[5:7]
arr[3::2] = [11, 22]


"""Словарь по типу defaultdict без defaultdict"""
d = dict{}
d.setdefault(key, []).append('data')

"""
__missing__ - вызывается если ключ[n] отсутствует в объекте
collections.ChainMap является списком словарей. Он позволяет искать
элемент сразу во всех словарях в этом списке

"""


"""
Строки, бинарные строки и массивы
"""
# можно получить bytes из строки, если известна кодировка
cafe = bytes('cafe', encoding='utf_8')
# срез bytes вернёт объект bytes
cafe[:1]
""">>> b'c' """
# обращение по массиву вернёт целое число от 0 до 255
cafe[0]
""">>> 99 """

# bytes можно преобразовать в bytesarray
cafe_arr = bytesarray(cafe)
# срез bytesarray также вернёт объект bytesarray
cafe_arr[-1:]
""">>> bytearray(b'c') """

"""Использование memoryview и struct для извлечения полей из заголовка
GIF-изображения"""
import struct
fmt = '<3s3sHH'
with open('filter.gif', 'rb') as fp:
    img = memoryview(fp.read())

header = img[:10]
bytes(header)
# >>> b'GIF89a+\x02\xe6\x00'
struct.unpack(fmt, header)
# >>> (b'GIF', b'89a', 555, 230)
#

import collections
Event = collections.namedtuple('Event', 'time proc action')
def taxi_proceess(ident, trips, start_time=0):
    time = yield Event(start_time, ident, 'leave garage')
    for i in rageng(trips):
        time = yield Event(time, ident, 'pick up passenger')
        time = yield Event(time, ident, 'drop off passenger')

    yield event(time, ident, 'going home')



