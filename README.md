# Image Conversion Tools

This tool allows you to convert images with different extensions on the command line.

## Command line Option

| Option | Value | Default |
| --- | --- | --- |
| -from | jpg, jpeg, png | jpg |
| -to | jpg, jpeg, png | png |
| -dir |  source directory | current directory |

## Example

### with go run
```
$ go run main.go -from png -to jpg -dir sample
```

### with go build
```
$ go build extchanger
$ ./extchanger -from png -to jpg -dir sample
```

## Specification

- Command line arguments (directory, extension before and after)
- Check command line arguments
- Use of user defined types
- Get all file paths with a specified extension in a specified directory
- Separate from the main package
- Use only self-made and standard packages
- Using Go Modules

## Reference
- [tenntenn/gohandson](https://github.com/tenntenn/gohandson/tree/master/imgconv/ja)
- [Gopher道場](https://github.com/gopherdojo/gopherdojo-studyroom)

## Other

I created this for my own Golang study.  
If anyone has any advice, feel free to share it with me.