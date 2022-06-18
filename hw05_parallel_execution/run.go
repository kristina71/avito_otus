package hw05parallelexecution

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrInvalidGoroutineCount = errors.New("invalid goroutine count given")

type Task func() error

func (t Task) Run() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("something went wrong, panic %v", p)
		}
	}()

	return t()
}

func validateError(errCount, errLimit int) error {
	log.Printf("err received %d/%d", errCount, errLimit)
	if errLimit >= 0 && errCount >= errLimit {
		log.Printf("err limit reached, terminate")
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(i int, chanTasks <-chan Task, chanErrors chan<- error, wg *sync.WaitGroup) {
	log.Printf("[%d] start worker", i)
	defer wg.Done()
	for t := range chanTasks {
		log.Printf("[%d] task received", i)
		if err := t.Run(); err != nil {
			log.Printf("got and [%d] error", i)
			chanErrors <- err
		}
	}
	log.Printf("[%d] terminate worker", i)
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 1 {
		return fmt.Errorf("%w: expected > 0, actual %d", ErrInvalidGoroutineCount, n)
	}

	var errCount int
	chanErrors := make(chan error, n)
	defer close(chanErrors)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	chanTasks := make(chan Task)
	defer close(chanTasks)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go worker(i, chanTasks, chanErrors, wg)
	}

	i := 0
	for i < len(tasks) {
		select {
		case <-chanErrors:
			errCount++
			if err := validateError(errCount, m); err != nil {
				return err
			}
		default:
		}

		select {
		case <-chanErrors:
			errCount++
			if err := validateError(errCount, m); err != nil {
				return err
			}
		case chanTasks <- tasks[i]:
			i++
			log.Printf("task %d pushed", i)
		}
	}
	return nil
}
