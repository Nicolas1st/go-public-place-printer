package jobq

import (
	"fmt"
	"os/exec"
	"printer/persistence/model"
	"strconv"
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
	cmd := fmt.Sprintf(`lp "%v" | cut -d ' ' -f 4`, job.StoredFilePath)
	command := exec.Command("bash", "-c", cmd)
	out, _ := command.CombinedOutput()

	// убрать \n
	jobID := model.JobID(strings.Replace(string(out), "\n", "", -1))

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

// RemoveCompletedAndCanceledJobs - удалить завершенные и отменненые работы
func (q *JobQueue) RemoveCompletedAndCanceledJobs() {
	for _, job := range q.jobs {
		cmd := fmt.Sprintf(`lpstat  | grep %v | wc -l`, job.ID)
		command := exec.Command("bash", "-c", cmd)
		out, _ := command.CombinedOutput()
		numberOfJobsWithIDProvided, _ := strconv.Atoi(string(out))
		if numberOfJobsWithIDProvided == 0 {
			q.CancelJob(job.ID)
		}
	}
}
