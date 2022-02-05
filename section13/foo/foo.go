package foo

const (
	Max = 100
	// 1文字目が小文字の場合はprivate
	min = 1
)

func ReturnMin() int {
	return min
}
