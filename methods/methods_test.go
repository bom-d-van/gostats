package method

import (
	"reflect"
	"testing"
)

type OneField struct {
}

func (o OneField) Method() {}

func BenchmarkReflectOneField(b *testing.B) {
	foo := reflect.ValueOf(OneField{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Method").Call(nil)
	}
}

func BenchmarkStandbardOneField(b *testing.B) {
	foo := OneField{}
	for i := 0; i < b.N; i++ {
		foo.Method()
	}
}

type FiveFields struct {
}

func (f FiveFields) Field1() {}
func (f FiveFields) Field2() {}
func (f FiveFields) Field3() {}
func (f FiveFields) Field4() {}
func (f FiveFields) Field5() {}

func BenchmarkReflectFiveFields(b *testing.B) {
	foo := reflect.ValueOf(FiveFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field1").Call(nil)
	}
}

func BenchmarkStandbardFiveFields(b *testing.B) {
	foo := FiveFields{}
	for i := 0; i < b.N; i++ {
		foo.Field1()
	}
}

func BenchmarkReflectLastFiveFields(b *testing.B) {
	foo := reflect.ValueOf(FiveFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field5").Call(nil)
	}
}

func BenchmarkStandbardLastFiveFields(b *testing.B) {
	foo := FiveFields{}
	for i := 0; i < b.N; i++ {
		foo.Field5()
	}
}

type TenFields struct{}

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

type FifteenFields struct {
}

func (ff FifteenFields) Field1()  {}
func (ff FifteenFields) Field2()  {}
func (ff FifteenFields) Field3()  {}
func (ff FifteenFields) Field4()  {}
func (ff FifteenFields) Field5()  {}
func (ff FifteenFields) Field6()  {}
func (ff FifteenFields) Field7()  {}
func (ff FifteenFields) Field8()  {}
func (ff FifteenFields) Field9()  {}
func (ff FifteenFields) Field10() {}
func (ff FifteenFields) Field11() {}
func (ff FifteenFields) Field12() {}
func (ff FifteenFields) Field13() {}
func (ff FifteenFields) Field14() {}
func (ff FifteenFields) Field15() {}

func BenchmarkReflectFifteenFields(b *testing.B) {
	foo := reflect.ValueOf(FifteenFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field1").Call(nil)
	}
}

func BenchmarkStandbardFifteenFields(b *testing.B) {
	foo := FifteenFields{}
	for i := 0; i < b.N; i++ {
		foo.Field1()
	}
}

func BenchmarkReflectLastFifteenFields(b *testing.B) {
	foo := reflect.ValueOf(FifteenFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field15").Call(nil)
	}
}

func BenchmarkStandbardLastFifteenFields(b *testing.B) {
	foo := FifteenFields{}
	for i := 0; i < b.N; i++ {
		foo.Field15()
	}
}

type TwentyFields struct {
}

func (tf TwentyFields) Field1()  {}
func (tf TwentyFields) Field2()  {}
func (tf TwentyFields) Field3()  {}
func (tf TwentyFields) Field4()  {}
func (tf TwentyFields) Field5()  {}
func (tf TwentyFields) Field6()  {}
func (tf TwentyFields) Field7()  {}
func (tf TwentyFields) Field8()  {}
func (tf TwentyFields) Field9()  {}
func (tf TwentyFields) Field10() {}
func (tf TwentyFields) Field11() {}
func (tf TwentyFields) Field12() {}
func (tf TwentyFields) Field13() {}
func (tf TwentyFields) Field14() {}
func (tf TwentyFields) Field15() {}
func (tf TwentyFields) Field16() {}
func (tf TwentyFields) Field17() {}
func (tf TwentyFields) Field18() {}
func (tf TwentyFields) Field19() {}
func (tf TwentyFields) Field20() {}

func BenchmarkReflectTwentyFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field1").Call(nil)
	}
}

func BenchmarkStandbardTwentyFields(b *testing.B) {
	foo := TwentyFields{}
	for i := 0; i < b.N; i++ {
		foo.Field1()
	}
}

func BenchmarkReflectLastTwentyFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field20").Call(nil)
	}
}

func BenchmarkStandbardLastTwentyFields(b *testing.B) {
	foo := TwentyFields{}
	for i := 0; i < b.N; i++ {
		foo.Field20()
	}
}

type TwentyFiveFields struct {
}

func (tff TwentyFiveFields) Field1()  {}
func (tff TwentyFiveFields) Field2()  {}
func (tff TwentyFiveFields) Field3()  {}
func (tff TwentyFiveFields) Field4()  {}
func (tff TwentyFiveFields) Field5()  {}
func (tff TwentyFiveFields) Field6()  {}
func (tff TwentyFiveFields) Field7()  {}
func (tff TwentyFiveFields) Field8()  {}
func (tff TwentyFiveFields) Field9()  {}
func (tff TwentyFiveFields) Field10() {}
func (tff TwentyFiveFields) Field11() {}
func (tff TwentyFiveFields) Field12() {}
func (tff TwentyFiveFields) Field13() {}
func (tff TwentyFiveFields) Field14() {}
func (tff TwentyFiveFields) Field15() {}
func (tff TwentyFiveFields) Field16() {}
func (tff TwentyFiveFields) Field17() {}
func (tff TwentyFiveFields) Field18() {}
func (tff TwentyFiveFields) Field19() {}
func (tff TwentyFiveFields) Field20() {}
func (tff TwentyFiveFields) Field21() {}
func (tff TwentyFiveFields) Field22() {}
func (tff TwentyFiveFields) Field23() {}
func (tff TwentyFiveFields) Field24() {}
func (tff TwentyFiveFields) Field25() {}

func BenchmarkReflectTwentyFiveFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFiveFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field1").Call(nil)
	}
}

func BenchmarkStandbardTwentyFiveFields(b *testing.B) {
	foo := TwentyFiveFields{}
	for i := 0; i < b.N; i++ {
		foo.Field1()
	}
}

func BenchmarkReflectLastTwentyFiveFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFiveFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field25").Call(nil)
	}
}

func BenchmarkStandbardLastTwentyFiveFields(b *testing.B) {
	foo := TwentyFiveFields{}
	for i := 0; i < b.N; i++ {
		foo.Field25()
	}
}

type ThirtyFields struct {
}

func (tf ThirtyFields) Field1()  {}
func (tf ThirtyFields) Field2()  {}
func (tf ThirtyFields) Field3()  {}
func (tf ThirtyFields) Field4()  {}
func (tf ThirtyFields) Field5()  {}
func (tf ThirtyFields) Field6()  {}
func (tf ThirtyFields) Field7()  {}
func (tf ThirtyFields) Field8()  {}
func (tf ThirtyFields) Field9()  {}
func (tf ThirtyFields) Field10() {}
func (tf ThirtyFields) Field11() {}
func (tf ThirtyFields) Field12() {}
func (tf ThirtyFields) Field13() {}
func (tf ThirtyFields) Field14() {}
func (tf ThirtyFields) Field15() {}
func (tf ThirtyFields) Field16() {}
func (tf ThirtyFields) Field17() {}
func (tf ThirtyFields) Field18() {}
func (tf ThirtyFields) Field19() {}
func (tf ThirtyFields) Field20() {}
func (tf ThirtyFields) Field21() {}
func (tf ThirtyFields) Field22() {}
func (tf ThirtyFields) Field23() {}
func (tf ThirtyFields) Field24() {}
func (tf ThirtyFields) Field25() {}
func (tf ThirtyFields) Field26() {}
func (tf ThirtyFields) Field27() {}
func (tf ThirtyFields) Field28() {}
func (tf ThirtyFields) Field29() {}
func (tf ThirtyFields) Field30() {}

func BenchmarkReflectThirtyFields(b *testing.B) {
	foo := reflect.ValueOf(ThirtyFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field1").Call(nil)
	}
}

func BenchmarkStandbardThirtyFields(b *testing.B) {
	foo := ThirtyFields{}
	for i := 0; i < b.N; i++ {
		foo.Field1()
	}
}

func BenchmarkReflectLastThirtyFields(b *testing.B) {
	foo := reflect.ValueOf(ThirtyFields{})
	for i := 0; i < b.N; i++ {
		foo.MethodByName("Field30").Call(nil)
	}
}

func BenchmarkStandbardLastThirtyFields(b *testing.B) {
	foo := ThirtyFields{}
	for i := 0; i < b.N; i++ {
		foo.Field30()
	}
}

// type Foo struct{}

// func (f *Foo) Method(
// 	arg1 *Bar,
// 	arg2 *Bar,
// 	arg3 *Bar,
// 	arg4 *Bar,
// 	arg5 *Bar,
// 	arg6 *Bar,
// 	arg7 *Bar,
// 	arg8 *Bar,
// 	arg9 *Bar,
// 	arg10 *Bar,
// 	arg11 *Bar,
// 	arg12 *Bar,
// 	arg13 *Bar,
// 	arg14 *Bar,
// 	arg15 *Bar,
// 	arg16 *Bar,
// 	arg17 *Bar,
// 	arg18 *Bar,
// 	arg19 *Bar,
// 	arg20 *Bar,
// 	arg21 *Bar,
// 	arg22 *Bar,
// 	arg23 *Bar,
// 	arg24 *Bar,
// 	arg25 *Bar,
// 	arg26 *Bar,
// 	arg27 *Bar,
// 	arg28 *Bar,
// 	arg29 *Bar,
// 	arg30 *Bar,
// 	arg31 *Bar,
// 	arg32 *Bar,
// 	arg33 *Bar,
// 	arg34 *Bar,
// 	arg35 *Bar,
// 	arg36 *Bar,
// 	arg37 *Bar,
// 	arg38 *Bar,
// 	arg39 *Bar,
// 	arg40 *Bar,
// ) {
// }

// type Bar struct {
// 	// Val1  string
// 	// Val2  string
// 	// Val3  string
// 	// Val4  string
// 	// Val5  string
// 	// Val6  string
// 	// Val7  string
// 	// Val8  string
// 	// Val9  string
// 	// Val10 string
// 	// Val11 string
// 	// Val12 string
// 	// Val13 string
// 	// Val14 string
// 	// Val15 string
// 	// Bar1  *Bar
// 	// Bar2  *Bar
// 	// Bar3  *Bar
// 	// Bar4  *Bar
// 	// Bar5  *Bar
// 	// Bar6  *Bar
// 	// Bar7  *Bar
// 	// Bar8  *Bar
// 	// Bar9  *Bar
// 	// Bar10 *Bar
// 	// Bar11 *Bar
// 	// Bar12 *Bar
// 	// Bar13 *Bar
// 	// Bar14 *Bar
// }

// func BenchmarkReflection(b *testing.B) {
// 	foo := reflect.ValueOf(&Foo{})
// 	args := []reflect.Value{
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		reflect.ValueOf(&Bar{}),
// 		// reflect.ValueOf("11"),
// 		// reflect.ValueOf("12"),
// 		// reflect.ValueOf("13"),
// 	}
// 	for i := 0; i < b.N; i++ {
// 		foo.MethodByName("Method").Call(args)
// 	}
// }

// func BenchmarkStandard(b *testing.B) {
// 	foo := &Foo{}
// 	for i := 0; i < b.N; i++ {
// 		foo.Method(
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			&Bar{},
// 			// "11",
// 			// "12",
// 			// "13",
// 		)
// 	}
// }
