package leetcode

// https://leetcode.com/problems/flood-fill/

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		dfs(image, sr, sc, color, newColor)
	}

	return image
}

func dfs(image [][]int, sr int, sc int, color int, newColor int) {
	if image[sr][sc] == color {
		image[sr][sc] = newColor

		if sr >= 1 {
			dfs(image, sr-1, sc, color, newColor)
		}

		if sc >= 1 {
			dfs(image, sr, sc-1, color, newColor)
		}

		if sr+1 < len(image) {
			dfs(image, sr+1, sc, color, newColor)
		}

		if sc+1 < len(image[0]) {
			dfs(image, sr, sc+1, color, newColor)
		}
	}
}
