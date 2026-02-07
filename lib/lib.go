package lib

import (
	"main/depth"
)

func ENTER_COMMAND() {
	depth.LOG()
}

func HOF_ACCESS_ACTION(funcion func()) {
	if depth.ACCESS_ACTION() {
	}
}

func INIT() {
}
