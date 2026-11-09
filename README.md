# Amerta Starter — React + Go (net/http) + PostgreSQL

Starter full-stack yang mengikuti `docker-compose.yml` di root project, dengan
fitur pertama berupa **login** memakai bcrypt, role `admin`/`user`, JWT access
token, dan refresh token yang disimpan di cookie.

| Service    | Teknologi                                    | Port host |
| ---------- | -------------------------------------------- | --------- |
| `frontend` | React 18 + Vite + Nginx                      | `3000`    |
| `backend`  | Go murni dengan `net/http` (tanpa framework) | `8080`    |
| `postgres` | PostgreSQL 16                                | `5432`    |

Routing, middleware, JSON, cookie, dan graceful shutdown memakai standard
library Go. `pgx/v5` hanya dipakai sebagai driver PostgreSQL.

## Struktur folder

```
.
├── docker-compose.yml
├── db/
│   └── init/001_init.sql       # tabel users + refresh_tokens, seed akun demo
├── backend/
│   ├── cmd/api/main.go         # entrypoint
│   ├── internal/
│   │   ├── auth/               # bcrypt, JWT, refresh token
│   │   ├── config/             # konfigurasi env
│   │   ├── database/           # koneksi PostgreSQL
│   │   ├── httpapi/            # router, middleware, handler auth
│   │   └── store/              # query users & refresh_tokens
│   ├── .env                    # environment untuk container
│   └── Dockerfile
└── frontend/
    ├── public/                 # letakkan logo resmi sebagai logo.png
    ├── src/
    │   ├── components/         # BrandLogo
    │   ├── pages/              # LoginPage, HomePage
    │   ├── api.ts              # client auth + auto refresh
    │   └── App.tsx
    ├── nginx.conf              # serve SPA + proxy /api ke backend
    └── Dockerfile
```

## Menjalankan dengan Docker

`docker-compose.yml` memakai network **external** bernama `api_net`, jadi buat
dulu sekali saja:

```bash
docker network create api_net
```

Lalu jalankan:

```bash
docker compose up --build
```

Buka:

- Frontend: <http://localhost:3000>
- API health: <http://localhost:8080/api/health>
- PostgreSQL: `localhost:5432` (user/password/db: `amerta`)

Hentikan:

```bash
docker compose down          # container berhenti, data tetap
docker compose down -v       # container + volume database dihapus
```

> **Penting:** schema dan akun demo hanya dibuat saat volume `pgdata` pertama
> kali dibuat. Kalau sebelumnya sudah pernah menjalankan versi dengan tabel
> `items`, jalankan `docker compose down -v` agar schema auth dibuat ulang.

## Akun demo

| Role  | Email               | Password   |
| ----- | ------------------- | ---------- |
| admin | `admin@amerta.local` | `Admin123!` |
| user  | `user@amerta.local`  | `User123!`  |

Akun ini hanya untuk development. Ganti atau hapus sebelum deploy.

Logo resmi diletakkan di `frontend/public/logo.png`. Selama file belum ada, UI
memakai wordmark sementara. Setelah login, halaman Beranda ringkas ditampilkan;
halaman departemen dan karyawan akan ditambahkan pada fitur berikutnya.

## Alur autentikasi

1. `POST /api/auth/login` memverifikasi password dengan **bcrypt**, lalu:
   - membuat **JWT access token** (HS256, default 15 menit) yang disimpan di
     cookie `amerta_access` (`HttpOnly`, `Path=/`, `SameSite=Lax`);
   - membuat **refresh token acak** (default 7 hari) yang disimpan di cookie
     `amerta_refresh` (`HttpOnly`, `Path=/api/auth`). Di database hanya
     SHA-256 hash-nya yang disimpan.
2. Frontend memanggil `/api/auth/me` saat dibuka. Jika access token sudah
   kedaluwarsa, client otomatis memanggil `/api/auth/refresh` satu kali lalu
   mengulang request.
3. `POST /api/auth/refresh` **merotasi** refresh token: token lama dicabut dan
   token baru dibuat dalam satu transaksi.
4. `POST /api/auth/logout` mencabut refresh token dan menghapus kedua cookie.

Role (`admin`/`user`) disimpan di kolom `users.role` dan di claim `role` pada
JWT. Halaman Beranda menampilkan area berbeda berdasarkan role.

## Endpoint API

| Method | Path                | Auth   | Keterangan                          |
| ------ | ------------------- | ------ | ----------------------------------- |
| POST   | `/api/auth/login`   | publik | Login, set cookie access + refresh  |
| POST   | `/api/auth/refresh` | cookie | Rotasi refresh token, set cookie baru |
| POST   | `/api/auth/logout`  | cookie | Cabut refresh token, hapus cookie   |
| GET    | `/api/auth/me`      | JWT    | Profil user yang sedang login       |
| GET    | `/api/health`       | publik | Status API + koneksi database       |

Contoh dengan `curl`:

```bash
# login (simpan cookie)
curl -i -c cookies.txt \
  -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@amerta.local","password":"Admin123!"}'

# profil
curl -b cookies.txt http://localhost:8080/api/auth/me

# refresh (cookie lama diputar)
curl -i -b cookies.txt -c cookies.txt -X POST http://localhost:8080/api/auth/refresh

# logout
curl -i -b cookies.txt -c cookies.txt -X POST http://localhost:8080/api/auth/logout
```

## Konfigurasi environment

`backend/.env` dipakai oleh container (`env_file` di compose):

| Variabel             | Default                       | Keterangan                              |
| -------------------- | ----------------------------- | --------------------------------------- |
| `PORT`               | `8080`                        | Port backend                            |
| `DATABASE_URL`       | `postgres://amerta:...`       | Koneksi PostgreSQL                      |
| `CORS_ALLOWED_ORIGIN`| `http://localhost:3000,...`   | Daftar origin, pisahkan dengan koma     |
| `JWT_SECRET`         | secret development            | **Wajib diganti di production**         |
| `ACCESS_TOKEN_TTL`   | `15m`                         | Umur access token                       |
| `REFRESH_TOKEN_TTL`  | `168h`                        | Umur refresh token (7 hari)             |
| `COOKIE_SECURE`      | `false`                       | Set `true` jika memakai HTTPS           |

## Development tanpa Docker (opsional)

1. Jalankan PostgreSQL saja:

   ```bash
   docker compose up -d postgres
   ```

2. Backend (butuh Go 1.26+). `backend/.env` berisi host `postgres` untuk
   container, jadi untuk lokal set `DATABASE_URL` secara eksplisit:

   ```powershell
   cd backend
   $env:DATABASE_URL = "postgres://amerta:amerta@localhost:5432/amerta?sslmode=disable"
   $env:JWT_SECRET = "dev-only-jwt-secret-change-me-0123456789abcdef"
   go run ./cmd/api
   ```

   ```bash
   # bash / zsh
   cd backend
   DATABASE_URL="postgres://amerta:amerta@localhost:5432/amerta?sslmode=disable" \
   JWT_SECRET="dev-only-jwt-secret-change-me-0123456789abcdef" \
   go run ./cmd/api
   ```

3. Frontend (butuh Node 22.12+). Vite mem-proxy `/api` ke `http://localhost:8080`:

   ```bash
   cd frontend
   npm install
   npm run dev
   ```

   Buka <http://localhost:5173>.

## Catatan keamanan

- Ganti `JWT_SECRET` dan password akun demo sebelum production.
- Set `COOKIE_SECURE=true` saat aplikasi berjalan di HTTPS.
- Access token berumur pendek; refresh token disimpan sebagai hash dan
  dirotasi setiap kali dipakai.
- `SameSite=Lax` membantu mitigasi CSRF. Untuk kebutuhan lebih ketat,
  tambahkan CSRF token dan rate limiting pada endpoint login.
- Schema di `db/init` hanya untuk bootstrap. Untuk perubahan schema
  selanjutnya, gunakan tool migrasi seperti `golang-migrate`.
