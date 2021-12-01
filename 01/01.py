def window (arr: list[int], n: int) -> list[int]:
    for i in range(len(arr) - n + 1):
        yield arr[i:i+n]


if __name__ == '__main__':

    ### Part 1 ###

    count = 0
    last = 0

    with open("inputs.txt", "r") as f:
        inputs = [int(x) for x in f.readlines()]

    last = inputs[0]
    for line in inputs:

        current = line
        if current > last:
            count += 1
        last = current

    print("Part 1:")
    print(count)

    ### Part 2 ###

    count = 0
    windows = window(inputs,3)
    last_sum = sum(next(windows))

    for group in windows:

        current_sum = sum(group)
        if current_sum > last_sum:
            count += 1
            print(current_sum, last_sum)

        last_sum = current_sum

    print("Part 2:")
    print(count)