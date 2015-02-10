	v := On("string").Required().Min(7).Max(10)
	v := On(123).Range(122, 124)
	v := On(time.Time{}).IsAfter(time.Time{})
