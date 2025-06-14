package util

import (
	"errors"
	"strconv"

	"github.com/nanorand/nanorand"
)

func GenerateOTP() (int64, error) {
	code, err := nanorand.Gen(6)
	if err != nil {
		return 0,errors.New("could not generate OTP")
	}

	convertedCode,err := strconv.ParseInt(code, 10, 64)
	if err != nil {
		return 0,errors.New("could not convert OTP")
	}

	return convertedCode,nil
}