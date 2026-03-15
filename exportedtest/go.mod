module github.com/kotaoue/go-exportedtest

go 1.14

replace github.com/kotaoue/go-exportedtest/packages/small => ./packages/small

replace github.com/kotaoue/go-exportedtest/packages/big => ./packages/big

require (
	github.com/kotaoue/go-exportedtest/packages/big v0.0.0-00010101000000-000000000000 // indirect
	github.com/kotaoue/go-exportedtest/packages/small v0.0.0-00010101000000-000000000000 // indirect
)
