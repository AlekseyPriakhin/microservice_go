package utils

func HandleError(err error, cb func()) {
	if err != nil {
		cb()
	}
}

func Map[TIn, TOut any](list []TIn, f func(TIn) TOut) []TOut {
	us := make([]TOut, len(list))
	for i := range list {
		us[i] = f(list[i])
	}
	return us
}
