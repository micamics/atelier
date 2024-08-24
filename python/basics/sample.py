weight = float(input("Weight: "))
unit = input("(K)g or (L)bs: ")

if unit.upper() == "L":
    weight *= 0.45
    print("Weight in Kg: ", weight)
else:
    weight /=  0.45
    print("Weight in Lbs: ", weight)

print("Done")
