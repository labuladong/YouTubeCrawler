package interfaces

type Task interface {
	Execute() (error, Task)
	Report()
}
