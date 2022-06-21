target triple = "x86_64-apple-darwin"

define i32 @main() {
main:
	%a = alloca double
	store double 23.0, double* %a
	br i1 false, label %0, label %1

0:
	br label %2

1:
	br label %2

2:
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
