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

tier_one = Elf()
tier_two = Elf()
tier_three = Elf()

for e in list_of_elves:
	if e.get_sum() >= tier_one.get_sum():
		tier_three = tier_two
		tier_two = tier_one
		tier_one = e
	elif e.get_sum() >= tier_two.get_sum():
		tier_three = tier_two
		tier_two = e
	elif e.get_sum() >= tier_three.get_sum():
		tier_three = e
	else:
		pass


# 	this_sum = e.get_sum()
# 	if this_sum > max_cal:
# 		max_cal = this_sum

# print(max_cal)

print(f"First place: {tier_one}")
print(f"Second place: {tier_two}")
print(f"Third place: {tier_three}")
print(f"Total: {tier_one.get_sum() + tier_two.get_sum() + tier_three.get_sum()}")