package utils

func Wrap[T any](x T) (r *T) {
	r = &x
	return
}
