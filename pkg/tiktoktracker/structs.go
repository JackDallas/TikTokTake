package tiktoktracker

const (
	//Queued Job
	Queued = 1
	//InProgress Job
	InProgress = 2
	//Done Job
	Done = 3

	//Error with Job
	Error = 4
)

//TrackerJob : Represents a tracker job
type TrackerJob struct {
	Username string
	ID       string
	Status   int
}
