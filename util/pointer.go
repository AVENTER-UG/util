package util

func StringToPointer(value string) *string {
	return func() *string { x := value; return &x }()
}

func Uint32ToPointer(value uint32) *uint32 {
	return func() *uint32 { x := value; return &x }()
}

func Uint64ToPointer(value uint64) *uint64 {
	return func() *uint64 { x := value; return &x }()
}

func IntToPointer(value int) *int {
	return func() *int { x := value; return &x }()
}

func BoolToPointer(value bool) *bool {
	return func() *bool { x := value; return &x }()
}
