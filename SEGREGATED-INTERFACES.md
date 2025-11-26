# Interface Segregation - Refactoring Final

## El Problema Anterior

### Interfaz Gigante (Anti-Patrón)
```go
// ❌ MAL: Una interfaz con todo mezclado
type Repository interface {
    // Perros
    AllDogBreeds() ([]*models.DogBreed, error)
    AllDogs() ([]*models.Dog, error)

    // Gatos
    AllCats() ([]*models.Cat, error)

    // Criadores
    AllBreeders() ([]*models.Breeder, error)
}
```

### Problemas:

1. **Acoplamiento forzado:**
```go
// Si solo quieres MongoDB para perros, tienes que implementar TODO
type MongoDBDogRepo struct {
    client *mongo.Client
}

func (m *MongoDBDogRepo) AllCats() ([]*models.Cat, error) {
    panic("no uso gatos!") // ❌ Problema!
}
```

2. **Viola el Interface Segregation Principle (SOLID)**
   - Los clientes no deberían depender de métodos que no usan

3. **Difícil de testear:**
```go
// Tienes que mockear TODO aunque solo uses perros
type MockRepo struct {
    // Tiene que implementar perros, gatos, criadores...
}
```

## La Solución: Interfaces Pequeñas y Separadas

### Nueva Estructura

```
repository/
├── dog_repository.go          # ✓ Solo interfaz de perros
├── cat_repository.go          # ✓ Solo interfaz de gatos
├── breeder_repository.go      # ✓ Solo interfaz de criadores
│
├── mysql/                     # Implementación MySQL
│   ├── dog_mysql.go          # Solo perros en MySQL
│   ├── cat_mysql.go          # Solo gatos en MySQL
│   └── breeder_mysql.go      # Solo criadores en MySQL
│
└── mock/                      # Implementación para tests
    ├── dog_mock.go           # Mock solo de perros
    ├── cat_mock.go           # Mock solo de gatos
    └── breeder_mock.go       # Mock solo de criadores
```

## Interfaces Separadas

### repository/dog_repository.go
```go
package repository

import "go-breeders/models"

// ✓ Interfaz SOLO para perros
type DogRepository interface {
    AllDogBreeds() ([]*models.DogBreed, error)
    GetDogBreedByID(id int) (*models.DogBreed, error)
    AllDogs() ([]*models.Dog, error)
    GetDogByID(id int) (*models.Dog, error)
    InsertDog(dog *models.Dog) (int, error)
    UpdateDog(dog *models.Dog) error
    DeleteDog(id int) error
}
```

### repository/cat_repository.go
```go
package repository

import "go-breeders/models"

// ✓ Interfaz SOLO para gatos
type CatRepository interface {
    AllCatBreeds() ([]*models.CatBreed, error)
    GetCatBreedByID(id int) (*models.CatBreed, error)
    AllCats() ([]*models.Cat, error)
    GetCatByID(id int) (*models.Cat, error)
    InsertCat(cat *models.Cat) (int, error)
    UpdateCat(cat *models.Cat) error
    DeleteCat(id int) error
}
```

### repository/breeder_repository.go
```go
package repository

import "go-breeders/models"

// ✓ Interfaz SOLO para criadores
type BreederRepository interface {
    AllBreeders() ([]*models.Breeder, error)
    GetBreederByID(id int) (*models.Breeder, error)
    InsertBreeder(breeder *models.Breeder) (int, error)
    UpdateBreeder(breeder *models.Breeder) error
    DeleteBreeder(id int) error
}
```

## Implementaciones MySQL

### repository/mysql/dog_mysql.go
```go
package mysql

import (
    "database/sql"
    "go-breeders/repository"
)

// ✓ Implementa SOLO DogRepository
type DogRepo struct {
    DB *sql.DB
}

func NewDogRepo(db *sql.DB) repository.DogRepository {
    return &DogRepo{DB: db}
}

func (r *DogRepo) AllDogBreeds() ([]*models.DogBreed, error) {
    // Consulta MySQL solo para dog_breeds
}
```

**Lo importante:** `DogRepo` NO necesita saber nada de gatos o criadores.

## Implementaciones Mock

### repository/mock/dog_mock.go
```go
package mock

import "go-breeders/repository"

// ✓ Mock SOLO para perros
type DogRepo struct{}

func NewDogRepo() repository.DogRepository {
    return &DogRepo{}
}

func (m *DogRepo) AllDogBreeds() ([]*models.DogBreed, error) {
    return []*models.DogBreed{
        {ID: 1, Breed: "Chihuahua", ...},
        {ID: 2, Breed: "German Shepherd", ...},
    }, nil
}
```

## Application Struct

### Antes (Acoplado)
```go
type application struct {
    Repo repository.Repository  // ❌ Una sola interfaz gigante
}
```

### Ahora (Desacoplado)
```go
type application struct {
    DogRepo     repository.DogRepository      // ✓ Solo perros
    CatRepo     repository.CatRepository      // ✓ Solo gatos
    BreederRepo repository.BreederRepository  // ✓ Solo criadores
}
```

## Uso en Main

### cmd/web/main.go
```go
func main() {
    db, _ := initMySQLDB(dsn)

    app := application{
        // Inyecta cada repositorio por separado
        DogRepo:     mysql.NewDogRepo(db),      // MySQL para perros
        CatRepo:     mysql.NewCatRepo(db),      // MySQL para gatos
        BreederRepo: mysql.NewBreederRepo(db),  // MySQL para criadores
    }

    // ...
}
```

## Uso en Handlers

### cmd/web/handlers.go
```go
func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
    var t toolbox.Tools

    // ✓ Usa el repositorio específico de perros
    dogBreeds, err := app.DogRepo.AllDogBreeds()
    if err != nil {
        _ = t.ErrorJSON(w, err, http.StatusInternalServerError)
        return
    }

    _ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}
```

**Ventaja:** Es obvio que este handler solo usa perros.

## Uso en Tests

### cmd/web/setup_test.go
```go
func TestMain(m *testing.M) {
    testApp = application{
        // ✓ Mock solo lo que necesitas
        DogRepo:     mock.NewDogRepo(),
        CatRepo:     mock.NewCatRepo(),
        BreederRepo: mock.NewBreederRepo(),
    }

    os.Exit(m.Run())
}
```

## Flexibilidad: Mix de Implementaciones

Ahora puedes mezclar implementaciones:

```go
app := application{
    DogRepo:     mysql.NewDogRepo(db),        // MySQL para perros
    CatRepo:     mongo.NewCatRepo(mongoClient), // MongoDB para gatos!
    BreederRepo: api.NewBreederRepo(apiClient), // API externa para criadores!
}
```

**¡Cada entidad puede usar una base de datos diferente!**

## Beneficios SOLID

### 1. Interface Segregation Principle ✓
Cada interfaz tiene solo los métodos que necesita.

### 2. Single Responsibility Principle ✓
Cada repositorio es responsable de UNA entidad.

### 3. Dependency Inversion Principle ✓
Los handlers dependen de interfaces, no de implementaciones concretas.

### 4. Open/Closed Principle ✓
Puedes agregar nuevas implementaciones sin modificar código existente.

## Ejemplo Real: Solo Quieres MongoDB para Perros

```go
package mongo

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go-breeders/repository"
)

// ✓ Solo implementa DogRepository
type DogRepo struct {
    client *mongo.Client
}

func NewDogRepo(client *mongo.Client) repository.DogRepository {
    return &DogRepo{client: client}
}

func (r *DogRepo) AllDogBreeds() ([]*models.DogBreed, error) {
    // MongoDB query solo para perros
    collection := r.client.Database("breeders").Collection("dog_breeds")
    // ...
}
```

**No necesitas implementar gatos ni criadores!**

Luego en tu app:

```go
mongoClient, _ := mongo.Connect(...)
mysqlDB, _ := sql.Open(...)

app := application{
    DogRepo:     mongo.NewDogRepo(mongoClient),   // MongoDB
    CatRepo:     mysql.NewCatRepo(mysqlDB),       // MySQL
    BreederRepo: mysql.NewBreederRepo(mysqlDB),   // MySQL
}
```

## Comparación Final

### Antes (Interfaz Gigante)
```
✗ Una interfaz con TODO
✗ Implementaciones forzadas a tener todo
✗ No puedes mezclar bases de datos
✗ Tests mock TODO aunque solo uses una cosa
```

### Ahora (Interfaces Segregadas)
```
✓ Interfaces pequeñas y enfocadas
✓ Implementa solo lo que necesitas
✓ Mezcla MySQL, MongoDB, APIs, etc.
✓ Tests mock solo lo que usas
✓ Código más claro y mantenible
```

## Convenciones de Nombres en Go

Go no tiene una convención rígida, pero las más comunes son:

### Para Repositorios:
```go
DogRepo        // Simple, claro
DogRepository  // Más explícito
dogRepo        // Variable (minúscula)
```

### Para Interfaces:
```go
DogRepository    // Nombre del contrato
Repository       // Genérico (evitar si es específico)
```

### En este proyecto usamos:
```go
// Interfaces (repository/*.go)
type DogRepository interface { ... }

// Structs de implementación (mysql/*.go, mock/*.go)
type DogRepo struct { ... }

// Variables en application
DogRepo repository.DogRepository
```

## Test Results

```bash
$ go test -v ./cmd/web/

=== RUN   TestApplication_GetAllDogBreedsJSON
    handlers_test.go:42: Response: [
      {"id":1,"breed":"Chihuahua",...},
      {"id":2,"breed":"German Shepherd",...},
      {"id":3,"breed":"Labrador Retriever",...}
    ]
--- PASS: TestApplication_GetAllDogBreedsJSON (0.00s)
PASS ✅
```

## Resumen

Esta refactorización sigue el **Interface Segregation Principle** de SOLID:

> "Los clientes no deberían verse forzados a depender de interfaces que no usan."

**Resultado:**
- Código más flexible
- Fácil de testear
- Fácil de extender
- Cada componente es independiente
- Puedes mezclar tecnologías sin problemas

¡Esto es arquitectura profesional en Go! 🎉
