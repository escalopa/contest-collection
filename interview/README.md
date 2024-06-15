# Interview Questions

## Cyclic channel loop

```go
func main() {
  const size = 10000
  list := []chan int{make(chan int)}

  for i := 0; i < size; i++ {
    list = append(list, make(chan int))
    go func(i int) {
      go func() {
        for v := range list[i] {
          list[i-1] <- v + 1
        }
        go func() {
          close(list[i-1])
        }()
      }()
    }(i + 1)
  }

  go func() {
    list[len(list)-1] <- 1
    close(list[len(list)-1])
  }()

  fmt.Println(<-list[0]) // what will be printed & why?
  fmt.Println(<-list[0]) // what will be printed & why?
}
```

## Merge channels

```go
func main() {
  var (
    c1 = make(chan int)
    c2 = make(chan int)
    c3 = make(chan int)
  )

  // write value & close channel 
  go func(){ c1 <- 1; close(c1) }()
  go func(){ c2 <- 2; close(c2) }()
  go func(){ c3 <- 3; close(c3) }()

  out := mergeChannels(c1, c2, c3)

  // what will be printed & why?
  for value := range out {
    fmt.Printf("%d\n", value)
  }

  fmt.Println("Done")
}

func mergeChannels(chans ...<-chan int) <-chan int {
  // implement me
}
```

## Map ordering

```go
m := map[string]int{"a":1,"b":2,"c":3}

for a, b := range m {
    fmt.Println(a, b) // what will be printed & why?
}
```

## Star question

```go
type Storage struct {
    cache *lru.Cache
}

func (s *Storage) Set(wh *warehouse.Warehouse) {
    s.cache.Put(s.Id, *wh)
}

func (s *Storage) Get(id types.WarehouseId) *warehouse.Warehouse {
    item, ok := s.cache.Get(id)

    if ok {
        // we do we always get flase?
        if wh, ok := item.(*warehouse.Warehouse); ok { 
            return wh
        }
    }

    return nil
}
```

## Mutate slice in function

```go
var mutate = func(a []int) {
  a[0] = 0
  a = append(a, 1)
  fmt.Println(a) // what will be printed & why?
}

func main() {
  a := []int{1, 2, 3, 4}
  mutate(a)
  fmt.Println(a) // what will be printed & why?
}
```

## Pointer to nil error

```go
type MyError struct{}

func (e *MyError) Error() string {
  return "my error"
}

func CheckError(err error) {
  if err != nil {
    panic(err)
  }
  fmt.Println("ok")
}

func main() {
  var err *MyError
  CheckError(err) // what will be printed & why?
}
```

## Merge slices

```go
func main() {
  out := MergeSlice([]int{1, 2, 5}, []int{4, 5, 6, 7, 8, 9})
  fmt.Println(out)
}

func MergeSlice(list1 []int, list2 []int) []int {
  // implement me
}
```

## Close channel

```go
func ProcessData(data []int) {
    results := make(chan int, len(data))
    for _, val := range data {
        go func(x int) {
            // Complex calculation
            time.Sleep(1 * time.Second)
            results <- x * 2
        }(val)

    }

    for i := 0; i < len(data); i++ {
        fmt.Println(<-results)
    }
}

func main() {
    // what will be printed & why?
    // how to fix it?
    ProcessData([]int{1, 2, 3, 4, 5})
}
```

## Parallel processing

```go
func main() {
    var max int

    for i := 1000; i > 0; i-- {
        go func() {
            if i%2 == 0 && i > max {
                max = i
            }
        }()
    }

    // what will be printed & why?
    // how to fix it?
    fmt.Printf("Maximum is %d", max)
}
```

## Nil map

```go
func main() {
    var m map[string]int

    fmt.Println(m["foo"])

    m["foo"] = 42

    fmt.Println(m["foo"])
}
```

## Mutate struct in function

```go
type Person struct {
    Name string
}

func (person Person) SetName(newName string) {
    person.Name = newName
}

func main() {
    person := Person{
    Name: "Bob",
}

    fmt.Println(person.Name) // what will be printed & why?

    person.SetName("Alice")

    fmt.Println(person.Name) // what will be printed & why?
}
```

## Error printing on painc

```go
func main() {
    // why is the order of printing?
    defer fmt.Println("world")
    fmt.Println("hello")
    panic("error")
}
```

## Select statement

```go
func main() {
        ch1 := make(chan bool)
        ch2 := make(chan bool)
        ch3 := make(chan bool)
        go func() {
                ch1 <- true
        }()
        go func() {
                ch2 <- true
        }()
        go func() {
                ch3 <- true
        }()

        // which channel will be selected?
        select {
        case <-ch1:
                fmt.Printf("val from ch1")
        case <-ch2:
                fmt.Printf("val from ch2")
        case <-ch3:
                fmt.Printf("val from ch3")
        }
}
```

## Channel deadlock

```go
func main() {
  ch := make(chan int)
  go func() {
    for i := 0; i < 5; i++ {
      ch <- i
    }
  }()
  for n := range ch {
    fmt.Println(n)
  }
}
```

## For loop with goroutine

```go
func main() {
  for i := 0; i < 5; i++ {
    go func() {
      // what will be printed & why?
      // how to fix it?
      fmt.Println(i)
    }()
  }
}
```

## Read from 2 channels

```go
func main() {
    timeStart := time.Now()
    _, _ = <-worker(), <-worker()
    // what will be printed & why?
    println(int(time.Since(timeStart).Seconds()))
}

func worker() chan int {
    ch := make(chan int)
    go func() {
        time.Sleep(3 * time.Second)
        ch <- 1
    }()
    return ch
}
```
