package main

func InsertSortDesc(vals []int) ([]int) {
	vLen := len(vals)
	for i:=1; i < vLen; i++ {
		key := vals[i]
		j := i - 1
		for 0 <= j && vals[j] < key {
			vals[j+1] = vals[j]
			j--
		}
		vals[j+1] = key
	}
	return vals
}
