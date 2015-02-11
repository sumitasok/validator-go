```
v := On("string").Required().Min(7).Max(10)
v := On(123).Range(122, 124)
v := On(time.Time{}).IsAfter(time.Time{})
```

alt:

```
v := On("string").Key("name").Required("custom message required").Min(7, "min custom message").Max(10, "max custom message")
```

```v.Errors()```

returns

```
[]error{
	errors.New("name: custom message required"),
	errors.New("name: min custom message")
}
```

#### Match

```
vMatch := On("email@exmple.com").Required().Match("([a-zA-Z0-9])+(@)([a-zA-Z0-9])+((.)[a-zA-Z0-9])+")
```

Or

```
vEmail := On("email@example.com").Required().Email()
```

Or

```
vMatch := On("email@exmple.com").Required().Match(validator.EMAIL_PATTERN)
```

TO-DO:

- move validation code into methods
- move validation code into struct and add check compatibility