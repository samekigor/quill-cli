package exitcodes

const (
	Success                = iota
	UnableToLoadConfigFile // 1
	FileNotFound           // 2
	InvalidUserInput       // 3
)
