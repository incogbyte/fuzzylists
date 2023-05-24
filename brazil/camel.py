# read a file and print the contents
# in camel case
def camel():
    with open('api-br.txt', 'r') as f:
        for line in f:
            print(line.title(), end='')
        print()


camel()