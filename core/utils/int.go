package utils

//Intersection 重叠
func Intersection(s1, s2 []int) []int {
	m := make(map[int]int)
	for k := range s1 {
		m[s1[k]] += 1
	}
	var a []int
	for k := range s2 {
		for key, value := range m {
			if key == s2[k] && value > 0 {
				m[k] -= 1
				a = append(a, key)
			}
		}
	}
	return a
}
