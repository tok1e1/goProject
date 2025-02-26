package main

import (
	"fmt"
	"math"
)

type Exchange struct {
	from, to, cost int
}

// Поиск цикла с отрицательным весом (алгоритм Беллмана-Форда)
func findNegativeCycle(n int, exchanges []Exchange) ([]int, bool) {
	const INF = math.MaxInt
	dist := make([]int, n+1)
	prev := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dist[i] = INF
		prev[i] = -1
	}

	dist[1] = 0

	// Применение алгоритма Беллмана-Форда
	for i := 1; i < n; i++ {
		for _, exchange := range exchanges {
			if dist[exchange.from] != INF && dist[exchange.from]+exchange.cost < dist[exchange.to] {
				dist[exchange.to] = dist[exchange.from] + exchange.cost
				prev[exchange.to] = exchange.from
			}
		}
	}

	for _, exchange := range exchanges {
		if dist[exchange.from] != INF && dist[exchange.from]+exchange.cost < dist[exchange.to] {
			return restoreCycle(prev, exchange.to), true
		}
	}

	return nil, false
}

// Восстановление пути отрицательного цикла
func restoreCycle(prev []int, start int) []int {
	visited := make([]bool, len(prev))
	v := start

	for !visited[v] {
		visited[v] = true
		v = prev[v]
	}

	cycle := []int{}
	startCycle := v
	for {
		cycle = append(cycle, v)
		v = prev[v]
		if v == startCycle {
			break
		}
	}
	for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
		cycle[i], cycle[j] = cycle[j], cycle[i]
	}

	return cycle
}

// Вывод цикла
func printCycle(cycle []int) {
	fmt.Println("YES")
	for _, v := range cycle {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	exchanges := make([]Exchange, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &exchanges[i].from, &exchanges[i].to, &exchanges[i].cost)
	}

	if cycle, found := findNegativeCycle(n, exchanges); found {
		printCycle(cycle)
	} else {
		fmt.Println("NO")
	}
}
