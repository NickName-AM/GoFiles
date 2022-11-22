# Create a bunch of files

## Working with files in Go
### Same as https://github.com/NickName-AM/PyFiles, but in Go


## What's new?
- Faster
- Better

## Installation
```
git clone https://github.com/NickName-AM/GoFiles.git
cd GoFiles
go build Main.go
```

## Example
This will create 100 files with the extension of php. (2 threads will create 50 files each)
```
go run Main.go -n 100 -t 2 -e php
```

## Help
```
  -dc string
        Write custom data to each file. (Not to be used with -dl)
  -dl int
        Length of random data to write in each file (Default: 0)
  -e string
        Extension to use (Default: txt (default "txt")
  -n int
        Number of files to create (Default: 1) (default 1)
  -t int
        Number of threads (Default: 1) (default 1)
  -v    Verbosity (Default: Off)
```