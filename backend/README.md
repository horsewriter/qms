# Quality System Backend - General Information Module

Backend en Go con plantillas Templ para el módulo de General Information.

## Requisitos

- Go 1.21 o superior
- Templ CLI para generar plantillas
- PostgreSQL (Supabase)

## Instalación

1. Instalar dependencias:
```bash
cd backend
go mod download
```

2. Instalar Templ CLI:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

3. Generar templates:
```bash
templ generate
```

## Configuración

El archivo `.env` ya está configurado con las credenciales de Supabase. Asegúrate de actualizar la conexión a la base de datos en `internal/database/database.go` con la contraseña correcta.

## Ejecución

```bash
cd cmd/server
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

## Estructura del Proyecto

```
backend/
├── cmd/
│   └── server/
│       └── main.go              # Punto de entrada
├── internal/
│   ├── database/
│   │   └── database.go          # Conexión a DB
│   ├── handlers/                # Handlers HTTP
│   │   ├── menu.go
│   │   ├── employee.go
│   │   ├── area.go
│   │   ├── level.go
│   │   ├── part_number.go
│   │   └── calibration_institution.go
│   └── models/
│       └── models.go            # Modelos de datos
└── templates/                    # Templates Templ
    ├── menu.templ
    ├── employee.templ
    ├── area.templ
    ├── level.templ
    ├── part_number.templ
    └── calibration_institution.templ
```

## Endpoints

### General Information Module

- `GET /general-information/menu` - Menú principal
- `GET /general-information/employee` - Formulario de empleados
- `GET /general-information/employee/list` - Lista de empleados
- `POST /general-information/employee` - Crear empleado
- `DELETE /general-information/employee/{id}` - Eliminar empleado

- `GET /general-information/area` - Formulario de áreas
- `GET /general-information/area/list` - Lista de áreas
- `POST /general-information/area` - Crear área
- `DELETE /general-information/area/{id}` - Eliminar área

- `GET /general-information/level` - Formulario de niveles
- `GET /general-information/level/list` - Lista de niveles
- `POST /general-information/level` - Crear nivel
- `DELETE /general-information/level/{id}` - Eliminar nivel

- `GET /general-information/part-number` - Formulario de números de parte
- `GET /general-information/part-number/list` - Lista de números de parte
- `POST /general-information/part-number` - Crear número de parte
- `DELETE /general-information/part-number/{id}` - Eliminar número de parte

- `GET /general-information/calibration-institution` - Formulario de instituciones de calibración
- `GET /general-information/calibration-institution/list` - Lista de instituciones
- `POST /general-information/calibration-institution` - Crear institución
- `DELETE /general-information/calibration-institution/{id}` - Eliminar institución

## Base de Datos

Las tablas ya fueron creadas en Supabase:
- employees
- areas
- levels
- part_numbers
- calibration_institutions

Todas las tablas incluyen RLS (Row Level Security) habilitado.
