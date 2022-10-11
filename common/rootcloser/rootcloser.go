package rootcloser

import "fmt"

var rootClosers []func()

func RegisterCloser(closer func()) {
	rootClosers = append(rootClosers, closer)
}

func Execute() {
	for _, closer := range rootClosers {
		defer recoverRootCloser()
		closer()
	}
}

func recoverRootCloser() {
	if err := recover(); err == nil {
		return
	} else {
		var (
			err error
			ok  bool
		)
		err, ok = err.(error)
		if !ok {
			err = fmt.Errorf("%v", err)
		}
		fmt.Printf("execute root defer failed | err=%s", err.Error())
	}
}
