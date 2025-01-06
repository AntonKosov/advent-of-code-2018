package maps

func SingleKey[K comparable, V any](m map[K]V) K {
	if len(m) != 1 {
		panic("map must have exactly one value")
	}

	for k := range m {
		return k
	}

	panic("not found")
}
