# 📌 Merge Overlapping Intervals in Go

## 🚀 Description
This Go program merges overlapping intervals from a given list of interval pairs. It ensures that overlapping intervals are combined into a single interval.

##🔍 How It Works
### The `mergeIntervals()` function:

- Sorts the intervals based on their start times.

- Iterates through the intervals and merges overlapping ones.

- Returns the merged list of intervals.

## 📂 Example
### Input:
```sh
intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
```
### Processing:
- `[1, 3]` and `[2, 6]` overlap → merged into `[1, 6]`.

- `[8, 10]` and `[15, 18]` remain as they are.

### Output:
```sh
Merged intervals: [[1 6] [8 10] [15 18]]
```

## 🛠️ Implementation Details
### - **Sorting**: Intervals are sorted based on their start time using `sort.Slice()`.

### - **Merging**:

- If an interval overlaps with the previous one, they are combined.

- Otherwise, the interval is added separately to the result.
