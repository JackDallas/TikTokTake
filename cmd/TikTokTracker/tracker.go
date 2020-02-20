package main

import (
	"errors"
	"sync"

	ttt "github.com/JackDallas/TikTokTake/pkg/tiktoktracker"
)

//Tracker :
type Tracker struct {
	jobQueue []ttt.TrackerJob

	sync.RWMutex
}

//NewTracker : Creates a new Tracker
func NewTracker() (Tracker, error) {
	tracker := Tracker{}
	//TODO DB Stuff
	return tracker, nil
}

//AddJob : Queues a TrackerJob
func (tracker *Tracker) AddJob(job ttt.TrackerJob) {
	tracker.Lock()
	defer tracker.Unlock()

	job.Status = ttt.Queued
	tracker.jobQueue = append(tracker.jobQueue, job)
}

//GetNextJob : Gets the next Job off the tracker and marks it as in progress
func (tracker *Tracker) GetNextJob() (*ttt.TrackerJob, error) {
	tracker.Lock()
	defer tracker.Unlock()

	for i := (len(tracker.jobQueue) - 1); i >= 0; i++ {
		if tracker.jobQueue[i].Status == ttt.Queued {
			tracker.jobQueue[i].Status = ttt.InProgress
			return &tracker.jobQueue[i], nil
		}
	}
	return nil, errors.New("No Job's Avaliable")
}

//GetJob : Gets a specified Job from the tracker, to be used to check status
func (tracker *Tracker) GetJob(id string) (*ttt.TrackerJob, error) {
	tracker.RLock()
	defer tracker.RUnlock()

	for i := range tracker.jobQueue {
		if tracker.jobQueue[i].ID == id {
			return &tracker.jobQueue[i], nil
		}
	}
	return nil, errors.New("No Job Found")
}
