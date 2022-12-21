# GoFiles

### Create a bunch of files quickly
### Same as [this](https://github.com/NickName-AM/PyFiles), but in Go


## What's new?
- Faster
- Can add prefix/suffix
- Uses threads

## Installation
```
git clone https://github.com/NickName-AM/GoFiles.git
cd GoFiles
go build Main.go
```

## Example
This will create 100 files with the extension of php. (2 threads will create 50 files each)
```
./Main -n 100 -t 2 -e php
```

This will create 50 files with prefix of "foo"
```
./Main -n 50 -prefix foo
```

This will create 50 files with suffix of "bar"
```
./Main -n 50 -prefix bar
```

Or you can use both
```
./Main -n 101 -prefix foo -suffix bar
```

## Help
```
  -dc string
    	Write custom data to each file (Not to be used with -dl)
  -dl int
    	Length of random data to write in each file (Default: 0)
  -e string
    	Extension to use (Default: txt (default "txt")
  -fl int
    	Length of the filename (Default: 10) (default 10)
  -n int
    	Number of files to create (Default: 1) (default 1)
  -prefix string
    	Prefix
  -suffix string
    	Suffix
  -t int
    	Number of threads (Default: 1) (default 1)
  -v	Verbosity (Default: Off)
```

## To-Do
```
Create files from a wordlist w/ threading
```