# Cron Expression Parse  

By Michal Gondar, Jan 2021

## Documentation

Slot name    | Allowed values  | Allowed special characters
----------   | --------------  | --------------------------
Minute       | 0-59            | * / , -
Hour         | 0-23            | * / , -
Day of month | 1-31            | * / , - ?
Month        | 1-12 or JAN-DEC | * / , -
Day of week  | 0-7 or SUN-SAT  | * / , - ?

## Possible extensions
- Implement `L` last day for `Day of month` and `Day of week`.
- Implement `W` nearest working day for given day for `Day of month`.

## Setup
- 

## Build

`go build -o bin/cron-parser cmd/cron-parser/main.go`

## Usage

`bin/cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"`

## Tests

`go clean -testcache && go test ./...`

