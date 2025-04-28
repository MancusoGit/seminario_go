package ejercicios

func Evaluar(num int) int {
	switch {
	case num < -18:
		return num * -1
	case num >= -18 && num <= -1:
		return num % 4
	case num >= 0 && num < 20:
		return num * num
	case num >= 20:
		return -num
	default:
		return 0
	}
}
