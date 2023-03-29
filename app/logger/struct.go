package logger

type logger struct {
	LogFile     string `json:"log_file"`
	LogLevel    string `json:"log_level"`
	LogToStdout bool   `json:"to_stdout"`
	LogToStderr bool   `json:"to_stderr"`
	LogToFile   bool   `json:"to_file"`
}
