
Test Subcase
```cmd
go test [path_to_test_file] -run="[method]/[case_name]" -v
```

Test Coverage Command
```cmd
go test ./[package] -cover
```

Test benchmark
```cmd
go test [path_to_test_file] -bench=. -benchmem
```

Run Document
```cmd
godoc -http=:8000
```

Unit Test VS Code Configuration
```
"go.coverOnSave": true,
"go.coverOnSingleTest": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",        
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
}
```

Integration Test
- for hiding the integration test (run only unit test for db purpose)
```cmd
//go:build integration

package ...
```
- for running the integration test
```cmd
go test [path_to_package] -v -tags=integration
```