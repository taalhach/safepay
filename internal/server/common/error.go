package common

import (
	"github.com/labstack/echo"
)

// CheckError helps handling errors
func CheckError(err error) {
	if err == nil {
		return
	}

	if _, ok := err.(*echo.HTTPError); ok {
		panic(err)
	}

	AbortErr(err)
}

func AbortErr(err error) {
	ret := echo.NewHTTPError(500)
	ret.SetInternal(err)
	panic(ret)
}
