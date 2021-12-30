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
    x1 = dir[0]
    y1 = dir[1]
    x2 = dir[2]
    y2 = dir[3]
    xDiff = int(x2) - int(x1)
    yDiff = int(y2) - int(y1)        
    movex = abs(xDiff)
    movey = abs(yDiff)

    print(xDiff)
    print(yDiff)
    print(movex)
    print(movey)
    


