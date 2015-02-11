```
v := On("string").Required().Min(7).Max(10)
v := On(123).Range(122, 124)
v := On(time.Time{}).IsAfter(time.Time{})
```

alt:

```
v := On("string").Key("name").Required("custom message required").Min(7, "min custom message").Max(10, "max custom message")
```

```
v.Errors()
```

returns

```
[]error{
	errors.New("name: custom message required"),
	errors.New("name: min custom message")
}
```

#### Match

```
vMatch := On("email@example.com").Required().Match("([a-zA-Z0-9])+(@)([a-zA-Z0-9])+((.)[a-zA-Z0-9])+")
```

Or

```
vEmail := On("email@example.com").Required().Email()
```

Or

```
vMatch := On("email@exmple.com").Required().Match(validator.EMAIL_PATTERN)
```

#### Time

time/thresholdTime is an object of `time.time`

```
time.Date(2011, time.November, 10, 23, 0, 0, 0, time.UTC)
```

In order to validate if the time provided is after a particular period

```
On(time).Min(thresholdTime)
```

Or

```
On(time).IsTimeAfter(thresholdTime)
```

In order to validate if the time provided is before a particular period

```
On(time).Max(thresholdTime)
```

Or

```
On(time).IsTimeBefore(thresholdTime)
```


TO-DO:

- add check compatibility