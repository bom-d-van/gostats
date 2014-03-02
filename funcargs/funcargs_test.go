package types

import "testing"

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

func NoByteArg()                                                                   {}
func OneByteArg(arg byte)                                                          {}
func FiveByteArgs(arg1, arg2, arg3, arg4, arg5 byte)                               {}
func TenByteArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 byte) {}
func FifteenByteArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 byte) {
}
func TwentyByteArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 byte) {
}

func BenchmarkNoByteArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoByteArg()
	}
}

func BenchmarkOneByteArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneByteArg(0)
	}
}

func BenchmarkFiveByteArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveByteArgs(0, 0, 0, 0, 0)
	}
}

func BenchmarkTenByteArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenByteArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	}
}

func BenchmarkFifteenByteArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenByteArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	}
}

func BenchmarkTwentyByteArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyByteArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	}
}

func NoStringArg()                                                                     {}
func OneStringArg(arg string)                                                          {}
func FiveStringArgs(arg1, arg2, arg3, arg4, arg5 string)                               {}
func TenStringArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 string) {}
func FifteenStringArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 string) {
}
func TwentyStringArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 string) {
}

func BenchmarkNoStringArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoStringArg()
	}
}

func BenchmarkOneStringArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneStringArg("0")
	}
}

func BenchmarkFiveStringArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveStringArgs("0", "0", "0", "0", "0")
	}
}

func BenchmarkTenStringArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenStringArgs("0", "0", "0", "0", "0", "0", "0", "0", "0", "0")
	}
}

func BenchmarkFifteenStringArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenStringArgs("0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0")
	}
}

func BenchmarkTwentyStringArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyStringArgs("0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0")
	}
}

func NoFloat32Arg()                                                                      {}
func OneFloat32Arg(arg float32)                                                          {}
func FiveFloat32Args(arg1, arg2, arg3, arg4, arg5 float32)                               {}
func TenFloat32Args(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 float32) {}
func FifteenFloat32Args(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 float32) {
}
func TwentyFloat32Args(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 float32) {
}

func BenchmarkNoFloat32Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoFloat32Arg()
	}
}

func BenchmarkOneFloat32Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneFloat32Arg(0.0)
	}
}

func BenchmarkFiveFloat32Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveFloat32Args(0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

func BenchmarkTenFloat32Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenFloat32Args(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

func BenchmarkFifteenFloat32Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenFloat32Args(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

func BenchmarkTwentyFloat32Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyFloat32Args(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

func NoFloat64Arg()                                                                      {}
func OneFloat64Arg(arg float64)                                                          {}
func FiveFloat64Args(arg1, arg2, arg3, arg4, arg5 float64)                               {}
func TenFloat64Args(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 float64) {}
func FifteenFloat64Args(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 float64) {
}
func TwentyFloat64Args(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 float64) {
}

func BenchmarkNoFloat64Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoFloat64Arg()
	}
}

func BenchmarkOneFloat64Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneFloat64Arg(0.0)
	}
}

func BenchmarkFiveFloat64Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveFloat64Args(0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

func BenchmarkTenFloat64Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenFloat64Args(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

func BenchmarkFifteenFloat64Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenFloat64Args(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

func BenchmarkTwentyFloat64Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyFloat64Args(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	}
}

type IntArray []int

func NoIntArrayArg()                                                                       {}
func OneIntArrayArg(arg IntArray)                                                          {}
func FiveIntArrayArgs(arg1, arg2, arg3, arg4, arg5 IntArray)                               {}
func TenIntArrayArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 IntArray) {}
func FifteenIntArrayArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 IntArray) {
}
func TwentyIntArrayArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 IntArray) {
}

func BenchmarkNoIntArrayArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoIntArrayArg()
	}
}

func BenchmarkOneIntArrayArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneIntArrayArg(IntArray{})
	}
}

func BenchmarkFiveIntArrayArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveIntArrayArgs(IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{})
	}
}

func BenchmarkTenIntArrayArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenIntArrayArgs(IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{})
	}
}

func BenchmarkFifteenIntArrayArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenIntArrayArgs(IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{})
	}
}

func BenchmarkTwentyIntArrayArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyIntArrayArgs(IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{}, IntArray{})
	}
}

type IntMap map[int]int

func NoIntMapArg()                                                                     {}
func OneIntMapArg(arg IntMap)                                                          {}
func FiveIntMapArgs(arg1, arg2, arg3, arg4, arg5 IntMap)                               {}
func TenIntMapArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 IntMap) {}
func FifteenIntMapArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 IntMap) {
}
func TwentyIntMapArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 IntMap) {
}

func BenchmarkNoIntMapArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoIntMapArg()
	}
}

func BenchmarkOneIntMapArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneIntMapArg(IntMap{})
	}
}

func BenchmarkFiveIntMapArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveIntMapArgs(IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{})
	}
}

func BenchmarkTenIntMapArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenIntMapArgs(IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{})
	}
}

func BenchmarkFifteenIntMapArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenIntMapArgs(IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{})
	}
}

func BenchmarkTwentyIntMapArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyIntMapArgs(IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{}, IntMap{})
	}
}

type EmptyStruct struct{}

func NoEmptyStructArg()                                                                          {}
func OneEmptyStructArg(arg EmptyStruct)                                                          {}
func FiveEmptyStructArgs(arg1, arg2, arg3, arg4, arg5 EmptyStruct)                               {}
func TenEmptyStructArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 EmptyStruct) {}
func FifteenEmptyStructArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 EmptyStruct) {
}
func TwentyEmptyStructArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 EmptyStruct) {
}

func BenchmarkNoEmptyStructArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoEmptyStructArg()
	}
}

func BenchmarkOneEmptyStructArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneEmptyStructArg(EmptyStruct{})
	}
}

func BenchmarkFiveEmptyStructArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveEmptyStructArgs(EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{})
	}
}

func BenchmarkTenEmptyStructArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenEmptyStructArgs(EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{})
	}
}

func BenchmarkFifteenEmptyStructArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenEmptyStructArgs(EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{})
	}
}

func BenchmarkTwentyEmptyStructArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyEmptyStructArgs(EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{}, EmptyStruct{})
	}
}

type OneField struct {
	Field1 int
}

func NoOneFieldArg()                                                                       {}
func OneOneFieldArg(arg OneField)                                                          {}
func FiveOneFieldArgs(arg1, arg2, arg3, arg4, arg5 OneField)                               {}
func TenOneFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 OneField) {}
func FifteenOneFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 OneField) {
}
func TwentyOneFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 OneField) {
}

func BenchmarkNoOneFieldArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoOneFieldArg()
	}
}

func BenchmarkOneOneFieldArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneOneFieldArg(OneField{})
	}
}

func BenchmarkFiveOneFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveOneFieldArgs(OneField{}, OneField{}, OneField{}, OneField{}, OneField{})
	}
}

func BenchmarkTenOneFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenOneFieldArgs(OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{})
	}
}

func BenchmarkFifteenOneFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenOneFieldArgs(OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{})
	}
}

func BenchmarkTwentyOneFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyOneFieldArgs(OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{}, OneField{})
	}
}

type FiveField struct {
	Field1 int
	Field2 int
	Field3 int
	Field4 int
	Field5 int
}

func NoFiveFieldArg()                                                                        {}
func OneFiveFieldArg(arg FiveField)                                                          {}
func FiveFiveFieldArgs(arg1, arg2, arg3, arg4, arg5 FiveField)                               {}
func TenFiveFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 FiveField) {}
func FifteenFiveFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 FiveField) {
}
func TwentyFiveFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 FiveField) {
}

func BenchmarkNoFiveFieldArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoFiveFieldArg()
	}
}

func BenchmarkOneFiveFieldArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneFiveFieldArg(FiveField{})
	}
}

func BenchmarkFiveFiveFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveFiveFieldArgs(FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{})
	}
}

func BenchmarkTenFiveFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenFiveFieldArgs(FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{})
	}
}

func BenchmarkFifteenFiveFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenFiveFieldArgs(FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{})
	}
}

func BenchmarkTwentyFiveFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyFiveFieldArgs(FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{}, FiveField{})
	}
}

type TenField struct {
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

func NoTenFieldArg()                                                                       {}
func OneTenFieldArg(arg TenField)                                                          {}
func FiveTenFieldArgs(arg1, arg2, arg3, arg4, arg5 TenField)                               {}
func TenTenFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 TenField) {}
func FifteenTenFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 TenField) {
}
func TwentyTenFieldArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 TenField) {
}

func BenchmarkNoTenFieldArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoTenFieldArg()
	}
}

func BenchmarkOneTenFieldArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OneTenFieldArg(TenField{})
	}
}

func BenchmarkFiveTenFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiveTenFieldArgs(TenField{}, TenField{}, TenField{}, TenField{}, TenField{})
	}
}

func BenchmarkTenTenFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenTenFieldArgs(TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{})
	}
}

func BenchmarkFifteenTenFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FifteenTenFieldArgs(TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{})
	}
}

func BenchmarkTwentyTenFieldArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwentyTenFieldArgs(TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{}, TenField{})
	}
}

// type Bar struct{}

// func NoArg()                                                                  {}
// func OneArg(arg Bar)                                                          {}
// func FiveArgs(arg1, arg2, arg3, arg4, arg5 Bar)                               {}
// func TenArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 Bar) {}
// func FifteenArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 Bar) {
// }
// func TwentyArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20 Bar) {
// }

// func BenchmarkNoArg(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		NoArg()
// 	}
// }

// func BenchmarkOneArg(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		// OneArg(0)
// 		OneArg(Bar{})
// 	}
// }

// func BenchmarkFiveArgs(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		// FiveArgs(0, 0, 0, 0, 0)
// 		FiveArgs(Bar{}, Bar{}, Bar{}, Bar{}, Bar{})
// 	}
// }

// func BenchmarkTenArgs(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		// TenArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
// 		TenArgs(Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{})
// 	}
// }

// func BenchmarkFifteenArgs(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		// FifteenArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
// 		FifteenArgs(Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{})
// 	}
// }

// func BenchmarkTwentyArgs(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		// TwentyArgs(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
// 		TwentyArgs(Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{}, Bar{})
// 	}
// }
