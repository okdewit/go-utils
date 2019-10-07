package catch

func Check(err error) {
	if err != nil {
		panic(err)
	}
}