package logs

//Level ..
type Level int8

const (
	// LevelDebug ..
	LevelDebug Level = iota
	//LevelInfo ..
	LevelInfo
	//LevelWarn ..
	LevelWarn
	//LevelError ..
	LevelError
)

const (
	//LevelKey ..
	LevelKey = "level"
)

func (l Level) Enabled(lv Level) bool {
	return lv >= l
}

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return ""
	}
}
