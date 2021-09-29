package utils

import "errors"

func VerfiyQuery(strSlice ...string) error {
	for _, s := range strSlice {
		if s == "" {
			return errors.New("参数缺失")
		}
	}
	return nil
}
