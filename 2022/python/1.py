import sys
from collections import defaultdict

def main():
    inputLines = sys.stdin.readlines()
    for i in range(len(inputLines)):
        inputLines[i] = inputLines[i].rstrip()

    #print(inputLines)
    numElvs = 0
    nextElf = True 
    elfCal = defaultdict(int)

    for cal in inputLines:
        if nextElf:
            numElvs+=1
            nextElf = False
        if cal == "":
            nextElf = True
            continue

        elfCal[numElvs] += int(cal)

    maxElf = -1
    maxCal = 0

    for key, value in elfCal.items():
        if value > maxCal:
            maxCal = value
            maxElf = key

    print(maxElf, maxCal)



if __name__ == "__main__":
    main()
