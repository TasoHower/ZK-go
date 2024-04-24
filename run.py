def extended_gcd(a, b):
    if b == 0:
        return a, 1, 0
    else:
        d, x, y = extended_gcd(b, a % b)
        return d, y, x - (a // b) * y

# 示例
if __name__ == "__main__":
    num1 = 35
    num2 = 3
    gcd_result, x, y = extended_gcd(num1, num2)

    print(f'{num1} 和 {num2} 的最大公约数是 {gcd_result}')
    print(f'满足贝祖等式的整数解为：{num1}*{x} + {num2}*{y} = {gcd_result}')