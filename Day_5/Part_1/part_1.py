lines = []
x1 = ""
y1 = ""
x2 = ""
y2 = ""
m = 10
n = 10
ventMap = [[0 for x in range(n)] for x in range(m)]

with open('input.txt') as dir:
    lines = dir.readlines()

for line in lines:
    line = line.replace(" -> ", ",")
    dir = line.split(",")
    x1 = int(dir[0])
    y1 = int(dir[1])
    x2 = int(dir[2])
    y2 = int(dir[3])
    xDiff = x2 - x1
    yDiff = y2 - y1        
    movex = abs(xDiff)
    movey = abs(yDiff)

    if xDiff < 0 & yDiff < 0:
        if movex >= movey:
            while movex > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 -= 1
                y1 -= 1
        else:
            while movey > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 -= 1
                y1 -= 1
    elif xDiff >= 0 & yDiff >= 0:    
        if movex >= movey:
            while movex > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 += 1
                y1 += 1
        else:
            while movey > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 += 1
                y1 += 1
    elif xDiff >= 0 & yDiff < 0:
        if movex >= movey:
            while movex > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 += 1
                y1 -= 1
        else:
            while movey > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 += 1
                y1 -= 1
    else:
        if movex >= movey:
            while movex > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 -= 1
                y1 += 1
        else:
            while movey > 0:
                ventMap[x1][y1] += 1
                movex -= 1
                x1 -= 1
                y1 += 1
total = 0

for item in ventMap:
    if item > 1:
        total += 1

print(total)
