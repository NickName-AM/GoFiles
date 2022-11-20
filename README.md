# Create a bunch of files

## Working with files in Go
### Same as https://github.com/NickName-AM/PyFiles, but in Go


## What's new?
- Faster
- Better

## Example
This will create 100 files with extension php. (2 threads will create 50 files each)
```
go run Main.go -n 100 -t 2 -e php
```

## Help
```
  -e string
        Extension to use (default "txt")
  -n int
        Number of files to create (Default: 0) (default 1)
  -t int
        Number of threads(Default: 1) (default 1)
  -v    Verbosity (Default: Off)
```