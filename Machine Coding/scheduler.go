//Problem Statement: Design a scheduler which can schedule tasks based on certain strategies

//Below is a scheduler that schedules job with Delay strategy.
//Since needed granularity is second, so our scheduler sleeps for 1 second and checks if there is any task to schedule.
//If found it schedules all those tasks which are needed to be scheduled at that time and then again goes to sleep for a second

package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"
)

const (
	DELAY = 0
)

var schedulerControl *SchedulerControl

type Task struct {
	name         string
	taskFunc     func()
	addTime      int64
	scheduleTime int64
}

type Scheduler interface {
	AddTask(task Task)
	Start()
}

type DelayScheduler struct {
	taskChannel *chan Task
	tasks       []Task
	mutex       sync.Mutex
}

func (ds *DelayScheduler) AddTask(task Task) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	ds.tasks = append(ds.tasks, task)
	sort.SliceStable(ds.tasks, func(i, j int) bool {
		return ds.tasks[i].scheduleTime > ds.tasks[j].scheduleTime
	})
}

func (ds *DelayScheduler) Start() {
	for {
		for {
			if len(ds.tasks) > 0 {
				task := ds.tasks[len(ds.tasks)-1]
				if task.scheduleTime == time.Now().Unix() {
					*ds.taskChannel <- task
					ds.mutex.Lock()
					ds.tasks = ds.tasks[:len(ds.tasks)-1]
					ds.mutex.Unlock()
				} else {
					break
				}
			}
		}
		time.Sleep(time.Second)
	}
}

type SchedulerControl struct {
	schedulers    map[int]Scheduler
	taskChannel   chan Task
	workers       int
	workerChannel chan func()
}

func GetSchedulerControl(numWorkers int) *SchedulerControl {
	if schedulerControl == nil {
		schedulerControl = &SchedulerControl{
			schedulers:    make(map[int]Scheduler),
			taskChannel:   make(chan Task),
			workers:       numWorkers,
			workerChannel: make(chan func(), 5),
		}
	}
	return schedulerControl
}

func (sc *SchedulerControl) worker(workerChannel chan func()) {
	for work := range workerChannel {
		work()
	}
}

func (sc *SchedulerControl) GetScheduler(schedulerType int) Scheduler {
	if schedulerType == DELAY {
		return &DelayScheduler{
			taskChannel: &sc.taskChannel,
		}
	}
	return nil
}

func (sc *SchedulerControl) CreateTask(schedulerType int, task Task) {
	if _, ok := sc.schedulers[schedulerType]; !ok {
		sc.schedulers[schedulerType] = sc.GetScheduler(schedulerType)
		go sc.schedulers[schedulerType].Start()
	}

	task.scheduleTime = time.Now().Unix() + task.scheduleTime
	sc.schedulers[schedulerType].AddTask(task)
}

func (sc *SchedulerControl) Run() {
	for i := 0; i < sc.workers; i++ {
		go sc.worker(sc.workerChannel)
	}

	for task := range sc.taskChannel {
		sc.workerChannel <- task.taskFunc
	}
}

func main() {
	myScheduler := GetSchedulerControl(3)
	go myScheduler.Run()

	myScheduler.CreateTask(DELAY, Task{name: "Task1", scheduleTime: 8, taskFunc: func() { log.Println("Worker Doing Task1") }})
	myScheduler.CreateTask(DELAY, Task{name: "Task2", scheduleTime: 6, taskFunc: func() { log.Println("Worker Doing Task2") }})
	myScheduler.CreateTask(DELAY, Task{name: "Task3", scheduleTime: 2, taskFunc: func() { log.Println("Worker Doing Task3") }})
	myScheduler.CreateTask(DELAY, Task{name: "Task4", scheduleTime: 5, taskFunc: func() { log.Println("Worker Doing Task4") }})
	myScheduler.CreateTask(DELAY, Task{name: "Task5", scheduleTime: 10, taskFunc: func() { log.Println("Worker Doing Task5") }})

	var a int
	fmt.Scanf("%d", &a)
	log.Println("Done")
}
