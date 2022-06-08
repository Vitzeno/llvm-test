define i32 @main() {
0:
	%a = alloca i32
	%b = alloca i32
	store i32 32, i32* %a
	store i32 16, i32* %b
	%1 = load i32, i32* %a
	%2 = load i32, i32* %b
	%3 = add i32 %1, %2
	ret i32 %3
}
