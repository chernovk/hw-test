package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// func Run(tasks []Task, n, m int) error {
// 	mu := sync.Mutex{}
// 	var errCount int32
// 	var result error

// 	wg := sync.WaitGroup{}

// 	for i := 0; i < n; i++ {
// 		wg.Add(1)
// 		go func() error {
// 			defer wg.Done()
// 			for {
// 				mu.Lock()
// 				if len(tasks) <= 0 {
// 					mu.Unlock()
// 					break
// 				}
// 				task := tasks[0]
// 				tasks = tasks[1:]
// 				mu.Unlock()

// 				if int(atomic.LoadInt32(&errCount)) >= m {
// 					mu.Lock()
// 					if result == nil {
// 						result = ErrErrorsLimitExceeded
// 					}
// 					mu.Unlock()
// 					return result
// 				}

// 				if err := task(); err != nil {
// 					atomic.AddInt32(&errCount, 1)
// 				}
// 			}
// 			return result
// 		}()
// 	}
// 	wg.Wait()
// 	return result
// }

func Run(tasks []Task, n, m int) error {
	mu := sync.Mutex{}
	var errCount int32
	var result error
	taskCh := make(chan Task, len(tasks))
	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)

	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() error {
			defer wg.Done()
			for task := range taskCh {
				if int(atomic.LoadInt32(&errCount)) >= m {
					mu.Lock()
					if result == nil {
						result = ErrErrorsLimitExceeded
					}
					mu.Unlock()
					break
				}

				if err := task(); err != nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
			return result
		}()
	}

	wg.Wait()
	return result
}
