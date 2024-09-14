romanNum = input("Enter your roman numeral: ")
romanNum = romanNum.lower()
romanNumList = list(romanNum)


totalAmount = 0
i = len(romanNumList) - 1


while i > 0:
    if romanNumList[i] == "v" and romanNumList[i - 1] == "i":
        totalAmount = totalAmount + 4
        romanNumList.pop[i]
        romanNumList.pop[i - 1]
        i = i - 1

    elif romanNumList[i] == "x" and romanNumList[i - 1] == "i":
        totalAmount = totalAmount + 9
        romanNumList.pop[i]
        romanNumList.pop[i - 1]
        i = i - 1
    
    elif romanNumList[i] == "l" and romanNumList[i - 1] == "x":
        totalAmount = totalAmount + 9
        romanNumList.pop[i]
        romanNumList.pop[i - 1]
        i = i - 1