target triple = "x86_64-pc-linux-gnu"

define i32 @main() {
entry:
	%a = alloca double
	store double 5.0, double* %a
	%b = alloca double
	store double 5.0, double* %b
	store double 10.0, double* %a
	%0 = fcmp ole double 15.0, 15.0
	br i1 %0, label %if.then, label %if.else

if.then:
	%msg = alloca i8
	%1 = call i32 (i8*, ...) @printf(i8* %msg, i32 10)
	br label %if.then

if.else:
	%msg = alloca i8
	%2 = call i32 (i8*, ...) @printf(i8* %msg, i32 25)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
