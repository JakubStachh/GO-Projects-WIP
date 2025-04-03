# 📌 Merge Two Sorted Arrays in Go

## 🚀 Description
This Go program merges two sorted arrays into a single sorted array efficiently using a two-pointer technique.

## 🔍 How It Works
### The `mergeArrays()` function:

- Iterates through both arrays, comparing elements and adding the smaller one to the result.

- Once one array is exhausted, the remaining elements of the other array are added.

- Returns a fully merged and sorted array.

## 📂 Example
### Input:
```sh
arr1 := []int{1, 3, 5, 7}
arr2 := []int{2, 4, 6, 8}
```

### Processing:
- Compares elements and merges step by step:

```sh
[1] → [1, 2] → [1, 2, 3] → [1, 2, 3, 4] → [1, 2, 3, 4, 5] → [1, 2, 3, 4, 5, 6] → [1, 2, 3, 4, 5, 6, 7] → [1, 2, 3, 4, 5, 6, 7, 8]
```
The final result is `[1, 2, 3, 4, 5, 6, 7, 8]`.

### Output:
```sh
[1 2 3 4 5 6 7 8]
```
## 🛠️ Implementation Details
- **Two-pointer technique**:

 - One pointer (`i`) for `arr1`, another (`j`) for `arr2`.

 - The smaller value at each step is added to `result`.

- **Handling leftovers**:

- If one array is fully traversed, the remaining elements of the other array are added.
