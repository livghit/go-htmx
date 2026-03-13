<div align="center">
<kbd>
<img src="https://cdn.dribbble.com/users/722246/screenshots/3977397/media/553d7f83386c55d83cada5cb3f72dc3e.gif" alt="Cool animation" />
</kbd>
<p>Animation by <a href="https://dribbble.com/yupnguyen">Yup Nguyen</a></p>

[![GO](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![HTMX](https://img.shields.io/badge/htmx-white.svg?style=for-the-badge&logo=htmx&logoColor=black)](https://htmx.org/)
[![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)](https://tailwindcss.com/)
[![Alpine.js](https://img.shields.io/badge/alpinejs-white.svg?style=for-the-badge&logo=alpinedotjs&logoColor=%238BC0D0)](https://alpinejs.dev/)

</div>

# go-htmx

> **One way** — not *the* way — to build full-stack Go apps with HTMX + Alpine.js.

A batteries-included starter framework that wires together [Chi](https://github.com/go-chi/chi), [Templ](https://templ.guide/), [HTMX 2](https://htmx.org/), and [Alpine.js](https://alpinejs.dev/) into a coherent project structure. SQLite out of the box, Postgres/MySQL with one env-var change.

> **Note:** Chi was chosen out of personal interest, not because it is the only option. The standard library alone can take you very far — here is a [great talk](https://www.youtube.com/watch?v=Qi9A6-xoOkA) showing exactly that.

---

## Features

| Package | What it does |
|---|---|
| `htmx/` | Inspect every HTMX request header; set every response header (`HX-Redirect`, `HX-Trigger`, `HX-Reswap`…) |
| `flash/` | Stateless HMAC-signed cookie flash messages — no session store needed, survive HTMX redirects |
| `validation/` | Explicit form validation helpers (`Required`, `MinLen`, `IsEmail`, `Matches`, `Unique`) |
| `middleware/` | JWT stored in an HttpOnly cookie — `Protect` and `Optional` route guards, `UserFromCtx` helper |
| `db/` | `database/sql` connection + embedded SQL migration runner (auto-runs on startup, no CLI needed) |
| `db/models/` | User CRUD: `GetByID`, `GetByUsername`, `GetByEmail`, `Create`, `Update`, `Delete` |
| `services/` | Auth service (register / login / logout with bcrypt + JWT) and a base web-handler helper |
| `views/` | Templ layouts and pages wired to Tailwind CDN — swap in a build step when you're ready |

---

## Project layout

```
cmd/
  main.go          ← entry point: connect DB, run migrations, start server
  seed/main.go     ← optional database seeder
config/            ← env-based config loader
db/
  migrations/      ← SQL migrations (auto-applied on startup)
  models/          ← database model types + queries
htmx/              ← HTMX request/response helpers
flash/             ← cookie flash message helpers
middleware/        ← JWT auth middleware
routes/            ← route registration
services/          ← business logic
utils/             ← shared utilities
validation/        ← form validation
views/
  layouts/         ← base HTML layout, nav, flash
  pages/           ← home, login, register
static/            ← embedded static assets (HTMX, Alpine, Bootstrap JS)
```

---

## Quick start

```bash
# 1. Clone
git clone https://github.com/livghit/go-htmx && cd go-htmx

# 2. Configure
cp example.env app.env
# Open app.env and set JWT_SECRET and FLASH_SECRET to long random strings:
#   openssl rand -hex 32

# 3. Run (with hot-reload)
make dev

# 4. Or build a binary
make build && ./bin/app
```

The server starts on `:3000`. The SQLite database is created automatically and migrations are applied on startup.

---

## Make targets

| Target | Description |
|---|---|
| `make dev` | Hot-reload via [air](https://github.com/air-verse/air) |
| `make run` | Run without hot-reload |
| `make build` | Compile production binary to `bin/app` |
| `make templ` | Regenerate Go code from `.templ` files |
| `make migrate-up` | Apply pending migrations via goose CLI |
| `make migrate-down` | Roll back the latest migration |
| `make seed` | Seed the database |
| `make lint` | Run golangci-lint |
| `make test` | Run all tests |
| `make clean` | Remove `bin/` and `app.db` |

---

## Configuration

Copy `example.env` to `app.env` (never commit `app.env`).

```env
APP_NAME="MyApp"
PORT=":3000"

DB_ENGINE=sqlite          # sqlite | postgres | mysql
DB_NAME=app.db

JWT_SECRET=change-me      # openssl rand -hex 32
JWT_EXPIRY=24             # hours

FLASH_SECRET=change-me    # openssl rand -hex 32
```

To use Postgres, set:
```env
DB_ENGINE=postgres
DB_CONN=postgres://user:pass@localhost:5432/myapp?sslmode=disable
```

---

## Contributing

Open a pull request or an issue — see [Contributing.md](Contributing.md) for details.
