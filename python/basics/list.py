names = ["John", "Bob", "Joseph", "Sam", "Mary"]
print(names)
print(names[0])  # first element on the list
print(names[-1])  # last element on the list

print()
names[1] = "Bobbie"  # replace element on the list indicated by the index
print(names)
# prints elements from 0 to 3 index, excluding the last index
print(names[0:3])
print(names[1:])  # prints from index 1 until the last element
print(names)

# list methods
print(len(names))
print(names.reverse())
print(names.count("Bob"))
