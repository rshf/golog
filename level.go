package golog

type level int

const (
	ALL level = iota * 10
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	SQL
	NONE
	up
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
	case 70:
		return "SQL"
	default:
		return "DEBUG"
	}
}
