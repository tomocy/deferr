# deferr

deferr provides a way to simplify error handling with defer.  

## Installtion
```
go get github.com/tomocy/deferr
```

## Expected usage
```go
func foo() (err error) {
	defer deferr.Format(&err, "failed to foo")

	a, err := doA()
	if err != nil {
		return err
	}

	b, err := doB()
	if err != nil {
		return err
	}

	return nil
}
```

This example reduces the conventional duplication of error handling like the below.
```go
func foo() error {
	a, err := doA()
	if err != nil {
		return fmt.Errorf("failed to foo: %v", err)
	}

	b, err := doB()
	if err != nil {
		return fmt.Errorf("failed to foo: %v", err)
	}

	return nil
}
```