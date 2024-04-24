package util

func String2Pointer(value string) *string {
	return func() *string { x := value; return &x }()

}
