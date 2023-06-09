# Подземная доставка

| Ресурс | Ограничение |
| :--- | :--- |
| Время | 2 секунды |
| Память | 256Mb |
| Ввод | стандартный ввод или input.txt |
| Вывод | стандартный вывод или output.txt |

Для ускорения работы служб доставки под городом Длинноградом был прорыт тоннель, по которому ходит товарный поезд, останавливающийся на промежуточных станциях возле логистических центров. На станциях к концу поезда могут быть присоединены вагоны с определенными товарами, а также от его конца может быть отцеплено некоторое количество вагонов или может быть проведена ревизия, во время которой подсчитывается количество вагонов с определенным товаром.

Обработайте операции в том порядке, в котором они производились, и ответьте на запросы ревизии.
## Формат ввода
В первой строке вводится число **_N_** **(1≤N≤100000)** — количество операций, произведенных над поездом.

В каждой из следующих **_N_** строк содержится описание операций. Каждая операция может иметь один из трех типов:

**_add_** <количество вагонов> <название товара> — добавить в конец поезда <количество вагонов> с грузом <название товара>. Количество вагонов не может превышать $10^9$, название товара — одна строка из строчных латинских символов длиной до 20.

**_delete_** <количество вагонов> — отцепить от конца поезда <количество вагонов>. Количество отцепляемых вагонов не превосходит длины поезда.

**_get_** <название товара> — определить количество вагонов с товаром <название товара> в поезде. Название товара — одна строка из строчных латинских символов длиной до 20.

## Формат вывода
На каждый запрос о количестве вагонов с определенным товаром выведите одно число — количество вагонов с таким товаром. Запросы надо обрабатывать в том порядке, как они поступали.
### Пример 1
```
Ввод						Вывод

7						20
add 10 oil					15
add 20 coal					21
add 5 oil
get coal
get oil
add 1 coal
get coal
```

### Пример 2
```
Ввод						Вывод

6						0
add 5 oil					2
get coal					0
add 7 liverstock
delete 10
get oil
get liverstock
```
