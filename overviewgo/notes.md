1. ### Pointer

- Get the variable reference: put an ampasand(`&`) in fron of it

```go
  hello := "hello"
  changeRef(*hello)
//   hello == "goodBye"
```

- Refer to an actual pointer: put an asterisk(`*`) in fron of it

```go
  func changeRef(s *string) {
    newStr := "goodBye"
    *s = newStr
  }
```

2. ### Variable, Function firstLetter

- Upercase: Useable outside of module
- Lowercase: Use only inside of module
