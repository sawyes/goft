package goft

import "sync"

//goft-task

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor //任务列表
var once sync.Once

func init() {
	chlist := getTaskList() //得到任务列表
	go func() {
		for t := range chlist {
			doTask(t)
		}
	}()
}

func doTask(t *TaskExecutor) {
	go func(t *TaskExecutor) {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		
		t.Exec() //执行任务
	}(t)
}

func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor) //初始化
	})
	return taskList
}

type TaskExecutor struct {
	f TaskFunc
	p []interface{}
	callback func()
}

func NewTaskExecutor(f TaskFunc, callback func(), p []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

func (this *TaskExecutor) Exec() { //执行任务
	this.f(this.p...)
}

func Task(f TaskFunc, callback func(), params ...interface{}) {
	go func(f TaskFunc, params ...interface{}) {
		getTaskList() <- NewTaskExecutor(f, callback, params) //增加任务队列
	}(f, params)
}
