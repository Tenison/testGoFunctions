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
#### To Contribute to this Project 
- Star Project :)
- Fork Project 
- git clone "git url" to local repo
- git remote -v :: 'To view all upstream branches'
- git remote add upstream "original project source url" :: 'So you push and pull upstream updates locally'
- git pull upstream main :: 'Pull upstream main branch updates"