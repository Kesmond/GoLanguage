def sum(x, y, z=None):
    if z is None:
        return x+y
    else:
        return x+y+z

x = 10
y = 10
z = 10
print(f"x = {x}\ny = {y}\nz = {z}\n")
print("Sum of x and y: ", sum(x,y))
print("Sum of x y z: ", sum(x,y,z))
