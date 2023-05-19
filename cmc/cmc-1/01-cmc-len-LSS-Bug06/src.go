// len(x) will return non negative value.
// so the expression len(x) < 0 is always false.
// CMC should generate a warning.

package testdata

import "fmt"

func test() {
	var a [10]int
	if len(a) < 0 {
		fmt.Println("Length is negative")
	}
}

//<<<<<191, 201>>>>>
