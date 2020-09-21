package golog

type level int

const (
	All level = iota * 10
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var Level level = INFO

func (l level) String() string {
	switch l {
	case 10:
		return "TRACE"
	case 20:
		return "DEBUG"
	case 30:
		return "INFO"
	case 40:
		return "WARN"
	case 50:
		return "ERROR"
	case 60:
		return "FATAL"
	default:
		return "INFO"
	}
}
