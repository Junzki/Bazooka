package assets

import (
	"os"
	"path"
	"strings"
)

const (
	userHomeDirFlag = "~"
)

func ExpandUserDir(p string) string {
	if ! strings.HasPrefix(p, userHomeDirFlag) {
		return p
	}

	userDir, err := os.UserHomeDir()
	if nil != err {
		// this error means current platform is not recognized by Golang.
		// See documents for os.UserHomeDir for detail.
		return p
	}

	if userHomeDirFlag == p {
		return userDir
	}

	p = p[1:] // Remove prefix.
	p = path.Join(userDir, p)

	return p
}
