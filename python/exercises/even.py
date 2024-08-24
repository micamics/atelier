# Write a program that prints all the even number from 1 - 10.
# Do not use the skip argument in range().
# Use a loop and iterate over the numbers.

counter = 0
for num in range(1, 10):
    if num % 2 == 0:
        print(num)
        counter += 1

print(f"We have {counter} even numbers")
