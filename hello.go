package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := time.Now()
	search("test")
	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func search(keyword string) {
	folders := [3]string{"one", "two", "three"}
	var wg sync.WaitGroup
	wg.Add(len(folders))
	defer wg.Wait()
	for _, folder := range folders {
		go searchFromFolder(keyword, folder, &wg)
	}
}

func searchFromFolder(keyword string, folder string, wg *sync.WaitGroup) {
	fmt.Println("Finding in " + folder)
	time.Sleep(1 * time.Second)
	wg.Done()
}
