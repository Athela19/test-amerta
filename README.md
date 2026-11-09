# Amerta Starter — React + Go (net/http) + PostgreSQL

Aplikasi full-stack untuk login dengan username, manajemen pengguna, dan
manajemen departemen. Karyawan tersimpan di database dan tampil di dalam
departemen; belum ada layar atau API untuk mengubah karyawan.

| Bagian | Teknologi | Port saat development |
| ------ | --------- | --------------------- |
| Frontend | React 18 + Vite + React Router. Di Docker dilayani Nginx. | Docker `3000`, lokal `5173` |
| Backend | Go 1.26, `net/http` (tanpa framework), `pgx/v5` | `8080` |
| Database | PostgreSQL 16 | `5432` |

Routing, middleware, JSON, cookie, dan graceful shutdown memakai standard
library Go.

Dokumentasi model dan alur:

- [Entity Relationship Diagram](docs/erd.md) — kamus tabel dan kardinalitas

### Use case

![Diagram use case: login, logout, kelola departemen, dan kelola pengguna](docs/use_case.png)

### Activity

![Diagram activity: alur login dan kelola departemen](docs/activity.png)

## Struktur folder

```
.
├── docker-compose.yml
├── docs/
│   ├── erd.md                  # model data
│   ├── use_case.png            # diagram use case
│   └── activity.png            # diagram activity
├── db/
│   └── init/001_init.sql       # schema + seed
├── backend/
│   ├── cmd/api/main.go
│   ├── internal/
│   │   ├── auth/               # bcrypt, JWT, refresh token
│   │   ├── config/
│   │   ├── database/
│   │   ├── httpapi/            # router, middleware, handler
│   │   └── store/              # query users, refresh_tokens, departments, employees
│   ├── .env                    # dipakai Compose; host database = postgres
│   ├── .env.example            # referensi nilai untuk shell lokal (tidak di-load otomatis)
│   └── Dockerfile
└── frontend/
    ├── public/                 # letakkan logo resmi sebagai logo.png
    ├── src/
    │   ├── components/         # AppLayout, BrandLogo, Modal
    │   ├── pages/              # LoginPage, DepartmentsPage, UsersPage
    │   ├── api.ts
    │   ├── App.tsx
    │   └── styles.css
    ├── nginx.conf              # SPA + proxy /api (hanya image Docker)
    └── Dockerfile
```

Pilih salah satu cara menjalankan di bawah. Keduanya memakai database, user,
dan password yang sama (`amerta` / `amerta` / database `amerta`).

## Menjalankan dengan Docker

Butuh Docker Engine dan Docker Compose. `docker-compose.yml` memakai network
**external** `api_net`. Buat sekali:

```bash
docker network create api_net
```

Lalu dari root repo:

```bash
docker compose up --build
```

Jika Docker hanya terpasang di WSL, jalankan dari PowerShell:

```powershell
wsl --cd "D:\path\ke\test-amerta" docker compose up --build
```

`down`, `down -v`, dan `ps` memakai pola `wsl --cd` yang sama.

Compose membaca `backend/.env` lewat `env_file`. Di file itu `DATABASE_URL`
mengarah ke host `postgres` (nama service), bukan `localhost`. Jangan memakai
file itu untuk proses Go di luar container.

Build frontend Docker menerima arg `VITE_API_URL=/api` agar browser memanggil
API lewat proxy Nginx pada origin yang sama.

Buka:

- Frontend: <http://localhost:3000>
- API health: <http://localhost:8080/api/health>
- PostgreSQL: `localhost:5432` (user, password, dan database: `amerta`)

Hentikan:

```bash
docker compose down          # container berhenti, data tetap
docker compose down -v       # container dan volume database dihapus
```

Schema serta akun demo hanya dijalankan saat volume `pgdata` pertama kali
dibuat (`db/init` dipasang ke `docker-entrypoint-initdb.d`). Kalau volume lama
masih ada dan schema-nya usang, jalankan `docker compose down -v` lalu `up`
lagi.

## Menjalankan tanpa Docker

PostgreSQL, backend, dan frontend berjalan langsung di mesin. Tidak ada
container dan tidak ada network `api_net`.

Yang perlu terpasang:

| Perangkat | Versi |
| --------- | ----- |
| PostgreSQL | 16 (`psql` ada di `PATH`) |
| Go | 1.26 atau lebih baru |
| Node.js | 22 LTS atau lebih baru. Image Docker memakai Node 24. |

### 1. Database

Buat role dan database sekali (sebagai superuser `postgres`). Lewati perintah
yang error karena objek sudah ada:

```bash
psql -U postgres -c "CREATE USER amerta WITH PASSWORD 'amerta' LOGIN;"
psql -U postgres -c "CREATE DATABASE amerta OWNER amerta;"
```

Di Windows, `psql` sering ada di
`C:\Program Files\PostgreSQL\16\bin`. Tambahkan folder itu ke `PATH`, atau
panggil `psql.exe` dengan path penuh.

Lalu terapkan schema dari root repo. Script aman dijalankan ulang.

```bash
psql "postgres://amerta:amerta@localhost:5432/amerta" -f db/init/001_init.sql
```

### 2. Backend

Proses Go **tidak** membaca file `.env`. `config.Load` memakai environment
proses saja. Bila variabel kosong, default-nya sudah mengarah ke
`localhost:5432`, dan `JWT_SECRET` memakai secret development.

Jalankan dari folder `backend`. Jangan mengekspor isi `backend/.env` ke shell:
host di file itu adalah `postgres`, yang hanya ada di jaringan Docker.
`backend/.env.example` hanya referensi nilai untuk disalin ke environment
shell secara manual.

PowerShell:

```powershell
cd backend
$env:DATABASE_URL = "postgres://amerta:amerta@localhost:5432/amerta?sslmode=disable"
$env:JWT_SECRET = "dev-only-jwt-secret-change-me-0123456789abcdef"
$env:CORS_ALLOWED_ORIGIN = "http://localhost:3000,http://localhost:5173"
go run ./cmd/api
```

bash:

```bash
cd backend
DATABASE_URL="postgres://amerta:amerta@localhost:5432/amerta?sslmode=disable" \
JWT_SECRET="dev-only-jwt-secret-change-me-0123456789abcdef" \
CORS_ALLOWED_ORIGIN="http://localhost:3000,http://localhost:5173" \
go run ./cmd/api
```

API siap di <http://localhost:8080/api/health>.

### 3. Frontend

Vite mem-proxy `/api` ke `http://localhost:8080`. Secara default
`VITE_API_URL` tidak perlu di-set (client memakai `/api`).

```bash
cd frontend
npm install
npm run dev
```

Buka <http://localhost:5173>.

## Akun demo

| Role | Username | Password |
| ---- | -------- | -------- |
| admin | `admin` | `admin123` |
| user | `user` | `user123` |

Login memakai **username**, bukan email. Akun ini hanya untuk development.
Hash seed diverifikasi oleh tes `TestSeedPasswords`.

Setelah login, aplikasi membuka **Departemen** (`/departments`). Semua peran
boleh melihat, menambah, mengubah, dan menghapus departemen. Hapus ditolak bila
departemen masih punya karyawan. Menu **Pengguna** (`/users`) hanya untuk
`admin`: tambah pengguna, ganti role, dan hapus. Admin tidak bisa mengubah role
atau menghapus akunnya sendiri.

Logo resmi diletakkan di `frontend/public/logo.png`. Selama file belum ada, UI
memakai wordmark sementara.

## Alur autentikasi

1. `POST /api/auth/login` memverifikasi password dengan **bcrypt**, lalu:
   - membuat **JWT access token** (HS256, default 15 menit) di cookie
     `amerta_access` (`HttpOnly`, `Path=/`, `SameSite=Lax`);
   - membuat **refresh token acak** (default 7 hari) di cookie
     `amerta_refresh` (`HttpOnly`, `Path=/api/auth`). Di tabel
     `refresh_tokens` hanya SHA-256 hash-nya yang disimpan.
2. Frontend memanggil `/api/auth/me` saat dibuka. Jika access token sudah
   kedaluwarsa, client memanggil `/api/auth/refresh` satu kali lalu mengulang
   request.
3. `POST /api/auth/refresh` **merotasi** refresh token: token lama dicabut dan
   token baru dibuat dalam satu transaksi.
4. `POST /api/auth/logout` menghapus kedua cookie. Jika cookie refresh ada,
   token tersebut dicabut di database. Endpoint tetap sukses tanpa cookie.

Role (`admin` / `user`) ada di `users.role`. JWT hanya membawa id pengguna.
Role dibaca ulang dari database pada setiap request, jadi perubahan role
langsung berlaku.

## Endpoint API

| Method | Path | Auth | Keterangan |
| ------ | ---- | ---- | ---------- |
| POST | `/api/auth/login` | publik | Login dengan username, set cookie |
| POST | `/api/auth/refresh` | cookie refresh | Rotasi refresh token, set cookie baru |
| POST | `/api/auth/logout` | publik | Cabut refresh jika ada, hapus cookie |
| GET | `/api/auth/me` | cookie access (JWT) | Profil user yang sedang login |
| GET | `/api/users` | admin | Daftar pengguna |
| POST | `/api/users` | admin | Tambah pengguna |
| PATCH | `/api/users/{id}` | admin | Ubah role pengguna |
| DELETE | `/api/users/{id}` | admin | Hapus pengguna |
| GET | `/api/departments` | cookie access (JWT) | Daftar departemen beserta karyawannya |
| POST | `/api/departments` | cookie access (JWT) | Tambah departemen |
| PATCH | `/api/departments/{id}` | cookie access (JWT) | Ubah nama departemen |
| DELETE | `/api/departments/{id}` | cookie access (JWT) | Hapus departemen (gagal jika masih ada karyawan) |
| GET | `/api/health` | publik | Status API dan koneksi database |

Kolom Auth “admin” berarti cookie access valid **dan** `users.role = 'admin'`.
Cookie access adalah JWT di cookie HttpOnly, bukan header `Authorization`.

Contoh dengan `curl` (API di port 8080):

```bash
curl -i -c cookies.txt \
  -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'

curl -b cookies.txt http://localhost:8080/api/auth/me

curl -b cookies.txt http://localhost:8080/api/users

curl -i -b cookies.txt \
  -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"operator","password":"operator123","role":"user"}'

curl -b cookies.txt http://localhost:8080/api/departments

curl -i -b cookies.txt -c cookies.txt -X POST http://localhost:8080/api/auth/refresh

curl -i -b cookies.txt -c cookies.txt -X POST http://localhost:8080/api/auth/logout
```

## Konfigurasi environment

### Backend

| Variabel | Default | Keterangan |
| -------- | ------- | ---------- |
| `PORT` | `8080` | Port backend |
| `DATABASE_URL` | `postgres://amerta:amerta@localhost:5432/amerta?sslmode=disable` | Di Docker, `backend/.env` menimpa host menjadi `postgres` |
| `CORS_ALLOWED_ORIGIN` | `http://localhost:3000,http://localhost:5173` | Daftar origin, pisahkan dengan koma |
| `JWT_SECRET` | secret development | **Wajib diganti di production** |
| `ACCESS_TOKEN_TTL` | `15m` | Umur access token |
| `REFRESH_TOKEN_TTL` | `168h` | Umur refresh token (7 hari) |
| `COOKIE_SECURE` | `false` | Set `true` jika memakai HTTPS |

### Frontend

| Variabel | Default | Keterangan |
| -------- | ------- | ---------- |
| `VITE_API_URL` | `/api` | Basis URL client. Di Docker di-set saat build image. |

## Catatan keamanan

- Ganti `JWT_SECRET` dan password akun demo sebelum production.
- Set `COOKIE_SECURE=true` saat aplikasi berjalan di HTTPS.
- Access token berumur pendek. Refresh token disimpan sebagai hash dan dirotasi
  setiap kali dipakai.
- Otorisasi admin memeriksa role terbaru di database, sehingga akun yang
  dihapus atau diturunkan role-nya langsung kehilangan akses.
- Semua user terautentikasi dapat mengubah data departemen. Batasi di API jika
  produk membutuhkan hak admin saja.
- `SameSite=Lax` membantu mitigasi CSRF. Untuk kebutuhan lebih ketat, tambahkan
  CSRF token dan rate limiting pada endpoint login.
- `db/init` hanya bootstrap. Perubahan schema berikutnya sebaiknya lewat tool
  migrasi seperti `golang-migrate`.
