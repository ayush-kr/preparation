//Mini Jira platform to manage tasks, users, sprint, etc.
//Not very robust but can be made with little modifications

package main

import (
	"errors"
	"log"
)

var uuid int

type Tasker interface {
	SetStatus(status string)
	SetAssignee(assigneeid string)
	GetTaskID() int
}

//----------Task-------------

type Task struct {
	taskID     int
	taskType   string
	taskStatus string
	assigneeID string
}

func (t *Task) SetStatus(status string) {
	t.taskStatus = status
	log.Println("Status of", t.taskType, t.taskID, "changed to", status)
}

func (t *Task) SetAssignee(userID string) {

	previousAssignee, ok := assignees[t.assigneeID]
	if ok {
		delete(previousAssignee.assignedTasks, t.taskID)
	}

	log.Println("Assignee of", t.taskType, t.taskID, "changed to", userID)
	t.assigneeID = userID

	assignees[userID].assignedTasks[t.taskID] = t

}

func (t *Task) GetTaskID() int {
	return t.taskID
}

//-------Task end----------------------

//-------TaskFactory-------------------

type TaskFactory struct{}

func (tf *TaskFactory) CreateTask(taskType string) (Tasker, error) {
	var task *Task
	if taskType == "STORY" {
		uuid++
		task = &Task{
			taskType:   "Story",
			taskStatus: "Open",
			taskID:     uuid,
		}
	} else if taskType == "BUG" {
		uuid++
		task = &Task{
			taskType:   "Bug",
			taskStatus: "Open",
			taskID:     uuid,
		}
	} else if taskType == "FEATURE" {
		uuid++
		task = &Task{
			taskType:   "Feature",
			taskStatus: "Open",
			taskID:     uuid,
		}
	}
	if task != nil {
		tasks[uuid] = task
		return task, nil
	}
	return nil, errors.New("Invalid Task")
}

//-------TaskFactory end-------------------

//-------Assigne(Users)-------------------

type Assignee struct {
	userID        string
	assignedTasks map[int]*Task
}

func (a *Assignee) GetTasks() {
	log.Println("Tasks assigned to", a.userID, "are:")
	for key, val := range a.assignedTasks {
		log.Println(val.taskType, key, ": status = ", val.taskStatus)
	}
	log.Println()
}

//-------Assigne end-------------------

//-------Common-------------------

var assignees map[string]*Assignee
var tasks map[int]*Task

func GetTasksOfAssignees() {
	for _, val := range assignees {
		val.GetTasks()
	}
}

func CreateUser(userID string) *Assignee {
	assignee := &Assignee{
		userID:        userID,
		assignedTasks: make(map[int]*Task, 25),
	}
	assignees[assignee.userID] = assignee
	return assignee
}

func CreateSprint(sprintID, description string) *Sprint {
	sprint := &Sprint{
		sprintid:    "Sprint1",
		description: "Macha denge is sprint me, Bolt ka record todna h",
		tasks:       make(map[int]*Task, 500),
	}
	return sprint
}

//-------Common end---------------

//-------Sprint-------------------

type Sprint struct {
	sprintid    string
	description string
	tasks       map[int]*Task
}

func (sprint *Sprint) AddTask(taskID int) {
	task := tasks[taskID]
	sprint.tasks[taskID] = task
}

func (sprint *Sprint) RemoveTask(taskID int) {
	delete(sprint.tasks, taskID)
}

func (sprint *Sprint) GetTasks() {
	log.Println("Tasks in sprint", sprint.sprintid, "are:")
	for key, val := range sprint.tasks {
		log.Println(val.taskType, key, ": status = ", val.taskStatus, ", Assigned to:", val.assigneeID)
	}
}

//-------Sprint end---------------

func main() {
	assignees = make(map[string]*Assignee, 20)
	tasks = make(map[int]*Task, 500)

	CreateUser("ayush")
	CreateUser("lucy")

	uuid = 0
	taskFactory := &TaskFactory{}
	story1, _ := taskFactory.CreateTask("STORY")
	story1.SetAssignee("ayush")

	story2, _ := taskFactory.CreateTask("STORY")
	story2.SetAssignee("lucy")

	bug1, _ := taskFactory.CreateTask("BUG")
	bug1.SetAssignee("ayush")

	bug2, _ := taskFactory.CreateTask("BUG")
	bug2.SetAssignee("lucy")

	sprint := CreateSprint("Sprint1", "Macha denge is sprint me. Bolt ka record todna hai")

	sprint.AddTask(story1.GetTaskID())
	sprint.AddTask(story2.GetTaskID())
	sprint.AddTask(bug1.GetTaskID())
	sprint.AddTask(bug2.GetTaskID())
	sprint.GetTasks()

	GetTasksOfAssignees()

	story1.SetAssignee("lucy")
	story2.SetAssignee("ayush")
	bug1.SetAssignee("lucy")
	bug2.SetAssignee("ayush")

	GetTasksOfAssignees()

	story1.SetStatus("InProgress")
	story2.SetStatus("Done")
	bug1.SetStatus("Blocked")
	bug2.SetStatus("InProgress")

	GetTasksOfAssignees()
	story1.SetStatus("Done")
	bug1.SetStatus("Done")
	bug2.SetStatus("Done")

	GetTasksOfAssignees()

	sprint.GetTasks()
}
