# Golang-FlagDir

Golang-FlagDir is a simple tool written in Go that allows creating directories with names based on user-defined flags provided through the command line.

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg) ![Go Gopher](https://golang.org/lib/godoc/images/gopher.png)

## Features

- Create directories with names specified by the user via command-line flags.
- Easy-to-use interface.

## Technologies Used

- [Go](https://golang.org/) ![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

## How to Use

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.

### Steps to Build and Run

1. Clone this repository:

    ```sh
    git clone https://github.com/WesleyBSa/Golang-FlagDir.git
    cd GoDirManager
    ```

2. Initialize the Go module:

    ```sh
    go mod init GoDirManager
    ```

3. Build the program:

    ```sh
    go build -o godirmanager.exe
    ```

4. Run the program:

    ```sh
    .\godirmanager.exe -channelName "YourDirectoryName"
    ```

## Example Usage

```sh
.\godirmanager.exe -channelName "TestChannel"
