target triple = "x86_64-pc-linux-gnu"

define i32 @main() {
entry:
	%a = alloca double
	store double 50.0, double* %a
	%msg = alloca i8
	%0 = call i32 (i8*, ...) @printf(i8* %msg, i32 0)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
