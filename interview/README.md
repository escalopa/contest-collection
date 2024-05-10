## Interview Questions

#### Cyclic channel loop

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

	fmt.Println(<-list[0]) // what will be printed & why?
	fmt.Println(<-list[0]) // what will be printed & why?
}
```

#### Merge channels

```go
package main

import "fmt"

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

  // merge channels into signle channel
  out := mergeChannels(c1, c2, c3)

  // read channels output
  for value := range out {
    fmt.Printf("%d\n", value)
  }

  fmt.Println("Done")
}

func mergeChannels(chans ...<-chan int) <-chan int {
  panic("implement me")
}
```

#### Map ordering

```go
m := map[string]int{"a":1,"b":2,"c":3}

for a, b := range m {
    fmt.Println(a, b) // what will be printed & why?
}
```

#### Star question

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