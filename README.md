# To Sql

An application that allows for rapid ingest of csv and tabular data and adding it to a sql table.

This is in place of an annoying workflow where data gets sent using excel or csv files and it having to be used for comparing with data that exists in a database.

Additionally it allows for quick ingest of the data to do comparison operations that would be easier to implement using sql or set based methods.

[![Super-Linter](https://github.com/kstolte/toSql/actions/workflows/super-linter.yml/badge.svg)](https://github.com/marketplace/actions/super-linter)

## project structure

root

- assets/* this is the output directory and site HTML
- server/
  - main.go: contains server for serving up the contents of the assets folder.
- wasm/
  - main.go: main package that holds logic for Wasm interface to JavaScript.
  - toSql/
    - testData/*: contains various preselected csv documents for testing
    - main.go: core logic package that gets built t
    - main_test.go: snapshot and benchmark testing for preventing regressions.

to build
`make build-wasm`

to dev
```bash
make dev
```

start server
```bash
cd ./cmd/server
go run main.go
```

