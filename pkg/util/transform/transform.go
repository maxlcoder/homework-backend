package transform

func ConvertSlice[T any, U any](in []T, covert func(T) U) []U {
	out := make([]U, len(in))
	for i := range in {
		out[i] = covert(in[i])
	}
	return out
}
