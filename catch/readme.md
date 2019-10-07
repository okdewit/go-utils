# Catch

This is a utility tool to make error catching a bit less verbose.

This is an abstraction which is not idiomatic go -- But sometimes, you just want to throw all your errors in a "logging bucket" and be done with it.

Instead of the default verbose error handling:

```go
user, err := getUser()

if (err != nil) {
    log.Println(err.Error())
}

color, err := user.GetFavoriteColor()

if (err != nil) {
    log.Println(err.Error())
}
```

You can defer a "Catch" at the top of a function, and throw using a "Check" function

```go
defer catch.Catch(func(err error) {
    log.Println(err.Error())
})

user, err := getUser()
catch.Check(err)

color, err := user.GetFavoriteColor()
catch.Check(err)
```
