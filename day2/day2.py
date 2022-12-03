#!/usr/bin/env python3

# Legend:
# A = Rock
# B = Paper
# C = Scissors
# X = Rock, 1 point
# Y = Paper, 2 points
# Z = Scissors, 3 points

# Round outcome: loss is 0, 3 if draw, 6 if win

# ABC: opponent, XYZ: me

def get_results(me, opponent):
	"""
	Returns "Win", "Loss", or "Draw"
	"""
	if me == "X":
		if opponent == "A":
			return "Draw"
		if opponent == "B":
			return "Loss"
		return "Win"

	if me == "Y":
		if opponent == "A":
			return "Win"
		if opponent == "B":
			return "Draw"
		return "Loss"

	if me == "Z":
		if opponent == "A":
			return "Loss"
		if opponent == "B":
			return "Win"
		return "Draw"

	return None

def get_score(scoreline):
	"""
	Gets a score based on the provided line,
	returns two values: Win/Loss/Draw, Total Score
	"""
	ScoreLegend = {"X": 1, "Y": 2, "Z": 3}
	opponent, me = scoreline.split()
	result = get_results(me, opponent)
	if result is None:
		return None
	if result == "Win":
		score = 6
	elif result == "Loss":
		score = 0
	elif result == "Draw":
		score = 3

	return result, score + ScoreLegend[me]

with open("input.dat", "r") as f:
	scores = f.read().splitlines()

totalScore = 0
for score in scores:
	result, gameScore = get_score(score)
	totalScore += gameScore
	# print(f"Game: {score}, Result: {get_score(score)}")

print(f"Total score: {totalScore}")