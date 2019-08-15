package engine

import (
	"log"
	"youtube/interfaces"
	"youtube/scheduler"
)

type SimpleEngine struct {
	scheduler   *scheduler.SimpleScheduler
	workerCount int
}

func NewSimpleEngine(count int) *SimpleEngine {
	return &SimpleEngine{workerCount: count}
}


func (e SimpleEngine) Run(tasks ...interfaces.Task) {
	in := make(chan interfaces.Task)
	out := make(chan interfaces.Task)

	for i := 0; i < e.workerCount; i++ {
		createWorker(in, out)
	}

	e.scheduler = scheduler.NewSimpleScheduler(in)
	for _, task := range tasks {
		e.scheduler.Submit(task)
	}
	
	count := 0
	for {
		newTask := <-out
		if newTask != nil {
			e.scheduler.Submit(newTask)
			continue
		}
		count++
		if count == len(tasks) {
			break
		}
	}
}

func createWorker(get chan interfaces.Task, post chan interfaces.Task) {
	go func() {
		for {
			task := <- get
			task.Report()
			err, nextTask := task.Execute()
			if err != nil {
				log.Println(err)
			}
			post <- nextTask
		}
	}()
}





