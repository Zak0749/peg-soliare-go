package pretify

func icon(t bool) string {
	if t {
		return "●"
	} else {
		return "○"
	}
}

func Board(board [45]bool) string {
	str := "      "
	str += icon(board[0]) + " " + icon(board[1]) + " " + icon(board[2])
	str += "\n"
	str += "      "
	str += icon(board[3]) + " " + icon(board[4]) + " " + icon(board[5])
	str += "\n"
	str += "      "
	str += icon(board[6]) + " " + icon(board[7]) + " " + icon(board[8])
	str += "\n"
	str += icon(board[9]) + " " + icon(board[10]) + " " + icon(board[11]) + " " + icon(board[12]) + " " + icon(board[13]) + " " + icon(board[14]) + " " + icon(board[15]) + " " + icon(board[16]) + " " + icon(board[17])
	str += "\n"
	str += icon(board[18]) + " " + icon(board[19]) + " " + icon(board[20]) + " " + icon(board[21]) + " " + icon(board[22]) + " " + icon(board[23]) + " " + icon(board[24]) + " " + icon(board[25]) + " " + icon(board[26])
	str += "\n"
	str += icon(board[27]) + " " + icon(board[28]) + " " + icon(board[29]) + " " + icon(board[30]) + " " + icon(board[31]) + " " + icon(board[32]) + " " + icon(board[33]) + " " + icon(board[34]) + " " + icon(board[35])
	str += "\n"
	str += "      "
	str += icon(board[36]) + " " + icon(board[37]) + " " + icon(board[38])
	str += "\n"
	str += "      "
	str += icon(board[39]) + " " + icon(board[40]) + " " + icon(board[41])
	str += "\n"
	str += "      "
	str += icon(board[42]) + " " + icon(board[43]) + " " + icon(board[44])

	return str
}
