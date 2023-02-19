package util

func Min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func NewString(s string) *string {
	return &s
}
