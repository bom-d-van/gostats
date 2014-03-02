package reflect

import (
	"reflect"
	"testing"
)

type OneField struct {
	Field int
}

func BenchmarkReflectOneField(b *testing.B) {
	foo := reflect.ValueOf(OneField{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field").String()
	}
}

func BenchmarkStandbardOneField(b *testing.B) {
	foo := OneField{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field
	}
}

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

type TenFields struct {
	Field1  int
	Field2  int
	Field3  int
	Field4  int
	Field5  int
	Field6  int
	Field7  int
	Field8  int
	Field9  int
	Field10 int
}

func BenchmarkReflectTenFields(b *testing.B) {
	foo := reflect.ValueOf(TenFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field1").String()
	}
}

func BenchmarkStandbardTenFields(b *testing.B) {
	foo := TenFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field1
	}
}

func BenchmarkReflectLastTenFields(b *testing.B) {
	foo := reflect.ValueOf(TenFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field10").String()
	}
}

func BenchmarkStandbardLastTenFields(b *testing.B) {
	foo := TenFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field10
	}
}

type FifteenFields struct {
	Field1  int
	Field2  int
	Field3  int
	Field4  int
	Field5  int
	Field6  int
	Field7  int
	Field8  int
	Field9  int
	Field10 int
	Field11 int
	Field12 int
	Field13 int
	Field14 int
	Field15 int
}

func BenchmarkReflectFifteenFields(b *testing.B) {
	foo := reflect.ValueOf(FifteenFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field1").String()
	}
}

func BenchmarkStandbardFifteenFields(b *testing.B) {
	foo := FifteenFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field1
	}
}

func BenchmarkReflectLastFifteenFields(b *testing.B) {
	foo := reflect.ValueOf(FifteenFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field15").String()
	}
}

func BenchmarkStandbardLastFifteenFields(b *testing.B) {
	foo := FifteenFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field15
	}
}

type TwentyFields struct {
	Field1  int
	Field2  int
	Field3  int
	Field4  int
	Field5  int
	Field6  int
	Field7  int
	Field8  int
	Field9  int
	Field10 int
	Field11 int
	Field12 int
	Field13 int
	Field14 int
	Field15 int
	Field16 int
	Field17 int
	Field18 int
	Field19 int
	Field20 int
}

func BenchmarkReflectTwentyFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field1").String()
	}
}

func BenchmarkStandbardTwentyFields(b *testing.B) {
	foo := TwentyFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field1
	}
}

func BenchmarkReflectLastTwentyFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field20").String()
	}
}

func BenchmarkStandbardLastTwentyFields(b *testing.B) {
	foo := TwentyFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field20
	}
}

type TwentyFiveFields struct {
	Field1  int
	Field2  int
	Field3  int
	Field4  int
	Field5  int
	Field6  int
	Field7  int
	Field8  int
	Field9  int
	Field10 int
	Field11 int
	Field12 int
	Field13 int
	Field14 int
	Field15 int
	Field16 int
	Field17 int
	Field18 int
	Field19 int
	Field20 int
	Field21 int
	Field22 int
	Field23 int
	Field24 int
	Field25 int
}

func BenchmarkReflectTwentyFiveFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFiveFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field1").String()
	}
}

func BenchmarkStandbardTwentyFiveFields(b *testing.B) {
	foo := TwentyFiveFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field1
	}
}

func BenchmarkReflectLastTwentyFiveFields(b *testing.B) {
	foo := reflect.ValueOf(TwentyFiveFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field25").String()
	}
}

func BenchmarkStandbardLastTwentyFiveFields(b *testing.B) {
	foo := TwentyFiveFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field25
	}
}

type ThirtyFields struct {
	Field1  int
	Field2  int
	Field3  int
	Field4  int
	Field5  int
	Field6  int
	Field7  int
	Field8  int
	Field9  int
	Field10 int
	Field11 int
	Field12 int
	Field13 int
	Field14 int
	Field15 int
	Field16 int
	Field17 int
	Field18 int
	Field19 int
	Field20 int
	Field21 int
	Field22 int
	Field23 int
	Field24 int
	Field25 int
	Field26 int
	Field27 int
	Field28 int
	Field29 int
	Field30 int
}

func BenchmarkReflectThirtyFields(b *testing.B) {
	foo := reflect.ValueOf(ThirtyFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field1").String()
	}
}

func BenchmarkStandbardThirtyFields(b *testing.B) {
	foo := ThirtyFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field1
	}
}

func BenchmarkReflectLastThirtyFields(b *testing.B) {
	foo := reflect.ValueOf(ThirtyFields{})
	for i := 0; i < b.N; i++ {
		_ = foo.FieldByName("Field30").String()
	}
}

func BenchmarkStandbardLastThirtyFields(b *testing.B) {
	foo := ThirtyFields{}
	for i := 0; i < b.N; i++ {
		_ = foo.Field30
	}
}

// type Foo struct {
// 	Val        Bar
// 	FloatVal   float64
// 	Val1       Bar
// 	FloatVal1  float64
// 	Val2       Bar
// 	FloatVal2  float64
// 	Val3       Bar
// 	FloatVal3  float64
// 	Val4       Bar
// 	FloatVal4  float64
// 	Val5       Bar
// 	FloatVal5  float64
// 	Val6       Bar
// 	FloatVal6  float64
// 	Val7       Bar
// 	FloatVal7  float64
// 	Val8       Bar
// 	FloatVal8  float64
// 	Val9       Bar
// 	FloatVal9  float64
// 	Val10      Bar
// 	FloatVal10 float64
// 	Val11      Bar
// 	FloatVal11 float64
// 	Val12      Bar
// 	FloatVal12 float64
// 	Val13      Bar
// 	FloatVal13 float64
// 	Val14      Bar
// 	FloatVal14 float64
// 	Val15      Bar
// 	FloatVal15 float64
// 	Val16      Bar
// 	FloatVal16 float64
// 	Val17      Bar
// 	FloatVal17 float64
// 	Val18      Bar
// 	FloatVal18 float64
// 	Val19      Bar
// 	FloatVal19 float64
// 	Val20      Bar
// 	FloatVal20 float64
// 	Val21      Bar
// 	FloatVal21 float64
// 	Val22      Bar
// 	FloatVal22 float64
// 	Val23      Bar
// 	FloatVal23 float64
// 	Val24      Bar
// 	FloatVal24 float64
// 	Val25      Bar
// 	FloatVal25 float64
// 	Val26      Bar
// 	FloatVal26 float64
// 	Val27      Bar
// 	FloatVal27 float64
// 	Val28      Bar
// 	FloatVal28 float64
// 	Val29      Bar
// 	FloatVal29 float64
// 	Val30      Bar
// 	FloatVal30 float64
// 	Val31      Bar
// 	FloatVal31 float64
// 	Val32      Bar
// 	FloatVal32 float64
// 	Val33      Bar
// 	FloatVal33 float64
// 	Val34      Bar
// 	FloatVal34 float64
// 	Val35      Bar
// 	FloatVal35 float64
// 	Val36      Bar
// 	FloatVal36 float64
// 	Val37      Bar
// 	FloatVal37 float64
// 	Val38      Bar
// 	FloatVal38 float64
// 	Val39      Bar
// 	FloatVal39 float64
// }

// type Bar struct {
// 	Val1  string
// 	Val2  string
// 	Val3  string
// 	Val4  string
// 	Val5  string
// 	Val6  string
// 	Val7  string
// 	Val8  string
// 	Val9  string
// 	Val10 string
// 	Val11 string
// 	Val12 string
// 	Val13 string
// 	Val14 string
// 	Val15 string
// 	Val16 string
// 	Val17 string
// 	Val18 string
// 	Val19 string
// 	Val20 string
// 	Val21 string
// 	Val22 string
// 	Val23 string
// 	Val24 string
// 	Val25 string
// 	Val26 string
// 	Val27 string
// 	Val28 string
// 	Val29 string
// 	Val30 string
// 	Val31 string
// 	Val32 string
// }

// func (f Foo) GetVal()   {}
// func (f Foo) GetVal1()  {}
// func (f Foo) GetVal2()  {}
// func (f Foo) GetVal3()  {}
// func (f Foo) GetVal4()  {}
// func (f Foo) GetVal5()  {}
// func (f Foo) GetVal6()  {}
// func (f Foo) GetVal7()  {}
// func (f Foo) GetVal8()  {}
// func (f Foo) GetVal9()  {}
// func (f Foo) GetVal10() {}
// func (f Foo) GetVal11() {}
// func (f Foo) GetVal12() {}
// func (f Foo) GetVal13() {}
// func (f Foo) GetVal14() {}
// func (f Foo) GetVal15() {}
// func (f Foo) GetVal16() {}
// func (f Foo) GetVal17() {}
// func (f Foo) GetVal18() {}
// func (f Foo) GetVal19() {}
// func (f Foo) GetVal20() {}
// func (f Foo) GetVal21() {}
// func (f Foo) GetVal22() {}
// func (f Foo) GetVal23() {}
// func (f Foo) GetVal24() {}
// func (f Foo) GetVal25() {}
// func (f Foo) GetVal26() {}
// func (f Foo) GetVal27() {}
// func (f Foo) GetVal28() {}
// func (f Foo) GetVal29() {}
// func (f Foo) GetVal30() {}
// func (f Foo) GetVal31() {}
// func (f Foo) GetVal32() {}
// func (f Foo) GetVal33() {}
// func (f Foo) GetVal34() {}
// func (f Foo) GetVal35() {}
// func (f Foo) GetVal36() {}
// func (f Foo) GetVal37() {}
// func (f Foo) GetVal38() {}
// func (f Foo) GetVal39() {}

// func BenchmarkReflect(b *testing.B) {
// 	foo := reflect.ValueOf(Foo{})
// 	for i := 0; i < b.N; i++ {
// 		// _ = foo.FieldByName("Val").String()
// 		// _ = foo.FieldByName("Val1").String()
// 		// _ = foo.FieldByName("Val2").String()
// 		// _ = foo.FieldByName("Val3").String()
// 		// _ = foo.FieldByName("Val4").String()
// 		// _ = foo.FieldByName("Val5").String()
// 		// _ = foo.FieldByName("Val6").String()
// 		// _ = foo.FieldByName("Val7").String()
// 		// _ = foo.FieldByName("Val8").String()
// 		// _ = foo.FieldByName("Val9").String()
// 		// _ = foo.FieldByName("Val10").String()
// 		// _ = foo.FieldByName("Val11").String()
// 		// _ = foo.FieldByName("Val12").String()
// 		// _ = foo.FieldByName("Val13").String()
// 		// _ = foo.FieldByName("Val14").String()
// 		// _ = foo.FieldByName("Val15").String()
// 		// _ = foo.FieldByName("Val16").String()
// 		// _ = foo.FieldByName("Val17").String()
// 		// _ = foo.FieldByName("Val18").String()
// 		_ = foo.FieldByName("Val19").String()
// 		// _ = foo.FieldByName("Val20").String()
// 		// _ = foo.FieldByName("Val21").String()
// 		// _ = foo.FieldByName("Val22").String()
// 		// _ = foo.FieldByName("Val23").String()
// 		// _ = foo.FieldByName("Val24").String()
// 		// _ = foo.FieldByName("Val25").String()
// 		// _ = foo.FieldByName("Val26").String()
// 		// _ = foo.FieldByName("Val27").String()
// 		// _ = foo.FieldByName("Val28").String()
// 		// _ = foo.FieldByName("Val29").String()
// 		// _ = foo.FieldByName("Val30").String()
// 		// _ = foo.FieldByName("Val31").String()
// 		// _ = foo.FieldByName("Val32").String()
// 		// _ = foo.FieldByName("Val33").String()
// 		// _ = foo.FieldByName("Val34").String()
// 		// _ = foo.FieldByName("Val35").String()
// 		// _ = foo.FieldByName("Val36").String()
// 		// _ = foo.FieldByName("Val37").String()
// 		// _ = foo.FieldByName("Val38").String()
// 		// _ = foo.FieldByName("Val39").String()
// 		// foo.MethodByName("GetVal19").Call(nil)
// 		// foo.MethodByName("GetVal1").Call(nil)
// 		// foo.MethodByName("GetVal2").Call(nil)
// 		// foo.MethodByName("GetVal3").Call(nil)
// 		// foo.MethodByName("GetVal4").Call(nil)
// 		// foo.MethodByName("GetVal5").Call(nil)
// 		// foo.MethodByName("GetVal6").Call(nil)
// 		// foo.MethodByName("GetVal7").Call(nil)
// 		// foo.MethodByName("GetVal8").Call(nil)
// 		// foo.MethodByName("GetVal9").Call(nil)
// 	}
// }

// func BenchmarkStandbard(b *testing.B) {
// 	foo := Foo{}
// 	for i := 0; i < b.N; i++ {
// 		// _ = foo.Val
// 		// _ = foo.Val1
// 		// _ = foo.Val2
// 		// _ = foo.Val3
// 		// _ = foo.Val4
// 		// _ = foo.Val5
// 		// _ = foo.Val6
// 		// _ = foo.Val7
// 		// _ = foo.Val8
// 		// _ = foo.Val9
// 		// _ = foo.Val10
// 		// _ = foo.Val11
// 		// _ = foo.Val12
// 		// _ = foo.Val13
// 		// _ = foo.Val14
// 		// _ = foo.Val15
// 		// _ = foo.Val16
// 		// _ = foo.Val17
// 		// _ = foo.Val18
// 		_ = foo.Val19
// 		// _ = foo.Val20
// 		// _ = foo.Val21
// 		// _ = foo.Val22
// 		// _ = foo.Val23
// 		// _ = foo.Val24
// 		// _ = foo.Val25
// 		// _ = foo.Val26
// 		// _ = foo.Val27
// 		// _ = foo.Val28
// 		// _ = foo.Val29
// 		// _ = foo.Val30
// 		// _ = foo.Val31
// 		// _ = foo.Val32
// 		// _ = foo.Val33
// 		// _ = foo.Val34
// 		// _ = foo.Val35
// 		// _ = foo.Val36
// 		// _ = foo.Val37
// 		// _ = foo.Val38
// 		// _ = foo.Val39
// 		// foo.GetVal19()
// 		// foo.GetVal1()
// 		// foo.GetVal2()
// 		// foo.GetVal3()
// 		// foo.GetVal4()
// 		// foo.GetVal5()
// 		// foo.GetVal6()
// 		// foo.GetVal7()
// 		// foo.GetVal8()
// 		// foo.GetVal9()
// 	}
// }
