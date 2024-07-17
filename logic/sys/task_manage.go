package sys

import (
	"dzhgo/addons/fileUpload/defineType"
	"dzhgo/addons/fileUpload/service"
	"sync"
)

type sTaskManager struct {
	sync.Mutex
	Tasks map[string]*defineType.Task
}

func init() {
	service.RegisterTaskManager(&sTaskManager{
		Tasks: make(map[string]*defineType.Task),
	})
}

// 添加任务
func (s *sTaskManager) AddTask(task *defineType.Task) {

	s.Lock()
	defer s.Unlock()
	s.Tasks[task.Id] = task
}

// 删除任务
func (s *sTaskManager) DelTask(id string) {

	s.Lock()
	defer s.Unlock()
	delete(s.Tasks, id)
}

// 获取任务
func (s *sTaskManager) GetTask(id string) *defineType.Task {

	s.Lock()
	defer s.Unlock()
	return s.Tasks[id]
}

// 更新任务
func (s *sTaskManager) UpdateTask(task *defineType.Task) {

	s.Lock()
	defer s.Unlock()
	s.Tasks[task.Id] = task
}
