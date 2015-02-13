#### Usage

```
import "github.com/sumitasok/validator-go"
```
In code

```
v := validator.On("string").Required().Min(7).Max(10)
v := validator.On(123).Range(122, 124)
v := validator.On(time.Time{}).IsTimeAfter(time.Time{})
```

alt:

```
v := validator.On("string").Key("name").Required("custom message of your choice").Min(7, "min custom message").Max(10, "max custom message")
```


#### Errors

```
v.Errors
```

returns all errors

```
[]error{
	errors.New("name: custom message required"),
	errors.New("name: min custom message")
}
```

```
v.Error()
```
returns first `error` or `nil`

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

I am thankful to https://github.com/jamieomatthews/validation, as this package is inspired on it, when it comes to the Validator chained signature and type aggregation for matching"
