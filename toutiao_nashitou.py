n = int(input())
stones = []
for i in range(0,n):
    stones.append(int(input()))
Our = 0
i = 0
my_turn = True
last_time = 1

def calc

while True:
    if my_turn:
        my_turn = False
        maxi = last_time * 2
        for i in range(1, maxi+1):
            tmp = stones[i:i+maxi]
            tmp_sum = sum(tmp)

