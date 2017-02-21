package cmd

type Result struct {
	Subdir   string
	ExitCode int
	Output   []byte
}


func (o Result) IsSuccess() bool {
	return o.ExitCode == 0
}