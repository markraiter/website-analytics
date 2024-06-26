# User Behavior Analysis

![GitHub](https://img.shields.io/github/license/markraiter/website-analytics)  ![GitHub top language](https://img.shields.io/github/languages/top/markraiter/website-analytics)  [![Go Report Card](https://goreportcard.com/badge/github.com/markraiter/website-analytics)](https://goreportcard.com/report/github.com/markraiter/website-analytics)

This is a Go project that analyzes user behavior data from CSV files. It loads data from two days and finds users who visited new pages on the second day. Time estimated - 4 hours.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.22.0)

### Installing

A step by step series of examples that tell you how to get a development environment running.

1. Clone the repository.
2. Follow the instructions to install [Taskfile](https://taskfile.dev/ru-ru/installation/) utility.
3. Make sure that you have `day1.csv` and `day2.csv` files in the project directory.
4. Run the app with `task run`.

## Running the tests

1. Run the tests with `task test`

## Built With

- [Go](https://golang.org/) - The programming language used.