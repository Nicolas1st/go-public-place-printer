package jobq

import (
	"fmt"
	"os/exec"
	"printer/persistence/model"
	"strings"
)

type JobQueue struct {
	jobs map[model.JobID]*model.Job
}

// NewJobQueue - создать новую очередь задач
func NewJobQueue() *JobQueue {
	return &JobQueue{
		jobs: map[model.JobID]*model.Job{},
	}
}

// Enqueue - добавляет работу к очереди
func (q *JobQueue) Push(job *model.Job) model.JobID {
	// отправить файл на принтер
	cmd := fmt.Sprintf("lp %v | cut -d ' ' -f 4", job.StoredFilePath)
	command := exec.Command("bash", "-c", cmd)
	idOut, _ := command.CombinedOutput()

	// убрать \n
	jobID := model.JobID(strings.Replace(string(idOut), "\n", "", -1))

	// сохранить работу
	job.ID = jobID
	q.jobs[jobID] = job

	return jobID
}

// Cancel - отменить работу
func (q *JobQueue) CancelJob(jobID model.JobID) {
	exec.Command("cancel", string(jobID)).Run()
	delete(q.jobs, jobID)
}

// GetAllJobs - вернуть все работы на принтере
func (q *JobQueue) GetAllJobs() map[model.JobID]*model.Job {
	return q.jobs
}
