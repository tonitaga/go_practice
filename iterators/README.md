# Go Iterators

## Основные понятия и концепции

- Ошибка на единицу (англ. **off-by-one error**) — логическая ошибка в алгоритме (или в математических вычислениях), когда количество итераций пошагового цикла оказывается на единицу меньше или больше необходимого.

До `Go 1.23` мы могли использовать `for range` для итерации по мапам, слайсам, массивам, строкам, каналам и целочисленным последовательностям.

- [`ForRange`](./for_range/main.go)

`for range` решение **off-by-one error**.

> **Как итерироваться по более сложной последовательности данных?** (Например, представленной в виде дерева, односвязного списка ...)

**Паттерн «Итератор» (Iterator)** — поведенческий паттерн проектирования, который предоставляет последовательный доступ ко всем элементам составного объекта, не раскрывая его внутреннего представления.

- [`SomeCollectionIterator`](./iterator_pattern/collection/main.go)

- [`LinkedListIterator`](./iterator_pattern/linked_list/main.go)

**В Go 1.23 появились итераторы**

Итератор в Go 1.23 — это функция, которая последовательно проходит через элементы последовательности и отправляет их в функцию обратного вызова, обычно называемую yield. Функция останавливается, когда достигает конца последовательности или когда yield сигнализирует о раннем прекращении, возвращая false.

- [`IterRange #1`](./std_iterators/iter_range/main.go)

Параметр функции называется `yield` по соглашению - само имя не имеет значения. `yield` может иметь `0,1,2` параметра и должна возвращать `bool`

- [`IterRange #2`](./std_iterators/iter_range_variations/main.go)

**Результат функции `yield` обязателен к проверке, чтобы отлавливать ранние выходы из `range`**

Если не проверять ранний выход, то происходит паника
- [`IterPanic`](./std_iterators/no_yield_handle/main.go)

```shell
panic: runtime error: range function continued iteration after function for loop body returned false
```

Динамическое формирование итератора. Использование пакета `iter`

- [`DynamicIterRange`](./std_iterators/dynamic_iter_range/main.go)
- [`IterSequence`](./std_iterators/iter_seq/main.go)

### Composing iterators

- [`EvenRange`](./std_iterators/composing_iterators/main.go)

### Backward

- [`BackwardRange`](./std_iterators/backward/main.go)

### Slices Package

- All
- Values
- Backward
- Collect
- AppendSeq
- Sorted
- SortedFunc
- SortedStableFunc
- Chunk

### Maps package

- All
- Keys
- Values
- Insert
- Collect

## Примеры использования

Итерирование на связанном списке при помощи Go-итераторов.

- [`LinkedListIterator`](./std_iterators/linked_list/main.go)

Итераторы при работе с большим набором данных это часто про ленивое чтение:

- Читая из базы или файла данные, где несколько миллионов строк, мы не хотим выгружать все эти строки в память. Мы хотим читать их частями и обрабатывать последовательно

## Особенности и ньюансы

### Pull/Push iterators

Push-итератор в языке программирования Go — это итератор, который передаёт каждое значение последовательности в функцию yield.

- [`Fibonacci`](./std_iterators/fibonacci/main.go)

Pull-итератор в языке программирования Go — это функция, которая позволяет извлекать одно значение за раз из последовательности

- [`Pull`](./std_iterators/pull_iterator/main.go)


> **Зачем?**

Например, когда нужно одновременно идти по двум итераторам

- [`TwoIterWalk`](./std_iterators/pull_two_iterators/main.go)

### Бенчмарки с итераторами

- [`Benchmark #1`](./benchmark/bench_test.go)
- [`Benchmark #2`](./benchmark/iter_slice/bench_test.go)

### Рекурсивные итераторы

- [`RecursiveIterator`](./std_iterators/recursive_iterator/main.go)

### Конкурентные итераторы

Неправильный подход использования итераторов и горутин

- [`WrongIterGoroutine`](./std_iterators/concurrent/goroutine_first/main.go)

Так как yield вызывается не последовательно, а параллельно.

> **Тогда как?**

Можно развернуть `push` итераторов в `pull` и уже обрабатывать значения в отдельных горутинах

- [`PullIterGoroutine`](./std_iterators/concurrent/pull_iter_goroutine/main.go)

### Паники с итераторами

- [`PanicFirst`](./std_iterators/panic/first/main.go)
- [`PanicSecond`](./std_iterators/panic/second/main.go)

### Naming conventions

В языке программирования Go (Golang) функции и методы итераторов (итераторов) называются в зависимости от последовательности, по которой они проходят.
Это соглашение закреплено в документации пакета iter. 

**Правила**

- Для метода итератора на типе коллекции традиционно используется имя `All`, потому что он итерирует последовательность всех значений в коллекции.

```go
func (s *Storage[T]) All() iter.Seq[T] { /* ... */ }
```

- Для типа, содержащего несколько возможных последовательностей, имя итератора может указывать, какая последовательность предоставляется.

```go
func (c *Mall) PetShops() iter.Seq[*PetShop] { /* ... */ } 
func (c *Mall) FoodShops() iter.Seq[*FoodShop] { /* ... */ }
```

- Если итератор требует дополнительной конфигурации, конструкторная функция может принимать дополнительные аргументы конфигурации.

```go
func (s *Storage[T]) Select(operation func(T) bool) iter.Seq[T] { /* ... */ }
```

- Когда существует несколько возможных порядков итерации, имя метода может указывать на этот порядок.

```go
func (s *Storage[T]) Backward() iter.Seq[T] { /* ... */ }

func (s *Tree[T]) InOrder() iter.Seq[T] { /* ... */ }
func (s *Tree[T]) PreOrder() iter.Seq[T] { /* ... */ }
```

## Внутреннее устройство

Итераторы - это синтаксический сахар

- [`SyntaxSugar`](./std_iterators/syntax_sugar/main.go)

### PUSH & PULL итераторы - кардинально разный поток управления

- PUSH итераторы управляют процессом итерации, передавая значения в функцию до тех пох, пока не закончатся данные или явно не попросят завершить итерирование

- PULL управляют извне и должны сохранять свое состояние между вызовами (Реализованы на основе корутин)
