// Identical expression on both side of '&&'
// Here len(s1) < cap(s1) && len(s1) < cap(s1) indicates a mistake
// CLSC should generate a warning

package testdata

func compare(s1, s2 []int) bool {
	if len(s1)+cap(s1) < len(s1)+cap(s1) {
		return true
	}
	return false
}

//<<<<<203, 236>>>>>
