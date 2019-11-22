This Weekends Challenge:

```
Your goal is to create a function that removes the first and last letters of a string. Strings with two characters or less are considered invalid: provide an appropriate message for that case.
```

This repo contains unit tests and a main.go that will allow you to run anywords you want to run through and return the results or error.

# To Run Unit Tests
```
cd internal/stringpeeler
```
```
go test -cover
```
This should return 100% coverage

# To peel your own word(s)
```
cd cmd
```
```
go run main.go <words to peel with a space between each word>
```

