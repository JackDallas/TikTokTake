# TikTokTracker

The central server that queues TikTok archive jobs and distributes them to the crawlers

## REST

### GET

- /status => Returns current stats on the tracker
- /job/{id} => Returns the status of the job with the id {id}

### POST

- /job/add => Add's a job, json with job details must be posted

## Socket

WIP

## R1

- Add job to queue
- Pull jobs from queue and send them to crawlers

## R2

- DB with list of past urls for filtering
- Job status tracking
- Persistance

## Workings

/job/add      AddJob(job)   
HTTPServer => Tracker    => SocketServer
