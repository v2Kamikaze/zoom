package background

import (
	"log"
	"sync"
	"time"
)

type Task func()

type BackgroundRunner struct {
	wg    *sync.WaitGroup
	tasks []Task
}

func NewBackgroundRunner() *BackgroundRunner {
	return &BackgroundRunner{
		wg:    &sync.WaitGroup{},
		tasks: []Task{},
	}
}

func (t *BackgroundRunner) Add(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *BackgroundRunner) RunAndWait() {

	rec := func(taskNumber int) {
		if r := recover(); r != nil {
			log.Printf("erro ao processar tarefa '%d'\n", taskNumber)
		}
	}

	for i := range t.tasks {
		t.wg.Add(1)
		go func() {
			now := time.Now()
			defer t.wg.Done()
			defer rec(i)

			t.tasks[i]()

			log.Printf("Tarefa %d demorou '%f' segundos\n", i, time.Since(now).Seconds())
		}()
	}

	t.wg.Wait()
}

func (t *BackgroundRunner) Clear() {
	t.tasks = nil
}
