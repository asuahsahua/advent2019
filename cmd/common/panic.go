package common

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIf(v bool, message string) {
	if v == true {
		panic(message)
	}
}