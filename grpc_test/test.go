package main

import "fmt"

func main() {
	res := isEscapePossible([][]int{}, []int{0, 0}, []int{999999, 999999})
	fmt.Println(res)
	fmt.Println("Hello, World!")
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	type pair struct {
		x int
		y int
	}
	directs := [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	maxConnRegion := len(blocked) * (len(blocked) - 1) / 2
	fmt.Println("maxConnRegion: ", maxConnRegion)
	var check func(blocked [][]int, source []int, target []int) bool
	check = func(blocked [][]int, source []int, target []int) bool {
		visit := map[pair]struct{}{}
		for _, block := range blocked {
			visit[pair{block[0], block[1]}] = struct{}{}
		}
		q := []pair{{source[0], source[1]}}
		cnt := 0
		for len(q) > 0 {
			l := len(q)

			fmt.Println("len(q) ", l)
			for ; l > 0; l-- {
				p := q[0]
				q = q[1:]
				for _, d := range directs {
					x := p.x + d[0]
					y := p.y + d[1]

					fmt.Println("x: ", x, "y: ", y)

					fmt.Println("10^6: ", 10^6)
					if x < 1000000 && x >= 0 && y < 1000000 && y >= 0 {
						if target[0] == x && target[1] == y {
							return true
						}
						if _, ok := visit[pair{x, y}]; !ok {
							cnt++
							visit[pair{x, y}] = struct{}{}
							q = append(q, pair{x, y})
						}
					}
				}

				fmt.Println("cnt: ", cnt)
				if cnt > maxConnRegion {
					return true
				}
			}
		}

		fmt.Println("-------------------------")
		return false
	}

	return check(blocked, source, target) && check(blocked, target, source)
}
