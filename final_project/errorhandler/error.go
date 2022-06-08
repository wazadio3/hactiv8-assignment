package errorhandler

func Check(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}
