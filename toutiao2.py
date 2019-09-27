s = input()
wordDict = input().split()
maxlen=0
for word in wordDict:
    if len(word)>maxlen:
        maxlen=len(word)
res=[0]*len(s)
for i in range(len(s)):
    p=i
    while(p>=0 and i-p<=maxlen):
        if (res[p]==1 and s[p+1:i+1] in wordDict) or (p==0 and s[p:i+1] in wordDict):
            res[i]=1
            break
        p-=1
if res[-1] == 1:
    print("True")
else:
    print("False")