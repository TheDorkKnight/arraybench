# arraybench

Compares benchmarks for returning slices xor arrays in the following situations:

 - return slice that is a wrapper around an existing array *(can be optimized out by compiler)*
 - return slice is a newly allocated copy of an existing array *(involves heap allocation)*
 - return array that exists *(can be optimized out by compiler)*
 - return copy of array *(involves stack allocation)*
