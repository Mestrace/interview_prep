package p8_2

// Robot in a Grid
// 如果可以走, 返回一个列表
// 如果不能走, 返回空

func RobotInAGridDFS(grid [][]uint8) [][]int {
	result := robotInAGridDFS(grid, 0, 0)
	reverse(result)
	return result
}

func robotInAGridDFS(grid [][]uint8, i, j int) [][]int {
	if grid[j][i] > 0 {
		return nil
	}
	if i == len(grid[0])-1 && j == len(grid)-1 {
		result := make([][]int, 0, i+j)
		result = append(result, []int{i, j})
		return result
	} else if i == len(grid[0])-1 {
		if result := robotInAGridDFS(grid, i, j+1); result != nil {
			return append(result, []int{i, j})
		}
	} else if j == len(grid)-1 {
		if result := robotInAGridDFS(grid, i+1, j); result != nil {
			return append(result, []int{i, j})
		}
	} else {
		if result := robotInAGridDFS(grid, i+1, j); result != nil {
			return append(result, []int{i, j})
		}
		if result := robotInAGridDFS(grid, i, j+1); result != nil {
			return append(result, []int{i, j})
		}
	}

	return nil
}

func reverse(path [][]int) {
	for i, tmp := range path {
		if i >= len(path)/2 {
			break
		}
		end := len(path) - i - 1
		println(i, end)
		path[i] = path[end]
		path[end] = tmp
	}
}

/*
LC 62
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
*/

func UniquePathDP(grid [][]uint8) int {
	// handle empty
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for j := range dp {
		dp[j] = make([]int, len(grid[0]))
	}

	// if the first one is zero
	if grid[0][0] == 0 {
		dp[0][0] = 1
	} else {
		return 0
	}

	for i := 1; i < len(dp[0]); i++ {
		if grid[0][i] == 0 {
			dp[0][i] = dp[0][i-1]
		}
	}

	for j := 1; j < len(dp); j++ {
		if grid[j][0] == 0 {
			dp[j][0] = dp[j-1][0]
		}
	}

	// make sure grid has more than 1 row/column, otherwise we are done
	if len(dp) == 1 || len(dp[0]) == 1 {
		return dp[len(grid)-1][len(grid[0])-1]
	}

	for j := 1; j < len(dp); j++ {
		for i := 1; i < len(dp[j]); i++ {
			if grid[j][i] == 0 {
				dp[j][i] = dp[j-1][i] + dp[j][i-1]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
