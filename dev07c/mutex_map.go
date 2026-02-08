package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mx sync.Mutex
	m map[string]string
}

func main() {
	{
		// Base Mutex
		// Мютекс используется для защиты мап от параллельного доступа и гонки данных
		var wg sync.WaitGroup
		wg.Add(3)

		cache := Cache{
			m: make(map[string]string),
		}
		
		cache.m["key"] = "abc"

		go func() {
			defer wg.Done()
			cache.mx.Lock()
			cache.m["key"] = "1abc"
			fmt.Println("Go write 1:", cache.m["key"])
			cache.mx.Unlock()
		}()

		go func() {
			defer wg.Done()
			cache.mx.Lock()
			cache.m["key"] = "2abc"
			fmt.Println("Go write 2:", cache.m["key"])
			cache.mx.Unlock()
		}()

		go func() {
			defer wg.Done()
			cache.mx.Lock()
			cache.m["key"] = "3abc"
			fmt.Println("Go write 3:", cache.m["key"])
			cache.mx.Unlock()
		}()
		
		wg.Wait()
		fmt.Println("Cache:", cache.m)
		fmt.Println("Main exit")
	}
}