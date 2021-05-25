package main

func PlayGame(player int, computer int, choices []*Choice) string {
	if player == computer {
		return "tie"
	} else if contains(choices[player-1].Beats, computer) {
		return "win"
	}
	return "lose"
}

func contains(data []int, number int) bool {
	for _, value := range data {
		if value == number {
			return true
		}
	}
	return false
}
