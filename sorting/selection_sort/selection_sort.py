# find the smallest integer
def find_smallest(arr): 
    smallest = arr[0]
    smallest_index = 0
    for i in range (1, len(arr)):
        if(arr[i]< smallest):
            smallest = arr[i]
            smallest_index = i
    return smallest_index

# sort by finding smallest integer and insert the same into new array
def selection_sort(arr):
    newArr = []
    copiedArray = list(arr)
    for i in range(len(copiedArray)):
        smallest = find_smallest(copiedArray)
        newArr.append(copiedArray.pop(smallest))
        print("new array after process")
        print(newArr)
    return newArr

print(selection_sort([5,7,1,2,8,3,9]))