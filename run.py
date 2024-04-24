def prime_factors(n):
    factors = []
    p = 2
    while p * p <= n:
        while n % p == 0:
            factors.append(p)
            n //= p
        p += 1
    if n > 1:
        factors.append(n)
    return factors

# 示例
if __name__ == "__main__":
    list = prime_factors(9)

    print(list)
    print(set(list))