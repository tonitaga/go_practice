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
### Бенчмарки с итераторами

- [`Benchmark`](./benchmark/bench_test.go)

```shell
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics         
BenchmarkIteratorRange-16       313387570                3.849 ns/op
BenchmarkUsualRange-16          310089063                3.869 ns/op
PASS
ok      command-line-arguments  2.412s
```

### Рекурсивные итераторы
### Конкурентные итераторы
### Паники с итераторами
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