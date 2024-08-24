# Create a fizz_buzz function that accepts an input and:
# - if input is divisible by 3, return Fizz
# - if input is divisible by 5, return Buzz
# - if input is divisible by both 3 and 5, return FizzBuzz
# - else, echo input.

def fizz_buzz(input):
    if input % 3 == 0 and input % 5 == 0:
        return "FizzBuzz"
    if input % 5 == 0:
        return "Buzz"
    if input % 3 == 0:
        return "Fizz"
    return input


print(f"3: {fizz_buzz(3)}")
print(f"5: {fizz_buzz(5)}")
print(f"15: {fizz_buzz(15)}")
print(f"7: {fizz_buzz(7)}")
