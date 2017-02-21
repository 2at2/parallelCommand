package cmd

import (
	"github.com/codeskyblue/go-sh"
	"fmt"
)

func Exec(chanel chan Result, path, subdir string, command string, args ...interface{}) {
	bytes, err := sh.NewSession().SetDir(path + "/" + subdir).Command(command, args...).Output()

	fmt.Println(subdir)
	fmt.Println(err)
	fmt.Println(string(bytes))

	if err != nil {
		chanel <- Result{
			Subdir: subdir,
			ExitCode: 1,
			Output: bytes,
		}
	} else {
		chanel <- Result{
			Subdir: subdir,
			ExitCode: 0,
			Output: bytes,
		}
	}
}
