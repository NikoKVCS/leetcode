n = int(input())
our_t = input().split()
ours = []
for i in range(0,n):
    ours.append(int(our_t[i]))

your_t = input().split()
yours = []
for i in range(0,n):
    yours.append(int(your_t[i]))

ours.sort(reverse=True)
yours.sort(reverse=True)
tmp=0
m=0
lose=0
for i in range(len(ours)):
    for j in range(m,len(yours)):
        if ours[i]>yours[j]:
            tmp+=1
            j+=1
            m=j
            break
        elif ours[i]<yours[j]:
            lose +=1
print(tmp-lose)