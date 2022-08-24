#### To Test functionationality found in Packages : 
#### Every package is functionationality I want to test
#### Read through the packages and import the package you want to test into the main.go file
- To implement the functions you want to test Eg: stringCharaters.ConvertStringToByteArray("hello")
```
    //Print to see results
    fmt.Println("Result is ", stringCharaters.ConvertStringToByteArray("hello"))
```
#### To Test/Run
- If you have makefile installed
```cmd
    make run-test
```
- else
```cmd
    go run main.go
```

- To run automated tests on all functions first
```cmd
    make test
```

