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
