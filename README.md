# Turkish Identification Number Validator
A library to validate Turkish identification number. It provides 2 functions.

**Example**
```
id := "43848743180"
if !trid.Valid() {
    log.Printf("id is not valid")
}
```

If you want more details about the error use Validate function instead:

```
id := "43848743180"
if err := trid.Validate(); err!=nil {
    log.Printf("id is not valid: %s", err.Error())
}
```

## References
- [Turkish Identification Number on Wikipedia](https://en.wikipedia.org/wiki/Turkish_Identification_Number)