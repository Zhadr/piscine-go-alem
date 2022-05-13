package piscine

func UltimateDivMod(a *int, b *int) {
	x := *a
	*a = x / (*b)
	*b = x % *b
}
