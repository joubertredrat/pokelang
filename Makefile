test-unit:
	go clean -cache
	go test -v ./app/*_default_test.go

test-unit-table:
	go clean -cache
	go test -v ./app/*_table_test.go

test-unit-table-parallel:
	go clean -cache
	go test -v ./app/*_parallel_test.go

test-all: test-unit test-unit-table