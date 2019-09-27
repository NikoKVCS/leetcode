
def median(A, B):
    m, n = len(A), len(B)
    if m > n:
        A, B, m, n = B, A, n, m
    if n == 0:
        raise ValueError

    imin, imax, half_len = 0, m, int((m + n + 1) / 2)
    print((m + n + 1) / 2, "  ", int((m + n + 1) / 2), "\n")
    while imin <= imax:
        i = int((imin + imax) / 2)

        print(((imin + imax) / 2), "  ", int((imin + imax) / 2), "\n")
        j = half_len - i
        if i < m and B[j-1] > A[i]:
            # i is too small, must increase it
            imin = i + 1
        elif i > 0 and A[i-1] > B[j]:
            # i is too big, must decrease it
            imax = i - 1
        else:
            # i is perfect

            if i == 0: max_of_left = B[j-1]
            elif j == 0: max_of_left = A[i-1]
            else: max_of_left = max(A[i-1], B[j-1])

            if (m + n) % 2 == 1:
                return max_of_left

            if i == m: min_of_right = B[j]
            elif j == n: min_of_right = A[i]
            else: min_of_right = min(A[i], B[j])

            return (max_of_left + min_of_right) / 2.0

A=[1,3]
B=[2]
A = [1,4,6,8,9,13,15,16,17]
B = [7,8,9,10,14,16,19,20]
C = median(A,B)
print(C)