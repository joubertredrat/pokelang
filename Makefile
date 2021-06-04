.PHONY: default
default: run ;

run:
	go run main.go

test-unit:
	go test -v ./app/*_unit_test.go

test-unit-default:
	go clean -cache
	go test -v ./app/*_default_test.go

test-unit-default-specific:
	go clean -cache
	go test -v ./app/*_default_test.go -run ^TestStorageWith100Items

test-unit-table:
	go clean -cache
	go test -v ./app/*_table_test.go

test-unit-table-specific:
	go clean -cache
	go test -v ./app/*_table_test.go -run .*/Test_storage_with_100_items
	go test -v ./app/*_table_test.go -run ^TestStorageTable/Test_storage_with_100_items

test-unit-table-parallel:
	go clean -cache
	go test -v ./app/*_parallel_test.go

test-all: test-unit-default test-unit-table