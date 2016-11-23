package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	N int
}

func main() {
	c0 := make(chan *Value)
	c1 := make(chan *Value)
	go senderRandom(c0, c1)
	go doer(c1, func(v *Value) {
		time.Sleep(2 * time.Second)
		fmt.Println(v)
	})

	for i := 0; i < 20; i++ {
		c0 <- &Value{i}
	}
	close(c0)
	fmt.Println("sent all values")

	select {}
}

func senderLatest(in <-chan *Value, out chan<- *Value) {
	haveValue := false
	var current *Value
	for {
		outc := out
		if !haveValue {
			outc = nil
		}
		select {
		case current, haveValue = <-in:
			if !haveValue {
				close(out)
				return
			}
		case outc <- current:
			haveValue = false
		}
	}
}

func senderAccum(in <-chan *Value, out chan<- *Value) {
	var vals []*Value
	var closed bool
	for {
		var sendv *Value
		outc := out
		if len(vals) == 0 {
			if closed {
				close(out)
				return
			}
			outc = nil
		} else {
			sendv = vals[0]
		}
		select {
		case v, ok := <-in:
			if !ok {
				closed = true
				break
			}
			vals = append(vals, v)
		case outc <- sendv:
			vals = vals[1:]
		}
	}
}

func senderRandom(in <-chan *Value, out chan<- *Value) {
	var wg sync.WaitGroup
	for v := range in {
		v := v
		wg.Add(1)
		go func() {
			out <- v
			wg.Done()
		}()
	}
	wg.Wait()
	close(out)
}

func doer(in <-chan *Value, f func(*Value)) {
	for v := range in {
		f(v)
	}
}

func process(values []V) {
	results := make([]R, len(values))
	
	run := parallel.NewRun(30)
	for i, v := range values {
		i, v := i, v
		run.Do(func() error {
			r, err := process(v)
			if err != nil {
				return err
			}
			results[i] = r
			return nil
		})
	}
	if err := run.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}
