# deferr

deferr provides a way to simplify error handling with defer.  

## Installtion
```
go get github.com/tomocy/deferr
```

## Example
```go
func ExampleFormat() {
	err := func() (err error) {
		defer deferr.Format(&err, "failed to foo")

		return errors.New("bar")
	}()

	fmt.Println(err)
	// Output:
	// failed to foo: bar
}

func ExampleWrapf() {
	err := func() (err error) {
		defer deferr.Wrapf(&err, "failed to foo")

		return errors.New("bar")
	}()

	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
	// Output:
	// failed to foo: bar
	// bar
}
```

This behavior can be customized by defining your `VerbMap`.  
```go
func ExampleCustomizedFormat() {
	err := func() (err error) {
		defer deferr.VerbMap{
            deferr.KeyFormat: {
                Flag: '#', Verb: 'v',
            },
        }.Format(&err, "failed to foo")

		return errors.New("bar")
	}()

	fmt.Println(err)
	// Output:
    // failed to foo: &errors.errorString{s:"bar"}
}
```