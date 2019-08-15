package interfaces

type Scheduler interface {
	Submit(Task)
	SetChannel(chan<- Task)
}
