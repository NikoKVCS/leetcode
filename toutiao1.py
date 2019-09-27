wb = input().split()
w = int(wb[0])
b = int(wb[1])

h = 0
def choose(color, h, left_w, left_b):
    if (color == 1 and left_w - (h+1) < 0 ) or (color == 2 and left_b - (h+1) < 0 ):
        return h, 1
    h = h+1
    if color == 1:
        left_w -= h
    else:
        left_b -= h
    h1, k1 = choose(1, h, left_w, left_b)
    h2, k2 = choose(2, h, left_w, left_b)

    if h1 > h2:
        return h1, k1
    elif h1 < h2:
        return h2, k2
    else:
        return h1, k1+k2

    
h1, k1 = choose(1, 1, w-1, b)
h2, k2 = choose(2, 1, w, b-1)
if h1 > h2:
    print(h1,  k1)
elif h1 < h2:
    print(h2,  k2)
else:
    print(h1, (k1+k2)) 

