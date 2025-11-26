# Basic Go Patterns

This document explains the fundamental patterns used in this Go web application and how they compare to traditional Object-Oriented Programming (OOP).

## The Receiver-Based Methods Pattern

### What It Is

In Go, we define a struct and then attach methods to it using receivers. These methods can be spread across multiple files within the same package, organized by concern.

**Example from this project:**

```go
// main.go - Define the type
type application struct {
    templateMap map[string]*template.Template
    config      appConfig
}

// handlers.go - HTTP handlers
func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
    app.render(w, "home.page.gohtml", nil)
}

// routes.go - Route configuration
func (app *application) routes() http.Handler {
    mux := chi.NewRouter()
    // ...
    return mux
}

// render.go - Template rendering
func (app *application) render(w http.ResponseWriter, t string, td *templateData) {
    // ...
}
```

### Alternative Names for This Pattern

- **Receiver-based methods** (most literal)
- **Struct with methods** (descriptive)
- **Method sets on types** (formal Go terminology)
- **Application struct pattern** (when used for web apps)
- **App container pattern** (informal)

## Go vs Traditional OOP

### Java Equivalent (Single Class)

```java
public class Application {
    private Map<String, Template> templateMap;
    private AppConfig config;

    // All methods in one file
    public void showHome(HttpServletRequest req, HttpServletResponse res) {
        render("home.page.html", null);
    }

    public Handler routes() {
        // ...
    }

    private void render(String template, TemplateData data) {
        // ...
    }
}
```

### Java Alternative (Composition)

```java
public class Application {
    private final HandlerService handlers;
    private final RouteService routes;
    private final RenderService renderer;

    public Application(HandlerService handlers,
                      RouteService routes,
                      RenderService renderer) {
        this.handlers = handlers;
        this.routes = routes;
        this.renderer = renderer;
    }

    // Delegates to composed services
}
```

## Advantages Over Traditional OOP

### 1. Simplicity & Explicitness

**No hidden behavior** - You know exactly where methods are defined. No need to navigate inheritance hierarchies.

```go
// Go - explicit, clear
func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
    app.render(w, "home.page.gohtml", nil)
}
```

vs

```java
// Java - where is render() defined? Parent class? Interface?
public class HomeController extends BaseController {
    @Override
    public void showHome() {
        super.render("home.page.html");
    }
}
```

### 2. No Inheritance Complexity

Go avoids common OOP problems:
- **Diamond problem** - Multiple inheritance conflicts
- **Fragile base class problem** - Changes to parent break children
- **Deep inheritance hierarchies** - Hard to understand and maintain
- **Gorilla/banana problem** - "You wanted a banana but got the gorilla holding it and the entire jungle"

### 3. Composition is Explicit

Dependencies are visible at the struct level:

```go
type application struct {
    db       *Database      // Clear what you depend on
    logger   Logger         // Interface = testable
    cache    *Cache
    config   appConfig
}
```

**Benefit:** No hunting through constructors and parent classes to understand dependencies.

### 4. Interface Satisfaction is Implicit

Types automatically satisfy interfaces without declaring it:

```go
type Renderer interface {
    render(w http.ResponseWriter, t string)
}

// application implements Renderer automatically
// No "implements" keyword needed
```

**Benefit:**
- Types can satisfy interfaces they don't know about
- Great for testing (create test doubles easily)
- Better decoupling

### 5. Small, Focused Interfaces

Go encourages small, single-method interfaces:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Compose them
type ReadWriter interface {
    Reader
    Writer
}
```

vs Java/C# with large interface contracts.

**Benefit:** Easier to implement, test, and mock.

### 6. Package-Level Encapsulation

Visibility is controlled by capitalization:

```go
// Public - exported (uppercase first letter)
func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {...}

// Private - package-only (lowercase first letter)
func (app *application) buildTemplateFromDisk(t string) {...}
```

**Benefit:** Encapsulation at the package boundary, not class level. More practical and flexible.

### 7. No Constructor Overloading Confusion

One clear initialization pattern:

```go
app := &application{
    templateMap: make(map[string]*template.Template),
    config:      cfg,
}
```

**Benefit:** No choosing between multiple constructor variants.

## When to Use Composition in Go

While our current pattern works well, you should consider composition when:

### 1. Different Concerns Need Different State

```go
type application struct {
    handlers *HandlerService  // Has its own state
    renderer *RenderService   // Has its own state
    db       *Database         // Shared resource
}
```

### 2. You Want to Swap Implementations

```go
type application struct {
    renderer Renderer  // Interface - can swap implementations
    logger   Logger    // Interface - can use different loggers
}
```

### 3. Components Are Reusable

When components can be used across multiple applications:

```go
// Reusable logger package
type Logger struct {
    output io.Writer
    level  LogLevel
}

// Use in multiple apps
app1 := &WebApp{logger: logger}
app2 := &APIApp{logger: logger}
```

### 4. Your Struct Gets Too Large

Rule of thumb: If your struct has >10-15 fields, consider breaking it up.

## Disadvantages Compared to OOP

To be fair, Go's approach has trade-offs:

1. **No polymorphic constructors** - Can't have factory hierarchies like OOP
2. **Less familiar** - Developers from Java/C#/C++ need adjustment period
3. **No method overloading** - Must use different method names
4. **Verbosity** - Receiver syntax can feel repetitive

## Real-World Example: Our Application

Our simple approach:

```go
type application struct {
    templateMap map[string]*template.Template
    config      appConfig
}
```

If this were traditional Java OOP, we might have:
- `AbstractApplication` base class
- `WebApplication extends AbstractApplication`
- `TemplateManager` separate class
- `ConfigurationService` separate class
- `TemplateFactory` for creating templates
- Various interfaces (`IApplication`, `ITemplateManager`, etc.)
- Dependency injection framework configuration

In Go: **One struct, methods organized by concern across files.**

## The Go Philosophy

> "Simplicity is complicated but the clarity is worth it."
> — Rob Pike (Go co-creator)

Go trades the power and flexibility of OOP inheritance for:
- **Simplicity** - Easier to understand
- **Explicitness** - No hidden behavior
- **Composition** - Build from small, focused pieces

For most applications, this is a net win.

## References

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)
- Dave Cheney's blog on [Practical Go](https://dave.cheney.net/practical-go)
