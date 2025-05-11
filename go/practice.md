<!-- TOC start (generated with https://github.com/derlin/bitdowntoc) -->

- [cyclic channel loop](#cyclic-channel-loop)
- [merge channels](#merge-channels)
- [map ordering](#map-ordering)
- [star question](#star-question)
- [mutate slice in function](#mutate-slice-in-function)
- [pointer to nil error](#pointer-to-nil-error)
- [merge slices](#merge-slices)
- [close channel](#close-channel)
- [parallel processing](#parallel-processing)
- [nil map](#nil-map)
- [mutate struct in function](#mutate-struct-in-function)
- [error printing on painc](#error-printing-on-painc)
- [select statement](#select-statement)
- [channel deadlock](#channel-deadlock)
- [for loop with goroutine](#for-loop-with-goroutine)
- [read from 2 channels](#read-from-2-channels)
- [nil response body](#nil-response-body)
- [slice modification and append behavior](#slice-modification-and-append-behavior)
- [defer behavior](#defer-behavior)
- [defer in for loop](#defer-in-for-loop)
- [goroutines, waitgroup, and execution order](#goroutines-waitgroup-and-execution-order)
- [race condition on map write](#race-condition-on-map-write)
- [merge multiple channels into one (2nd edition)](#merge-multiple-channels-into-one-2nd-edition)
- [system design: database scaling](#system-design-database-scaling)

<!-- TOC end -->

<!-- TOC --><a name="cyclic-channel-loop"></a>
### cyclic channel loop

```go
package main

import "fmt"

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

  fmt.Println(<-list[0])
  fmt.Println(<-list[0])
}
```

---

<!-- TOC --><a name="merge-channels"></a>
### merge channels

```go
package main

import "fmt"

func main() {
  var (
    c1 = make(chan int)
    c2 = make(chan int)
    c3 = make(chan int)
  )

  go func(){ c1 <- 1; close(c1) }()
  go func(){ c2 <- 2; close(c2) }()
  go func(){ c3 <- 3; close(c3) }()

  out := mergeChannels(c1, c2, c3)

  for value := range out {
    fmt.Printf("%d\n", value)
  }

  fmt.Println("Done")
}

func mergeChannels(chans ...<-chan int) <-chan int {
  panic("implement me")
}
```

---

<!-- TOC --><a name="map-ordering"></a>
### map ordering

```go
package main

import "fmt"

func main () {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	for a, b := range m {
		fmt.Println(a, b)
	}
}
```

---

<!-- TOC --><a name="star-question"></a>
### star question

```go
package main

type Storage struct {
    cache *Cache
}

func (s *Storage) Set(wh *warehouse.Warehouse) {
    s.cache.Put(wh.Id, *wh)
}

func (s *Storage) Get(id types.WarehouseId) *warehouse.Warehouse {
    item, ok := s.cache.Get(id)

    if ok {
        if wh, ok := item.(*warehouse.Warehouse); ok { 
            return wh
        }
    }

    return nil
}
```

---

<!-- TOC --><a name="mutate-slice-in-function"></a>
### mutate slice in function

```go
package main

import "fmt"

var mutate = func(a []int) {
  a[0] = 0
  a = append(a, 1)
  fmt.Println(a)
}

func main() {
  a := []int{1, 2, 3, 4}
  mutate(a)
  fmt.Println(a)
}
```

---

<!-- TOC --><a name="pointer-to-nil-error"></a>
### pointer to nil error

```go
package main

import "fmt"

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
  CheckError(err)
}
```

---

<!-- TOC --><a name="merge-slices"></a>
### merge slices

```go
package main

import "fmt"

func main() {
  out := MergeSlice([]int{1, 2, 5}, []int{4, 5, 6, 7, 8, 9})
  fmt.Println(out)
}

func MergeSlice(list1 []int, list2 []int) []int {
  panic("implement me")
}
```

---

<!-- TOC --><a name="close-channel"></a>
### close channel

```go
package main

import (
	"fmt"
	"time"
)

func ProcessData(data []int) {
    results := make(chan int, len(data))
    for _, val := range data {
        go func(x int) {
            time.Sleep(1 * time.Second)
            results <- x * 2
        }(val)
    }

    for i := 0; i < len(data); i++ {
        fmt.Println(<-results)
    }
}

func main() {
    ProcessData([]int{1, 2, 3, 4, 5})
}
```

---

<!-- TOC --><a name="parallel-processing"></a>
### parallel processing

```go
package main

import "fmt"

func main() {
    var m int

    for i := 1000; i > 0; i-- {
        go func() {
            if i%2 == 0 && i > m {
                m = i
            }
        }()
    }

    fmt.Printf("Maximum is %d", m)
}
```

---

<!-- TOC --><a name="nil-map"></a>
### nil map

```go
package main

import "fmt"

func main() {
    var m map[string]int

    fmt.Println(m["foo"])

    m["foo"] = 42

    fmt.Println(m["foo"])
}
```

---

<!-- TOC --><a name="mutate-struct-in-function"></a>
### mutate struct in function

```go
package main

import "fmt"

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

    fmt.Println(person.Name)

    person.SetName("Alice")

    fmt.Println(person.Name)
}
```

---

<!-- TOC --><a name="error-printing-on-painc"></a>
### error printing on painc

```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
    panic("error")
}
```

---

<!-- TOC --><a name="select-statement"></a>
### select statement

```go
package main

import "fmt"

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

---

<!-- TOC --><a name="channel-deadlock"></a>
### channel deadlock

```go
package main

import "fmt"

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

---

<!-- TOC --><a name="for-loop-with-goroutine"></a>
### for loop with goroutine

```go
package main

import "fmt"

func main() {
  for i := 0; i < 5; i++ {
    go func() {
      fmt.Println(i)
    }()
  }
}
```

---

<!-- TOC --><a name="read-from-2-channels"></a>
### read from 2 channels

```go
package main

import (
	"time"
)

func main() {
    timeStart := time.Now()
    _, _ = <-worker(), <-worker()
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

---

<!-- TOC --><a name="nil-response-body"></a>
### nil response body

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	client := http.Client{}
	u, _ := url.Parse("https://no.such.host")
	request := &http.Request{
		Method: "GET",
		URL:    u,
	}

	response, err := client.Do(request)
	defer func() { _ = response.Body.Close() }()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Request failed: %s\n", err.Error())
	}

	_, _ = io.Copy(os.Stdout, response.Body)
}
```

---

<!-- TOC --><a name="slice-modification-and-append-behavior"></a>
### slice modification and append behavior

```go
package main

import (
    "fmt"
)

func mod(a []int) {
    // a = append(a, 125)

    for i := range a {
        a[i] = 5
    }

    fmt.Println(a) 
}

func main() {
    sl := []int{1, 2, 3, 4}
    mod(sl)
    fmt.Println(sl)
}
```

---

<!-- TOC --><a name="defer-behavior"></a>
### defer behavior

```go
package main

import (
    "fmt"
)

func main() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

---

<!-- TOC --><a name="defer-in-for-loop"></a>
### defer in for loop

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    wg := new(sync.WaitGroup)

    for i := 0; i < 10; i++ {
        wg.Add(1)
        defer fmt.Println(i)
    		defer func() { fmt.Println(i) }()
        wg.Done()
    }

    wg.Wait()
}
```

<!-- TOC --><a name="goroutines-waitgroup-and-execution-order"></a>
### goroutines, waitgroup, and execution order

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        time.Sleep(2 * time.Second)
        fmt.Println("1")
        wg.Done()
    }()

    wg.Wait()

    go func() {
        fmt.Println("2")
    }()

    fmt.Println("3")
}
```

<!-- TOC --><a name="race-condition-on-map-write"></a>
### race condition on map write

```go
package main

import (
    "fmt"
    "time"
)

var (
    m  map[string]int
)

func main() {
    m = make(map[string]int)
    go f1()
    go f2()

    time.Sleep(time.Second * 1)
    fmt.Printf("print map = %+v\n", m)
}

func f1() {
    for i := 0; i < 100000; i++ {
        m["f1"]++
    }
}

func f2() {
    for i := 0; i < 100000; i++ {
        m["f2"]++
    }
}
```

---

<!-- TOC --><a name="merge-multiple-channels-into-one-2nd-edition"></a>
### merge multiple channels into one (2nd edition)

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func RandNumbers(length, m int) []int {
    var s []int
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < length; i++ {
        s = append(s, rand.Intn(m))
    }
    return s
}

func writeToChan(ch chan<- int) {
    defer close(ch)
    for _, v := range RandNumbers(100, 100) {
        ch <- v
    }
}

func mergeChan(chs ...chan int) chan int {
    panic("implement me")
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
    ch4 := make(chan int)

    mergedChan := mergeChan(ch1, ch2, ch3, ch4)

    go writeToChan(ch1)
    go writeToChan(ch2)
    go writeToChan(ch3)
    go writeToChan(ch4)

    for v := range mergedChan {
        fmt.Println(v)
    }
}
```

---

<!-- TOC --><a name="system-design-database-scaling"></a>
### system design: database scaling

```
You have a database that stores user data.
Multiple microservices read and write to it in various combinations.
After scaling up the number of services and queries, the database can no longer handle the load.
Suggest ways to scale the database and resolve performance issues.
```
