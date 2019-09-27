
NM = input().split()
N = int(NM[0])
M = int(NM[1])


def seq(i,j):
    return i*M + j
matrix = []

for i in range(0,N):
    val = input().split()
    row = []
    for j in range(0,M):
        row.append(int(val[j]))
    matrix.append(row)

graph = [[] for i in range(0,N*M)]


for i in range(0,N):
    for j in range(0,M):
        if 0 == matrix[i][j]:
            continue
        relations = []
        if j+1 < M and matrix[i][j+1] == 1:
            relations.append(seq(i,j+1))
        if j-1 >= 0 and i+1 < N and matrix[i+1][j-1] == 1:
            relations.append(seq(i+1,j-1))
        if i+1 < N and matrix[i+1][j] == 1:
            relations.append(seq(i+1,j))
        if i+1 < N and j+1 < M  and matrix[i+1][j+1] == 1:
            relations.append(seq(i+1,j+1))
        graph[seq(i,j)] = relations

vis = [0 for i in range(0,N*M)]

def dfs(i):
    vis[i] = 1
    stack = []
    for item in graph[i]:
        stack.append(item)
    while len(stack) > 0:
        id = stack.pop()
        if vis[id] == 1:
            continue
        vis[id] = 1
        for item in graph[id]:
            stack.append(item)
count = 0
for i in range(0,N*M):
    if vis[i] == 1 or matrix[i//M][i%M] == 0:
        continue
    count += 1
    dfs(i)

print(count)

