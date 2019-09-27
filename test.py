a = [405, 395, 374, 410, 417, 426, 383, 398, 390, 402]
mu = sum(a) / len(a)
sig = 0
for i in range(0,len(a)):
    sig += (a[i] - mu)*(a[i] - mu)
sig /= len(a) - 1
print(mu)
print(sig)



n = int(input())
nums = []
for i in range(0,n):
    nums.append(int(input()))
total = [0,0,0]
zero_pos = [-1,-1,-1]
zero_seq = 0
max_ = 0
for i in range(0,n):
    if 0 == nums[i]:
        if zero_seq < 2:
            zero_seq += 1
        else:
            if sum(total) > max_:
                max_ = sum(total)
            temp = total[1:3]
            temp.append(0)
            total = temp
    else:
        total[zero_seq] += nums[i]

if sum(total) > max_:
    max_ = sum(total)
print(max_)