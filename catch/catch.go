package catch

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Catch (callback func(err error)) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			callback(err)
		} else {
			panic(r)
		}
	}
}