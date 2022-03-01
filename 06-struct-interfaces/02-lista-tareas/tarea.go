package main

import (
	"fmt"
)

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}
func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)

}

type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripción: %s \nCompeltado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}
func main() {
	t1 := Task{
		name:      "Curso de Go",
		desc:      "Completar el curso este mes",
		completed: false,
	}
	t2 := Task{
		name:      "Curso de HTM",
		desc:      "Completar el curso esta semana",
		completed: true,
	}
	lista := TaskList{}
	lista.appendTask(&t1)
	lista.appendTask(&t2)
	t1.toPrint()
	t2.toPrint()

	t3 := Task{
		name:      "Curso de SCC",
		desc:      "Completar el curso esta semana",
		completed: false,
	}
	lista.appendTask(&t3)
	lista.removeTask(1)
	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}
}
