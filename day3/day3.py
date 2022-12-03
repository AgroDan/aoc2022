#!/usr/bin/env python3

import string
import pprint

# This is what I started to do before I realized that _we can fix this. we have the technology._
# priority = {
# 	"a" : 1,
# 	"b" : 2,
# 	"c" : 3,
# 	"d" : 4,
# 	"e" : 5,
# 	"f"
# }

# Define a function to basically do all the work for us
def score_rucksack(r):
	"""
		takes a rucksack and splits it, finds the common item
		and returns the "score", which is to say the numeric
		value of the priority assigned to the common item.
	"""
	# First, let's define a priority matrix
	# And for the record, I can't believe this works
	priority = {l: num+1 for num, l in enumerate(string.ascii_letters)}

	left = r[:int(len(r)/2)]
	right = r[int(len(r)/2):]

	for item in left:
		if item in right:
			return priority[item]

	# Crap I think that's basically all I need to do.

# Now let's ingest the input
with open("input", "r") as f:
	rucksacks = f.read().splitlines()

# Just to test, let's start with 1
# rucksack = rucksacks[5]
# print(len(rucksack))
# print(f"One half: {rucksack[:int(len(rucksack)/2)]}")
# print(f"One half: {rucksack[int(len(rucksack)/2):]}")

# priority = {l: num+1 for num, l in enumerate(string.ascii_letters)}
# pprint.pprint(priority)
# print(f"Result of rucksack 0: {score_rucksack(rucksacks[0])}")

total_score = 0
for rucksack in rucksacks:
	total_score += score_rucksack(rucksack)

print(f"Final score: {total_score}")