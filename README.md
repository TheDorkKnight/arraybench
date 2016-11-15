# arraybench

Compares benchmarks for returning slices xor arrays in the following situations:

 - return slice that is a wrapper around an existing array *(can be optimized out by compiler)*
 - return slice is a newly allocated copy of an existing array *(involves heap allocation)*
 - return array that exists *(can be optimized out by compiler)*
 - return copy of array *(involves stack allocation)*

## Sample Results

```
BenchmarkSlice-8                        30000000                57.3 ns/op
BenchmarkSliceWrappingArray-8           100000000               14.4 ns/op
BenchmarkArray-8                        100000000               14.5 ns/op
BenchmarkArrayCopy-8                    50000000                29.7 ns/op
```
