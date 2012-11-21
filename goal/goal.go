package goal

type Status int

const (
	StatusTodo = iota
	StatusInprogress
	StatusDone
	StatusBacklog
)

func (status Status) String() string {
	statusStrings := map[Status]string{
		StatusTodo:       "Todo",
		StatusInprogress: "In progress",
		StatusDone:       "Done",
		StatusBacklog:    "Backlog",
	}

	return statusStrings[status]
}

//type Goal struct {
//	title       string
//	description string
//	status      Status
//}
