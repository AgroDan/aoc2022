#!/usr/bin/env python3

# Legend:
# A = Rock
# B = Paper
# C = Scissors
# X = need to lose
# Y = Need to draw
# Z = need to win

# ALSO...
# 	rock = 1
# 	paper = 2
# 	scissors = 3

# Round outcome: loss is 0, 3 if draw, 6 if win

# ABC: opponent, XYZ: me

def get_results(me, opponent):
	"""
	Returns total score of the end result
	"""
	if me == "X":
		# need to lose
		if opponent == "A":
			# we throw scissors, which is 3
			return 3
		if opponent == "B":
			# we throw rock, which is 1
			return 1

		# Otherwise opponent throws scissors, we throw paper
		return 2

	if me == "Y":
		# need to draw
		if opponent == "A":
			# we throw rock
			return 4
		if opponent == "B":
			# we throw paper
			return 5

		# Otherwise opponent throws scissors, we throw scissors
		return 6

	if me == "Z":
		# need to win
		if opponent == "A":
			# we throw paper
			return 8
		if opponent == "B":
			# we throw scissors
			return 9

		# Otherwise opponent throws scissors, we throw rock
		return 7

	return None

with open("input.dat", "r") as f:
	scores = f.read().splitlines()

totalScore = 0
for score in scores:
	opponent, me = score.split()
	totalScore += get_results(me, opponent)

print(f"Total score: {totalScore}")