package Gate

func buf(v bool) bool {
	return v
}

func inv(v bool) bool {
	return !v
}

func or(v, s bool) bool {
	return v || s
}

func and(v, s bool) bool {
	return v && s
}

func xor(v, s bool) bool {
	return (v || s) && !(v && s)
}

func nor(v, s bool) bool {
	return !(v || s)
}

func nand(v, s bool) bool {
	return !(v && s)
}

func xnor(v, s bool) bool {
	return !((v || s) && !(v && s))
}
