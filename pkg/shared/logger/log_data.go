package logger

type LogData struct {
	ID        int64  `json:"id"`
	CreatedOn int64  `json:"created_on"`
	Level     string `json:"level"`
	FuncName  string `json:"func_name"`
	FileName  string `json:"file_name"`
	Line      int64  `json:"line"`
	Time      string `json:"time"`
	Message   string `json:"message"`
}
