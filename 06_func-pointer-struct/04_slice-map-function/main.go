package main

import "fmt"

func main() {
	mp := map[string]int{"foo": 1}
	updateMap(mp)
	fmt.Println(mp)

	vals := []int{1, 2, 3, 4, 5}
	// updateSlice(vals)
	fmt.Println(vals)

	vals = addElement(vals, 6)
	fmt.Println(vals)
}

func updateMap(mp map[string]int) {
	mp["foo"] = 42
	mp["bar"] = 56
}

func updateSlice(vals []int) {
	// แบบนี้จะเปลี่ยนค่า slice ที่ส่งมา
	for i := range vals {
		vals[i] += 1
	}

	// vals = append(vals, 6) // อันนี้ค่าที่ถูกส่งมาจะไม่ถูกเปบี่ยนเพราะเป็นการ new slice
}

func addElement(vals []int, elem int) []int {
	return append(vals, elem)
}
