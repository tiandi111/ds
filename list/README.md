# List

## Benchmark1 - Index Deletion: SingleLinkedList vs ArrayList(slice)

ArrayList deletion is implemented by index reslice:
```go
*l.arr = append((*l.arr)[:n], (*l.arr)[n+1:]...)
```
**Description**:<br/>
BenchmarkListSize = 100000 // list has 100000 elements<br/>
BenchmarkDelReps  = 50000 // repeat deletion for 50000 reps<br/>
BenchmarkDelHeadPosition   = 1 // delete from 10%<br/>
BenchmarkDelMiddlePosition = 5 // delete from 50%<br/>
BenchmarkDelTailPosition   = 9 // delete from 90%<br/>
BenchmarkDelDeNominator    = 10 // helper for integer math<br/>

**Result**:<br/>
BenchmarkGenericArrayList_Del_Middle-8             50000             20461 ns/op<br/>
BenchmarkGenericLinkedList_Del_Middle-8                1        8212797700 ns/op<br/>
BenchmarkGenericArrayList_Del_Head-8                   1        7457259300 ns/op<br/>
BenchmarkGenericLinkedList_Del_Head-8                  1        1436253600 ns/op<br/>
BenchmarkGenericArrayList_Del_Tail-8            1000000000               0.12 ns/op<br/>
BenchmarkGenericLinkedList_Del_Tail-8                  1        14998345200 ns/op<br/>

