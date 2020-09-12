package helpers

import "github.com/sirupsen/logrus"

func ThrowIfError(err error) {
	if err != nil {
		logrus.Fatalln(err)
	}
}

func LogIfError(err error) {
	if err != nil {
		logrus.Errorln(err)
	}
}
