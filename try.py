def divide(a, b):
    if b == 0:
        return (None, "division by zero")  # (result, error)
    return (a / b, None)

# Usage:
a = float(input("Numerator: "))
b = float(input("Denominator: "))

result, err = divide(a, b)
if err:
    print(f"Error: {err}")
else:
    print(f"Result: {result}")
