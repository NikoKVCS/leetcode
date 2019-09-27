n = int(input())
our_t = input().split()
ours = []
for i in range(0,n):
    ours.append(int(our_t[i]))

your_t = input().split()
yours = []
for i in range(0,n):
    yours.append(int(your_t[i]))
a = yours
b = ours
a.sort(reverse=True)
b.sort(reverse=True)
c = []
for i in range(0,n):
    if a[i]>b[0]:
        c.append(b.pop())
    else:
        c.append(b.pop(0))
win,lose,equal = 0,0,0
for i in range(0,n):
    if a[i] < c[i]:
        win += 1
    elif a[i] == c[i]:
        equal+=1
    else:
        lose += 1

marks = win - lose
print(marks)