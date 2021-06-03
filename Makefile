test-unit:
	go test -v ./app/*_unit_test.go

test-unit-default:
	go clean -cache
	go test -v ./app/*_default_test.go

test-unit-table:
	go clean -cache
	go test -v ./app/*_table_test.go

test-unit-table-parallel:
	go clean -cache
	go test -v ./app/*_parallel_test.go

test-all: test-unit-default test-unit-table