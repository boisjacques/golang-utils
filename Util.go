package util

import (
	"log"
	"runtime"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Tracer() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return file + " " + string(line) + " " + f.Name()
}
