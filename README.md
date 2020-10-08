# go-cobra-escaping

Small repo just to help visualize how cobra escapes incoming arguments.

With the cobra escaping intact:
```
$ go run main.go with "hello\nworld"
```

With the cobra escaping countered using `strings.ReplaceAll`:
```
$ go run main.go without "hello\nworld"
```
