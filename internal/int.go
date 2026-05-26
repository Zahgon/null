package internal

type Integer interface {
	int64 | int32 | int16 | byte
}

func UnmarshalIntJSON[T Integer, U int64 | uint64](data []byte, value *T, valid *bool, bits int, parse func(string, int, int) (U, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalIntText[T Integer, U int64 | uint64](text []byte, value *T, valid *bool, bits int, parse func(string, int, int) (U, error)) error {
	_ = "STUB: not implemented"
	return nil
}
