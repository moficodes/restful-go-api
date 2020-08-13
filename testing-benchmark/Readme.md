# Testing and Benchmarking

## Testing

There are two main kinds of tests

1. Functional Testing types
    - Unit Testing
    - Integration Testing
    - System Testing
    - Sanity Testing
    - Smoke Testing
    - Interface Testing
    - Regression Testing
    - Beta/Acceptance Testing

2. Non-functional Testing
    - Performance Testing
    - Load Testing
    - Stress Testing
    - Volume Testing
    - Security Testing
    - Compatibility Testing
    - Install Testing
    - Recovery Testing
    - Reliability Testing
    - Usability Testing
    - Compliance Testing
    - Localization Testing

We will be mainly looking at unit testing our REST API in this section. Some of these testing are out of scope for our code altogether. For example, System Testing or Regression Testing. 

## Unit Testing

Unit test checks the functionality of single unit of code and confirms agains know output that the result is consistent. 

Go has a built-in package [httptest](https://golang.org/pkg/net/http/httptest) for testing http methods.

Look at `courses_test.go` for some examples of http tests in go. I leave writing test for the other handlers as an exercise you you.

## Running the Test

To run the code in this section

```bash
git checkout origin/testing-benchmarking-01
```

If you are not already in the folder

```bash
cd testing-benchmark
```

```bash
go test -v ./...
```


