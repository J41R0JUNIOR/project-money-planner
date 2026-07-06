package middleware

func Noop() func(next func()) func() {
	return func(next func()) func() {
		return func() {
			next()
		}
	}
}
