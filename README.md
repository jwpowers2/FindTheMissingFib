# Find The Missing Fib

## create a Fibonacci sequence

```go
    f := NewFibGame()
	f.Fibs = f.MakeFib(21)
```

## remove a number from the sequence

```go
    f.RemoveNumber(2)
```

## search for which number was removed and return the numbers above and below in the sequence

```go
    f.FindMissingFib(f.Fibs)
```
