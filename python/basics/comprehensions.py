# List comprehension is the preferred way to perform mapping of list to a new list.
# It's cleaner and more performant than map() or filtered().

# We can use comprehension for lists, sets, arrays, dictionaries data types

# before we had --> prices = list(map(lambda item: item[1], items))
# list comprehension syntax -> [expression for parameter in list condition] --> condition is optional

items = [
    ("Product1", 10),
    ("Product2", 7),
    ("Product3", 5),
]

prices = [item[1] for item in items]
print(f"Prices {prices}")

# before --> filtered = list(filter(lambda item: item[1] >= 10, items))
filtered = [item for item in items if item[1] >= 10]
print(f"Filtered: {filtered}")

values = {}
for x in range(5):
    values[x] = x*2
print(f"Values: {values}")


print()
print("Using comprehension below: ")

# lines 22 - 25 is the same as below using comprehension
values = {x: x*2 for x in range(5)}
print(f"Values: {values}")
