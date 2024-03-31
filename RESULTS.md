# Сравнение двух реализаций lock-free стека

Все результаты тестирования представлены в [таблице](https://docs.google.com/spreadsheets/d/1vM2VX7_RzhkFGnl_R0tDx7UFg_vUejbcIK0tfLbvSbM/edit?usp=sharing).

Характеристики ПК:
- Процессор AMD Ryzen 5 5500U с общим количество ядер - 6 и с 12 логическими процессорами;
- Объем оперативной памяти: 8 Гб.

## Результаты

Исходя из тестовых данных, можно сделать следующий выводы

- _Горутинами надо пользоваться с умом_ (10000 горутин с 1 задачей медленнее чем 10 горутин с 1000
задачами, т.к. тратится меньше времени на шедулер и памяти на стеки горутин);
- Оптимальное количество горутин для работы стека -- 4, больше только замедляет (время на шедулер + память);
- При рандомной вставке и рандомном удалении все-таки стоит пользоваться стеком с оптимизацией
соответствующих операций
