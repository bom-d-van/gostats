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

Access first field:

![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-first.png)
![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-first-plot.png)

Access last field:

![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-last.png)
![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/fields-last-plot.png)

### Access methods via reflection

Similar to fields in structs while fields is replaces by methods, a benchmark example:

```
type TenFields struct {}

func (f TenFields) Field1()  {}
func (f TenFields) Field2()  {}
func (f TenFields) Field3()  {}
func (f TenFields) Field4()  {}
func (f TenFields) Field5()  {}
func (f TenFields) Field6()  {}
func (f TenFields) Field7()  {}
func (f TenFields) Field8()  {}
func (f TenFields) Field9()  {}
func (f TenFields) Field10() {}

func BenchmarkReflectTenFields(b *testing.B) {
	foo := reflect.ValueOf(TenFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field1").Call(nil)
	}
}

func BenchmarkStandbardTenFields(b *testing.B) {
	foo := TenFields{}
	for i := 0; i < b.N; i++ {
		foo.Field1()
	}
}

func BenchmarkReflectLastTenFields(b *testing.B) {
	foo := reflect.ValueOf(TenFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field10").Call(nil)
	}
}

func BenchmarkStandbardLastTenFields(b *testing.B) {
	foo := TenFields{}
	for i := 0; i < b.N; i++ {
		foo.Field10()
	}
}
```

Access first method:

![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/methods-first.png)
![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/methods-first-plot.png)

Access last method:

![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/methods-last.png)
![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/methods-last-plot.png)

## Different types have different performances

Benchmark example:

```
func NoIntArg()                                                                  {}
func OneIntArg(arg int)                                                          {}
func FiveIntArgs(arg1, arg2, arg3, arg4, arg5 int)                               {}
func TenIntArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 int) {}
func FifteenIntArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 int) {
}
func TwentyIntArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 int) {
}

func BenchmarkNoIntArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoIntArg()
	}
}

func BenchmarkOneIntArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneIntArg(0)
	}
}

func BenchmarkFiveIntArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveIntArgs(0, 0, 0, 0, 0)
	}
}

func BenchmarkTenIntArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenIntArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	}
}

func BenchmarkFifteenIntArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenIntArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	}
}

func BenchmarkTwentyIntArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyIntArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	}
}
```

![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/funcargs.png)
![](https://raw.githubusercontent.com/bom-d-van/gostats/master/imgs/funcargs-plot.png)

## Conclusion

Reflection is indeed slow.

Simple is better. When we design data structure, we should always prefer simple solutions. If int could

The result of this exploration is actually not surprising at all. We do know reflection is slow and map is costly.

These not-yet-complete benchmarks could also serve as a simple guide line when designing programs.

This post has no intension to discourage any usage of so-called costly things. But I think as a programmer, when we are designing something, we should understand what we are actually doing. Using inappropriate data structure or over complicated design is wrong indeed.
