package gotils

func HandleError(err error) (result error) {
	if err != nil {
		result = err
		defer func() {
			panic(err)
		}()
	}
	return
}
