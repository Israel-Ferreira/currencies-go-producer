package jobs


type Job interface {
	configure()
	RunTask()
}