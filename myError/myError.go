package MyError

import (
	"fmt"
)

//封装错误类型，MyError 类型记录了文件，行号，相关的错误信息
type MyError struct {
	File string
	Line int
	Msg  string
}

//PathError 除了底层错误外还提供了使用哪个文件，执行哪个操作等相关信息。
type PathError struct {
	Path string
	Op   string
	Msg  string
}

func (e *MyError) Error() string {
	//fmt.Errorf(format string, a ...interface{})
	return fmt.Sprintln("%s:%d: %s", e.File, e.Line, e.Msg)
}

func (e *PathError) Error() string {
	return fmt.Sprintln("%s:%s: %s", e.Path, e.Op, e.Msg)
}
