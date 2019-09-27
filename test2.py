str = input()

while i < len(str):
    a = str[i]
    if ord(a) >= ord('0') and ord(a) <= ord('9') and i+1 < len(str) and str[i+1]=='%':
        Mul.append(ord(a) - ord('0'))
        Front.append(front)
        front = ""
        i += 1
    elif a == '#':
        mul = Mul.pop()
        mid_temp = ""
        for j in range(0,mul):
            mid_temp += front
        front = Front.pop() + mid_temp

    else:
        front += a
    i += 1
print(front)