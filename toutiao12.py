wb = input().split()
w = int(wb[0])
b = int(wb[1])

sum = w+b
ans = 0
cnt = [0 for i in range(0,1010)]
f = [[0 for j in range(0,4*123456)] for i in range(0,2)]
f[0][r] = 1

now = 0
pre = 1

for i in range(1, 1001):
    now 
 
	for (int i = 1; i <= 1000; i++) {
		now ^= 1;
		pre = (now+1)%2; 
		cnt[now] = cnt[pre]+i;
 
		if (cnt[now] > sum) {
			for (int j = 0; j <= r; j++) {
				ans += f[pre][j];
				if (ans >= modn)  ans -= modn;
			}
 
			break;
		}
 
		for (int j = 0; j <= r; j++) {
			f[now][j] = 0;
 
			if (sum-cnt[pre]-j >= i)
				f[now][j] = f[pre][j];
 
			f[now][j] += f[pre][j+i];
			if (f[now][j] >= modn)  f[now][j] -= modn;
		}
	}
 
	printf("%d\n", ans);
}
 
int main() {
	init();
	solve();
}