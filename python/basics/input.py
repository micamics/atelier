print("Tell me anything...")
anything = input()
print("Hmm...", anything, "... Really?")

print()
anything = input("Again, tell me anything...")
print("Hmm...", anything, "...Really?")

## return value of input() is a string
## type-casting input to float
print()
anything = float(input("Enter a number: "))
something = anything ** 2.0
print(anything, "to the power of 2 is", something)

## type-casting input to int
print()
anything = int(input("Enter a number: "))
something = anything ** 2
print(anything, "to the power of 2 is", something)
