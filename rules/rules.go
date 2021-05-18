// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func test(m dsl.Matcher) {
	m.Match(`a := 42`).Report("A found")
	//m.Match(`b := 42`).Report("B found")
}
