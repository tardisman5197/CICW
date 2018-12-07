# CICW
Computer Intelligence Coursework

## Prerequisites

[Go](https://golang.org/doc/install) needs to be installed.

## Building the project

1. Navigate to the project directory
2. Run the command:
    ```
    go build
    ```

This will output an executable for your opperating system.

## Running

There are two ways of running the program

### With Build

1. In the command promt, navigate to the executable
2. Run the command:
    ```
    ./CICW.exe
    ```
3. By default it will run a demo, with seed 0.

### Without Building

1. In the command promt, navigate to the project
2. Run the command:
    ```
    go run algorithms.go main.go particle.go price.go pricingProblem.go
    ```
3. By default it will run a demo, with seed 0.

## Help

If you add the flag `-h` when running the program will display the help text

```
Usage of CICW.exe:
  -alg string
        Sets algorithm to run: 'PSO', 'AIS' (default "PSO")
  -c    Collects data over a selection of seeds for a certain algorithm
  -ext string
        Sets the file extention to the csv file (default "final")
  -seed int
        Set the default seed
```