# Create a program that checks the most commonly used character in an inputted text/phrase.

sentence = "This is a common interview question."

char_frequency = {}
for char in sentence:
    val = char_frequency.get(char, 0)
    if val == 0:
        char_frequency[char] = 1

    char_frequency[char] = val + 1


asList = list(char_frequency.items())
asList.sort(key=lambda char: char[1], reverse=True)

print(f"Most Commonly Used: {asList[0]}")
