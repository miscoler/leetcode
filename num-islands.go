package main

func markLand(grid [][]byte, x, y int) {
	if x < 0 ||  y < 0 || x >= len(grid) || y >= len((grid)[0])  {
		return
	}
	if (grid)[x][y] == '1' {
		(grid)[x][y] = '0'
        markLand(grid, x-1, y)
        markLand(grid, x, y+1)
        markLand(grid, x+1, y)
        markLand(grid, x, y-1)
	}
}

func numIslands(grid [][]byte) (res int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res += 1
				markLand(grid, i, j)
			}
		}
	}
	return
}
