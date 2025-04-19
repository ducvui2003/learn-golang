# Notes

- File need a main func to run via `go run`
- Go treat all file in same folder is a part of package. So running `go init` need only one main func declared in same folder, go will merge all file in same folder when run. If you want run main func in each file, you can wrap each file in another folder like this repository
- If file has any declared variable not use, this file can not run
