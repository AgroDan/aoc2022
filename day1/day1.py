#!/usr/bin/env python

with open('input.dat', 'r') as f:
	calories = f.read().splitlines()

# print(calories[0])

# for each line, create a new complex object

class Elf:
	def __init__(self):
		self.food_item = []

	def __repr__(self):
		return f"Elf with {len(self.food_item)} items totalling {self.get_sum()}"

	def push(self, cals):
		cals = int(cals)
		self.food_item.append(cals)

	def get_sum(self):
		return sum(self.food_item)

# Now create a working list
list_of_elves = []

e = Elf()
for cal in calories:
	if cal == '':
		list_of_elves.append(e)
		e = Elf()
	else:
		e.push(cal)

# Now let's just find the highest calorie count
max_cal = 0
for e in list_of_elves:
	this_sum = e.get_sum()
	if this_sum > max_cal:
		max_cal = this_sum

print(max_cal)