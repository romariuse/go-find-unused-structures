Search structures in code. Useless for find unused public structures (errors, translations, etc...)

```go
structures := finder.New(finder.PathConfig{
    Path:      "./internal/errorEnum",
    Recursive: true,
})
missing, err := structures.Unused(finder.PathConfig{
    Path:      "./",
    Recursive: true,
}
if err != nil {
	panic(err)
}
fmt.Println(missing)
```
