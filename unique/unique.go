package unique

func reverseArray(arr [45]bool) [45]bool {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func Boards(boards [][45]bool) [][45]bool {
	var unique [][45]bool
	m := map[[45]bool]bool{}

	for _, v := range boards {
		if !m[v] && !m[reverseArray(v)] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
