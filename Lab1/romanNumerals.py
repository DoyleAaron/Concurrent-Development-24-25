
# This function finds the double values and removes them from the roman numeral list
def find_doubles(romanNumList, totalAmount):
    i = len(romanNumList) - 1
    while i > 0:
        if romanNumList[i] == "v" and romanNumList[i - 1] == "i":
            totalAmount = totalAmount + 4
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            i = i - 2

        elif romanNumList[i] == "x" and romanNumList[i - 1] == "i":
            totalAmount = totalAmount + 9
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            i = i - 2
        
        elif romanNumList[i] == "l" and romanNumList[i - 1] == "x":
            totalAmount = totalAmount + 40
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            i = i - 2

        elif romanNumList[i] == "c" and romanNumList[i - 1] == "x":
            totalAmount = totalAmount + 90
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            i = i - 2
        
        elif romanNumList[i] == "d" and romanNumList[i - 1] == "c":
            totalAmount = totalAmount + 400
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            i = i - 2

        elif romanNumList[i] == "m" and romanNumList[i - 1] == "c":
            totalAmount = totalAmount + 900
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            i = i - 2
        
        elif romanNumList[i] == "i" and romanNumList[i - 1] == "i" and romanNumList[i - 2] == "i":
            totalAmount = totalAmount + 3
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            romanNumList.pop(i - 2)
            i = i - 3
        
        elif romanNumList[i] == "i" and romanNumList[i - 1] == "i":
            totalAmount = totalAmount + 2
            romanNumList.pop(i)
            romanNumList.pop(i - 1)
            i = i - 2

        else:
            i = i - 1
        
    return romanNumList, totalAmount

# This function is to calculate the rest of the single values
def calculate_rest(romanNumList, totalAmount):
    i = 0
    while i < len(romanNumList):
        if romanNumList[i] == "i":
            totalAmount = totalAmount + 1
            i = i + 1
        
        elif romanNumList[i] == "v":
            totalAmount = totalAmount + 5
            i = i + 1
        
        elif romanNumList[i] == "x":
            totalAmount = totalAmount + 10
            i = i + 1
        
        elif romanNumList[i] == "l":
            totalAmount = totalAmount + 50
            i = i + 1

        elif romanNumList[i] == "c":
            totalAmount = totalAmount + 100
            i = i + 1

        elif romanNumList[i] == "d":
            totalAmount = totalAmount + 500
            i = i + 1
        
        elif romanNumList[i] == "m":
            totalAmount = totalAmount + 1000
            i = i + 1

        else:
            i = i + 1

    return totalAmount

romanNum = input("Enter your roman numeral: ").lower()
valid_letters = {'i', 'v', 'x', 'l', 'c', 'd', 'm'}

# Check for invalid characters and alert the user
if not set(romanNum).issubset(valid_letters):
    print("Invalid input, please enter a valid roman numeral containing only: i, v, x, l, c, d, m.")
else:
    romanNumList = list(romanNum)

    if len(romanNum) > 15:
        print("This input is too long.")
    else:
        totalAmount = 0
        romanNumList, totalAmount = find_doubles(romanNumList, totalAmount)
        finalAmount = calculate_rest(romanNumList, totalAmount)
        if finalAmount >= 4000:
            print("This number is too big, number has to be between 1 - 3999")
        else:
            print(finalAmount)