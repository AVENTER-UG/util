package util

func StringToPointer(value string) *string {
	return func() *string { x := value; return &x }()
}

func Uint32ToPointer(value uint32) *uint32 {
	return func() *uint32 { x := value; return &x }()
}

func IntToPointer(value int) *int {
	return func() *int { x := value; return &x }()
}
