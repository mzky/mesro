package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	//	"sync"
	"time"
)

func main() {
	cr("one")
	cr("two")
	cr("three")
	select {}
}

func cr(str string) {

	//var wg sync.WaitGroup
	//wg.Add(1)
	c := gron.New()
	c.AddFunc(gron.Every(1*time.Second), func() {
		fmt.Println("+runs " + str)
	})
	c.AddFunc(gron.Every(1*time.Second), func() {
		fmt.Println("-runs " + str)
	})
	c.AddFunc(gron.Every(10*time.Second), func() {
		//	defer wg.Done()
		c.Stop()
	})
	c.Start()
	//wg.Wait()
}
