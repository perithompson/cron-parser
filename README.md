# cron-parser
Cron-Parser is a Simple CLI written in Go to parse cron expressions.

## Prerequisites
- [Go v1.18.3](https://go.dev/doc/install)

## Description
This project has been written in Go and tested with version 1.18.3.  For speed and consistency with other Go CLI tools, 
it was initialised using the [Cobra CLI](https://github.com/spf13/cobra).  


## Getting Started


Build the project from within the Project Directory
``` sh
cd project_dir #!replace with project location

go mod tidy
go build -o cron-parser .

chmod +x ./cron-parser
```

## Usage

``` sh
./cron-parser "23 0-20/2 * * * /usr/bin find"     
```

``` sh
minute        23
hour          0 2 4 6 8 10 12 14 16 18 20
day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   0 1 2 3 4 5 6
command       /usr/bin/find
```

## Contributing

In order to add further CLI functions, you should install the `cobra-cli` following their [instructions](https://github.com/spf13/cobra-cli/blob/main/README.md) and then using `cobra-cli add <command>` as described [here](https://github.com/spf13/cobra-cli/blob/main/README.md#add-commands-to-a-project)
