# lambda func syntax: lambda parameter:expression

items = [
    ("Product1", 10),
    ("Product2", 7),
    ("Product3", 5),
]

# sorting
items.sort(key=lambda item: item[1])
print(items)


# getting all prices and pass it to a new list
# used map()
prices = list(map(lambda item: item[1], items))
print(f"Prices: {prices}")

# filter based on a condition
filtered = list(filter(lambda item: item[1] >= 10, items))
print(f"Filtered: {filtered}")
