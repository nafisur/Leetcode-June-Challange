package leetcode

/*
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'.
Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'.
Two cells are connected if they are adjacent cells connected horizontally or vertically.
*/
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if (i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'B' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 || board[i][j] == 'B' || board[i][j] == 'X' {
		return
	}
	board[i][j] = 'B'
	dfs(board, i-1, j)
	dfs(board, i, j-1)
	dfs(board, i+1, j)
	dfs(board, i, j+1)
}
