# 📌 QuickSort in Go

## 🚀 Description
This Go program implements the QuickSort algorithm, which is a highly efficient, comparison-based sorting algorithm. It recursively partitions the array around a pivot and sorts the elements.

## 🔍 How It Works
### The `quickSort()` function:

- **Base case**: If the array has fewer than two elements, it's already sorted.

- **Pivot Selection**: The first element of the array is chosen as the pivot.

- **Partitioning**: The array is partitioned into two subarrays:

    - Elements less than or equal to the pivot.

    - Elements greater than the pivot.

- **Recursive Sorting**: The function is then recursively called to sort the left and right subarrays.

## 🛠️ Implementation Details
- **Partitioning**:

  - Two pointers (`left` and `right`) are used to find elements on opposite sides of the pivot that need to be swapped.

- **Recursion**: The function is applied recursively to both the left and right subarrays to ensure all elements are sorted.

## 📂 Example
### Input:
```sh
arr := []int{10, 7, 8, 9, 1, 5}
fmt.Println("Sorted array:", quickSort(arr))
```
### Processing:
1. Start with pivot `10`.

2. Partition the array into elements less than `10` and greater than `10`.

3. Recursively sort the two partitions.

4. Return the final sorted array.

### Output:
```sh
Sorted array: [1 5 7 8 9 10]
```
