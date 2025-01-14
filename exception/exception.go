package exception

import (
	"fmt"
	log "github.com/spance/deblocus/golang/glog"
	"runtime"
)

// injection
var DEBUG bool

type Exception struct {
	msg     string
	code    int
	warning bool
}

func (e *Exception) Error() string {
	return e.msg
}

func (e *Exception) Code() int {
	return e.code
}

func (e *Exception) Warning() bool {
	return e.warning
}

func (e *Exception) Apply(appendage interface{}) *Exception {
	newE := new(Exception)
	newE.code = e.code
	newE.msg = fmt.Sprintf("%s %v", e.msg, appendage)
	return newE
}

func NewW(msg string) *Exception {
	return &Exception{msg: msg, warning: true}
}

func New(code int, msg string) *Exception {
	return &Exception{msg: msg, code: code}
}

func CatchException(e interface{}) bool {
	if ex, y := e.(*Exception); y && ex.warning {
		log.Errorln(ex.msg)
		return true
	} else if e != nil {
		if DEBUG {
			buf := make([]byte, 1600)
			runtime.Stack(buf, false)
			fmt.Println(e)
			fmt.Println(string(buf))
		} else {
			if s, y := e.(string); y {
				log.Warningln(s)
			} else {
				log.Errorln(e)
			}
		}
		return true
	}
	return false
}
