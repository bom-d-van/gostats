# Aspects of Go: How slow is reflection and How types of a function affect its performance

Bellow are some benchmarks designed to explore the performance of reflection package and of arguments of different types in functions, for Go.

## How slow is reflection

### Access fields in struct via reflection

There are 7 kind of structs, among which the difference is their field number, starting from 1 field, then 5 fields, 10 fields, ..., up to 30 fields. There are four benchmarks, two of reflection and two of non-reflection.

A example of these benchmarks:

```
type FiveFields struct {
	Field1 int
	Field2 int
	Field3 int
	Field4 int
	Field5 int
}

func BenchmarkReflectFiveFields(b *testing.B) {
	foo := reflect.ValueOf(FiveFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field1").String()
	}
}

func BenchmarkStandbardFiveFields(b *testing.B) {
	foo := FiveFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field1
	}
}

func BenchmarkReflectLastFiveFields(b *testing.B) {
	foo := reflect.ValueOf(FiveFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field5").String()
	}
}

func BenchmarkStandbardLastFiveFields(b *testing.B) {
	foo := FiveFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field5
	}
}
```

And running result is:

Reflection:

![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-reflection.png)
![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-reflection-plot.png)

Non-reflection:

![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-standard.png)
![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-standard-plot.png)

### Access methods via reflection

## Different types have different performances

## Conclusion

Reflection is indeed slow.

Simple is better. When we design data structure, we should always prefer simple solutions. If int could

The result of this exploration is actually not surprising at all. We do know reflection is slow and map is costly.

These not-yet-complete benchmarks could also serve as a simple guide line when designing programs.

This post has no intension to discourage any usage of so-called costly things. But I think as a programmer, when we are designing something, we should understand what we are actually doing. Using inappropriate data structure or over complicated design is wrong indeed.
