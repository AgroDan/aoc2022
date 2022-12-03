#!/usr/bin/env python3

import string

# Define a function to determine the common badge in 3 rucksacks and score appropriately
def find_common_badge(ruck1, ruck2, ruck3):
	"""
	Returns the score of the common item in all 3 rucksacks
	"""
	priority = {l: num+1 for num, l in enumerate(string.ascii_letters)}
	for l in ruck1:
		if l in ruck2 and l in ruck3:
			return priority[l]


# Now let's ingest the input
with open("input", "r") as f:
	rucksacks = f.read().splitlines()

total_score = 0

for i in range(0, len(rucksacks), 3):
	total_score += find_common_badge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])

print(f"Final score: {total_score}")