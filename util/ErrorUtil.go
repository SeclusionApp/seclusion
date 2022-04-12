package util

func HandleError(err error, msg string) {
	if err != nil {
		panic(msg + " | " + err.Error())
	}
}
