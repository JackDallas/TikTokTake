package main

import "errors"

//SocketServer :
type SocketServer struct {
	Tracker *Tracker
	Port    int
	started bool
}

//NewSocketServer : Creates and starts the SocketServer, returns an error if the server has already started or cannot start
func NewSocketServer(t *Tracker, p int) (SocketServer, error) {
	ss := SocketServer{}
	ss.Tracker = t
	ss.Port = p

	err := ss.init()
	if err != nil {
		return SocketServer{}, err
	}

	err = ss.start()
	if err != nil {
		return SocketServer{}, err
	}

	return ss, nil
}

func (ss *SocketServer) init() error {
	ss.started = false
	return nil
}

func (ss *SocketServer) start() error {
	if !ss.started {
		ss.started = true

	} else {
		return errors.New("SocketServer already running")
	}
	return nil
}
