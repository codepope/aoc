import re
f=open("../input.txt")
lines=f.readlines()
total=0
for n in lines:
    r=re.search("([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)",n)
    (np10,np11,np20,np21)=(int(r.group(1)),int(r.group(2)),int(r.group(3)),int(r.group(4)))
    if (np10 >= np20 and np11 <= np21) or (np20 >= np10 and np21 <= np11):
        total += 1    
print(total)
