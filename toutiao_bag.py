tmp = input().split()
c = int(tmp[0])
m = int(tmp[1])

w = []
v = []

for i in range(0,m):
    tmp = input().split()
    w.append(int(tmp[0]))
    v.append(int(tmp[1]))

dp = [[0] * (c + 1) for i in range(len(w) + 1)]
 
for i in range(1, len(w) + 1):
    for j in range(1, c + 1):
        if j >= w[i - 1]:
            dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - w[i - 1]] + v[i - 1])
        else:
            dp[i][j] = dp[i - 1][j]
print(max(dp[len(w)]))