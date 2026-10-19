# Go Big O Notation

A practical introduction to Big O Notation using Go examples. This repository explores common time complexities, compares algorithms and data structures, and measures actual performance through Go benchmarks.

## Overview

Big O Notation describes how an algorithm's resource requirements grow as the input size increases. Understanding this growth helps developers evaluate algorithm efficiency and make better decisions when building scalable applications.

This repository provides runnable examples and tests to demonstrate the differences between common time complexities.

## Topics

| Directory | Complexity / Topic | Description |
|---|---|---|
| `01-constant` | `O(1)` | Constant-time operations |
| `02-logarithmic` | `O(log n)` | Binary Search |
| `03-linear` | `O(n)` | Linear Search |
| `04-linearithmic` | `O(n log n)` | Merge Sort |
| `05-quadratic` | `O(n²)` | Nested loops and pair comparisons |
| `06-exponential` | `O(2ⁿ)` | Naive recursive Fibonacci |
| `07-search` | Search Algorithms | Compare Linear Search and Binary Search |
| `08-data-structures` | Data Structures | Compare slice and map lookups |
| `09-benchmark` | Benchmarking | Measure lookup performance using Go benchmarks |

## Prerequisites

- Go installed on your machine.
- Basic familiarity with Go syntax and functions.
- A terminal for running examples and tests.

## Getting Started

Clone the repository:

```bash
git clone https://github.com/andriantp/golang-big-o.git
cd golang-big-o
```

Replace `USERNAME` with the GitHub username that owns the repository.

Run an individual example:

```bash
go run ./01-constant
go run ./02-logarithmic
go run ./03-linear
go run ./04-linearithmic
go run ./05-quadratic
go run ./06-exponential
```

Run the search tests:

```bash
go test ./07-search -v
```

Run the data structure tests:

```bash
go test ./08-data-structures -v
```

Run the benchmarks:

```bash
go test ./09-benchmark -bench=. -benchmem
```

## Benchmark Results

The benchmark compares slice and map lookups using datasets of different sizes.

The following results were measured on a machine running Linux on an Intel Core i5-1235U processor.

| Dataset Size | Slice Lookup | Map Lookup |
|---:|---:|---:|
| 100 | 175.4 ns/op | 10.03 ns/op |
| 1,000 | 1,662 ns/op | 8.489 ns/op |
| 10,000 | 14,192 ns/op | 10.46 ns/op |

These measurements illustrate two different growth patterns:

- **Slice lookup:** typically `O(n)` in the worst case because elements may need to be checked one by one.
- **Map lookup:** average `O(1)` for key lookup, although worst-case behavior can be linear.

Benchmark results depend on hardware, Go version, data distribution, and implementation details. They should not be treated as universal performance guarantees.

## Key Takeaways

- Big O describes how algorithm complexity grows with input size.
- Different algorithms can solve similar problems with different complexity.
- Data structure selection affects lookup performance and memory usage.
- Benchmarking measures actual implementation performance, complementing theoretical analysis.
- A lower complexity does not automatically guarantee better performance for every workload.

## Related Article

This repository accompanies the Medium article:

**[Golang] Big O Notation**

The article introduces the concepts behind the examples, while this repository provides the Go implementations and experiments.

