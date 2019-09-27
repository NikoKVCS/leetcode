// +build ignore

package main

import (
	"fmt"
)

func seq(i, j, M int) int {
	return i*M + j
}

func main() {
	N := 0
	M := 0
	fmt.Scan(&N, &M)

	matrix := make([][]int, 0)

	for i := 0; i < N; i++ {
		temp := make([]int, 0)
		for j := 0; j < M; j++ {
			ai := 0
			fmt.Scan(&ai)
			temp = append(temp, ai)
		}
		matrix = append(matrix, temp)
	}

	graph := make([][]int, N*M)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			relations := make([]int, 0)
			if j+1 < M && matrix[i][j+1] == 1 {
				relations = append(relations, seq(i, j+1, M))
			}
			if j-1 >= 0 && i+1 < N && matrix[i+1][j-1] == 1 {
				relations = append(relations, seq(i+1, j-1, M))
			}
			if i+1 < N && matrix[i+1][j] == 1 {
				relations = append(relations, seq(i+1, j, M))
			}
			if i+1 < N && j+1 < M && matrix[i+1][j+1] == 1 {
				relations = append(relations, seq(i+1, j+1, M))
			}
			graph[seq(i, j, M)] = relations
		}
	}

	vis := make([]int, N*M)
	count := 0
	for i := 0; i < len(graph); i++ {
		if matrix[i/M][i%M] == 0 || vis[i] == 1 {
			continue
		}
		bfs(i, graph, vis)
		count++
	}

	fmt.Printf("%d\n", count)

}
func bfs(i int, graph [][]int, vis []int) {
	vis[i] = 1
	if len(graph[i]) == 0 {
		return
	}
	top := 0
	queue := make([]int, 0)
	queue = append(queue, graph[i]...)
	for top < len(queue) {
		id := queue[top]
		if vis[id] == 1 {
			top++
			continue
		}
		queue = append(queue, graph[id]...)
		vis[id] = 1
		top++
	}
}
