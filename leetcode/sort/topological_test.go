package sort

/**
应用场景: 要将图论里的顶点按照相连的性质进行排序
前提: 1.必须是有向图 2.图中没有环

时间复杂度O(n)
统计顶点的入度 O(n)
每个顶点被遍历一次 O(n)

*/
var rs []int

func topoSort(indegree []int, adj [][]int) {
	q := make([]int, 0)

	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		//出队
		i := q[0]
		q = q[1:]
		rs = append(rs, i)

		for j := 0; j < len(adj[i]); j++ {
			if adj[i][j] == 1 {
				indegree[j]--
				if indegree[j] == 0 {
					q = append(q, j)
				}
			}
		}
	}
}
