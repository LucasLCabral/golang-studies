# Basic Test
go test .
go test -v

# Cover
go test -coverprofile=cover.out
go test -coverprofile=cover.out
go tool cover -html=cover.out

# Bench
go test -bench=.
go test -bench=. -run=^#

# Fuzz
go test -fuzz=. -run=^#
