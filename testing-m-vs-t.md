# testing.M vs testing.T in Go

## Quick Summary

| Type | Purpose | Scope | When to Use |
|------|---------|-------|-------------|
| `testing.M` | **Main test controller** | Entire test package | Setup/teardown for ALL tests |
| `testing.T` | **Individual test controller** | Single test function | Each individual test |

---

## testing.T - Individual Test Control

### What is it?

`testing.T` is passed to **each individual test function**. It controls that specific test and provides methods to report failures, log information, and control test execution.

### Usage

```go
func TestSomething(t *testing.T) {
    // t is used to control THIS test only
    result := add(2, 2)

    if result != 4 {
        t.Errorf("Expected 4, got %d", result)  // Mark test as failed
    }

    t.Log("Test passed!")  // Log information
}
```

### Common `testing.T` Methods

```go
// Reporting failures
t.Error("test failed")       // Mark as failed, continue test
t.Errorf("got %d, want %d", got, want)  // Formatted error
t.Fail()                     // Mark as failed, continue test
t.FailNow()                  // Mark as failed, stop test immediately
t.Fatal("critical error")    // Same as Error + FailNow
t.Fatalf("got %d", got)      // Formatted Fatal

// Logging
t.Log("debug info")          // Only shown if test fails or -v flag
t.Logf("value: %d", value)   // Formatted log

// Control
t.Skip("skipping test")      // Skip this test
t.Skipf("skipping: %s", reason)
t.Parallel()                 // Run this test in parallel with others

// Cleanup
t.Cleanup(func() {
    // Cleanup code runs after test finishes
    closeConnection()
})
```

### Example

```go
func TestDatabase_Insert(t *testing.T) {
    // Setup for this test only
    db := setupTestDB()

    // Cleanup for this test only
    t.Cleanup(func() {
        db.Close()
    })

    // Test logic
    err := db.Insert("test")
    if err != nil {
        t.Fatalf("Insert failed: %v", err)
    }

    t.Log("Insert succeeded")
}
```

---

## testing.M - Package-Level Test Control

### What is it?

`testing.M` is used in `TestMain` to control the **entire test suite** for a package. It runs **once before all tests** and handles setup/teardown for the whole package.

### Usage

```go
func TestMain(m *testing.M) {
    // Setup - runs ONCE before ALL tests
    setupGlobalResources()

    // Run all tests
    exitCode := m.Run()

    // Teardown - runs ONCE after ALL tests
    cleanupGlobalResources()

    // Exit with same code as tests
    os.Exit(exitCode)
}
```

### Key Points

1. **TestMain is optional** - Only create it if you need package-level setup/teardown
2. **Only ONE TestMain per package** - Go allows only one
3. **Must call m.Run()** - This actually runs the tests
4. **Must call os.Exit()** - This ensures proper exit code
5. **Runs before any Test functions** - Perfect for expensive setup

### Common `testing.M` Methods

```go
// Only one method!
exitCode := m.Run()  // Runs all tests, returns exit code (0 = pass, 1 = fail)
```

---

## Real-World Example from Your Codebase

### setup_test.go

```go
package main

import (
	"go-breeders/models"
	"log"
	"os"
	"testing"
)

var testApp application  // Shared across ALL tests

// testing.M - Package-level setup
func TestMain(m *testing.M) {
	// Setup ONCE for ALL tests
	dsn := "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?..."
	db, err := initMySQLDB(dsn)
	if err != nil {
		log.Panic(err)
	}

	// Initialize shared test application
	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	// Run ALL tests in the package
	exitCode := m.Run()

	// Teardown ONCE after ALL tests
	// (Could close db connection here)
	// db.Close()

	// Exit with the same code as the tests
	os.Exit(exitCode)
}
```

### handlers_test.go

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// testing.T - Individual test
func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	// Use the shared testApp initialized in TestMain

	// Setup for THIS test only
	req, err := http.NewRequest("GET", "/api/dog-breeds", nil)
	if err != nil {
		t.Fatal(err)  // t.Fatal - fail this test and stop
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	// Execute test
	handler.ServeHTTP(rr, req)

	// Assertions for THIS test
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status: got %v want %v", status, http.StatusOK)
	}

	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("wrong content type: got %v want %v", ct, "application/json")
	}

	t.Log("Test passed!")  // t.Log - log for this test
}

// Another test using the same testApp
func TestApplication_ShowHome(t *testing.T) {
	// Uses same testApp - no need to recreate it!
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.ShowHome)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status: got %v", status)
	}
}
```

---

## Execution Flow

```
1. TestMain runs
   ├─ Setup code (database connection, etc.)
   ├─ m.Run() executes:
   │   ├─ TestApplication_GetAllDogBreedsJSON(t)
   │   │   ├─ Test-specific setup
   │   │   ├─ Test assertions with t.Error, t.Log, etc.
   │   │   └─ Test-specific cleanup (t.Cleanup)
   │   │
   │   ├─ TestApplication_ShowHome(t)
   │   │   ├─ Test-specific setup
   │   │   ├─ Test assertions
   │   │   └─ Test-specific cleanup
   │   │
   │   └─ ... (more tests)
   │
   ├─ Teardown code (close connections, etc.)
   └─ os.Exit(exitCode)
```

---

## When to Use Each

### Use `testing.M` (TestMain) when:

1. **Expensive setup** that should run once
   - Database connection
   - Loading large test fixtures
   - Starting test servers

2. **Global cleanup** needed after all tests
   - Closing database connections
   - Cleaning up temp directories
   - Stopping servers

3. **Configuration** needed for all tests
   - Environment variables
   - Feature flags
   - Test database migrations

### Use `testing.T` for:

1. **Every individual test** - it's required!
2. **Test-specific setup/teardown** - use `t.Cleanup()`
3. **Reporting test results** - `t.Error()`, `t.Fatal()`
4. **Logging test information** - `t.Log()`
5. **Controlling test execution** - `t.Skip()`, `t.Parallel()`

---

## Benefits of Using TestMain (testing.M)

### Without TestMain (Inefficient)

```go
func TestGetAllDogs(t *testing.T) {
	db := connectToDB()  // Slow!
	defer db.Close()
	// test...
}

func TestGetOneDog(t *testing.T) {
	db := connectToDB()  // Slow again!
	defer db.Close()
	// test...
}

// Database connection created multiple times!
```

### With TestMain (Efficient)

```go
var testDB *sql.DB  // Shared

func TestMain(m *testing.M) {
	testDB = connectToDB()  // Once!
	exitCode := m.Run()
	testDB.Close()
	os.Exit(exitCode)
}

func TestGetAllDogs(t *testing.T) {
	// Use shared testDB
	// Fast!
}

func TestGetOneDog(t *testing.T) {
	// Use shared testDB
	// Fast!
}

// Database connection created only once!
```

---

## Common Patterns

### Pattern 1: Database Tests

```go
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Setup
	testDB = setupTestDatabase()
	runMigrations(testDB)

	// Run tests
	exitCode := m.Run()

	// Cleanup
	testDB.Close()
	os.Exit(exitCode)
}
```

### Pattern 2: Test Server

```go
var testServer *httptest.Server

func TestMain(m *testing.M) {
	// Start server once
	testServer = httptest.NewServer(handler)

	exitCode := m.Run()

	// Stop server
	testServer.Close()
	os.Exit(exitCode)
}
```

### Pattern 3: Configuration

```go
func TestMain(m *testing.M) {
	// Set environment for all tests
	os.Setenv("ENV", "test")
	os.Setenv("LOG_LEVEL", "debug")

	exitCode := m.Run()

	// Cleanup
	os.Unsetenv("ENV")
	os.Exit(exitCode)
}
```

---

## Key Differences Summary

| Feature | testing.M | testing.T |
|---------|-----------|-----------|
| **Function signature** | `TestMain(m *testing.M)` | `TestXxx(t *testing.T)` |
| **Runs** | Once per package | Once per test function |
| **Purpose** | Setup/teardown for ALL tests | Control individual test |
| **Methods** | Only `m.Run()` | Many: Error, Log, Fatal, Skip, etc. |
| **Required** | Optional | Required for every test |
| **Number allowed** | Only one per package | Unlimited |
| **When runs** | Before any tests | For each Test function |
| **Use for** | Expensive global setup | Test assertions & control |

---

## Important Notes

1. **TestMain must call m.Run()** or tests won't execute
2. **TestMain must call os.Exit()** with the exit code from m.Run()
3. **Order matters**: Setup → m.Run() → Teardown → os.Exit()
4. **No t.Parallel() in TestMain** - it controls all tests
5. **TestMain doesn't count as a test** - it's infrastructure

---

## Your Specific Use Case

In your `go-breeders` project:

**testing.M** (setup_test.go):
- Connects to MySQL **once**
- Initializes `testApp` with database and models
- All tests share this connection (fast!)

**testing.T** (handlers_test.go):
- Each test uses the shared `testApp`
- Each test makes HTTP requests and assertions
- Each test can fail independently with `t.Error()` or `t.Fatal()`

This pattern is **perfect** for web applications with databases!
