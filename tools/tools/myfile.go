package tools

import (
	"os"
)


func DirCreate(path string) (err error) {
	err = nil
	var exist = true

	exist, err = PathExists(path)
	if err != nil {
		return
	}
	if exist {
		return
	} else {
		err = os.Mkdir(path, 0777)
		if err != nil {
			return
		}
		exist = true
	}

	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}