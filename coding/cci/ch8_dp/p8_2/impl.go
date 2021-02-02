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
