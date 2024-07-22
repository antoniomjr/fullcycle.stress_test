# Go Challenge for FullCycle Pós Go Expert

## Overview

This repository contains a Go application developed as a part of the "`fullcyle.stress-test`" (Go Challenge) from the Pós Go Expert.

## Challenge Requirements

Objective: Create a CLI system in Go to perform load tests on a web service. The user must provide the service URL, the total number of requests and the number of simultaneous calls.

The system must generate a report with specific information after carrying out the tests.

Parameter Entry via CLI:

--url: URL of the service to be tested.
--requests: Total number of requests.
--concurrency: Number of simultaneous calls.

Test Execution:

Make HTTP requests to the specified URL.
Distribute requests according to the defined level of competition.
Ensure that the total number of requests is fulfilled.
Report Generation:

Present a report at the end of the tests containing:
Total time spent executing
Total number of requests made.
Number of requests with HTTP status 200.
Distribution of other HTTP status codes (such as 404, 500, etc.).
Application execution:
We can use this application by making a call via docker. Ex:
docker run <your docker image> —url=http://google.com —requests=1000 —concurrency=10

## Project Structure
|-- main.go

## Instructions

To build and run this application, follow these steps:

#### Running the Go application by Docker
This will run the application
```bash
docker build -t loadtester .
docker run loadtester --url=http://httpbin.org/anything --requests=10000 --concurrency=600
```

#### Running the Go application locally
```bash
go build -o loadtester main.go

./loadtester --url=http://httpbin.org/anything --requests=10000 --concurrency=600
docker run loadtester --url=https://economia.awesomeapi.com.br/json/last/USD-BRL --requests=99900 --concurrency=100
```
> [!WARNING]
> docker run loadtester --url=http://httpbin.org/anything --requests=10000 --concurrency=600
><img width="675" alt="Screenshot 2024-07-22 at 20 30 28" src="https://github.com/user-attachments/assets/35a27ade-bd00-4f3c-aef4-0e33072126d9">


