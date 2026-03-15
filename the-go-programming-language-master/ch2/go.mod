module github.com/kotaoue/the-go-programming-language/ch2

go 1.14

replace (
	github.com/kotaoue/the-go-programming-language/ch2/popcount => ./popcount
	github.com/kotaoue/the-go-programming-language/ch2/tempconv => ./tempconv
)

require (
	github.com/kotaoue/the-go-programming-language/ch2/popcount v0.0.0-00010101000000-000000000000
	github.com/kotaoue/the-go-programming-language/ch2/tempconv v0.0.0-00010101000000-000000000000
	gopl.io v0.0.0-20200323155855-65c318dde95e
)
