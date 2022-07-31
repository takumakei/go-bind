package bind_test

import (
	"fmt"

	"github.com/takumakei/go-bind"
)

func Example() {
	div := func(a, b int) (d, r int) {
		d = a / b
		r = a % b
		return
	}

	// This is equivalent to
	//
	//  fn := func() (int, int) { return div(10, 3) }
	//
	fn := bind.P2R2(div, 10, 3)

	d, r := fn()

	fmt.Printf("%d, %d", d, r)
	// output: 3, 1
}
