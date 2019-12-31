package add_numbers_mutex_sync

import (
	"fmt"
	"sync"
)

var numbers []int

func AddNumbers(){
	mutex:= &sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		mutex.Lock()
		fmt.Println("#1 locked")
		numbers = append(numbers, 0, 1, 2, 3, 4)
		mutex.Unlock()
		fmt.Println("#1 unlocked")
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		fmt.Println("#2 locked")
		numbers = append(numbers, 5, 6, 7, 8, 9)
		mutex.Unlock()
		fmt.Println("#2 unlocked")
	}()
	wg.Wait()
	fmt.Printf("%+v", numbers)
}
