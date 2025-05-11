## Questions

---

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

### system design: database scaling

```
You have a database that stores user data.
Multiple microservices read and write to it in various combinations.
After scaling up the number of services and queries, the database can no longer handle the load.
Suggest ways to scale the database and resolve performance issues.
```
