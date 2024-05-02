# Installation Guide

### Via Docker

Prerequisites:

1. Have [Docker](https://www.docker.com/ 'Docker official website') installed

```
git clone git@github.com:TineoC/worldtech-test.git
docker build -t worldtech-test
docker run worldtech-test
```

### Via CMD

Prerequisites:

1. Have [Golang](https://go.dev/ 'Golang official docs') installed

```
git@github.com:TineoC/worldtech-test.git
go run main.go
```

### Output

```
![Project Output](https://github.com/TineoC/worldtech-test/blob/main/image/README/1714658490723.png "Project Output")

```

## Errors Fixed

✨ Added WaitGroups for handling async tasks on the getData subroutine.

✨ Added error handling for the results from the getData request.

🐛 Removed os.Exit function for preventing unexpected exit while processing async tasks.

🐛 Fixed "location" url value parsing inside getData() method.

🐛 Pass line argument into getData Goroutine.

✨ Containerize the project
