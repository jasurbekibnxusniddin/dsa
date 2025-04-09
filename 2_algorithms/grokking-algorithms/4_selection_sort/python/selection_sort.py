def min_element(l) -> int:

    min_element = l[0]
    min_index = 0

    for i in range(1, len(l)):

        if l[i] < min_element:

            min_element = l[i]
            min_index = i

    return min_index

def selection_sort(l) -> list:

    new_list = []

    for i in range(len(l)):

        min_index = min_element(l)

        new_list.append(l.pop(min_index))
    
    return new_list


print(selection_sort([5, 3, 6, 2, 10]))