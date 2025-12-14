# Nx & Nexa  
Neural Execution Assistant for Go

Nx & Nexa adalah tooling CLI untuk Go yang menyediakan pengalaman developer seperti **Laravel Artisan** atau **Django manage.py**, namun tetap mengikuti **filosofi Go (simple, explicit, binary-first)**.

---

## Konsep

Proyek ini terdiri dari **dua CLI terpisah** dengan peran yang jelas:

| Tool | Fungsi |
|----|----|
| **nx** | Global project generator & manager |
| **nexa** | Project-level CLI (artisan untuk project) |

**Aturan utama:**
- `nx` digunakan secara global
- `nexa` digunakan di dalam project

---

## Prasyarat

- Go **1.22 atau lebih baru**
- `$GOBIN` atau `$GOPATH/bin` sudah ada di `PATH`

Cek dengan:
```bash
go version
```

---

## Instalasi Nx (Global)

Install Nx menggunakan Go:

```bash
go install github.com/brian2zz/nexa-tools/cmd/nx@latest
```

Verifikasi instalasi:

```bash
nx --version
```

---

## Cek Environment

Untuk memastikan environment siap:

```bash
nx doctor
```

Perintah ini akan mengecek:
- Go terinstall
- Versi Go
- GOPATH / GOBIN
- PATH
- OS dan arsitektur

---

## Membuat Project Baru

Buat project baru dengan Nx:

```bash
nx create-project my-app
```

Struktur project yang dihasilkan:

```
my-app/
├─ cmd/
│  ├─ server/
│  └─ nexa/
│     └─ commands/
├─ app/
├─ go.mod
└─ .env.example
```

---

## Install CLI Project (Nexa)

Masuk ke direktori project:

```bash
cd my-app
```

Install CLI project:

```bash
nx install
```

Setelah install, CLI `nexa` bisa dijalankan langsung.

---

## Menjalankan Project

Jalankan server project:

```bash
nexa serve
```

---

## Update Nx

Untuk memperbarui Nx ke versi terbaru:

```bash
nx update
```

Nx akan mengunduh dan menginstall versi terbaru dari repository.

---

## Filosofi Desain

Nx dan Nexa dipisahkan untuk menjaga:
- Tidak ada konflik binary
- Pemisahan tanggung jawab yang jelas
- Dukungan multi-project
- Maintainability jangka panjang

Nx **tidak menangani logic aplikasi**.  
Semua command terkait aplikasi ada di **Nexa**.

---

## Roadmap

### Nx (Global Tool)
- [x] create-project
- [x] install
- [x] doctor
- [x] update
- [x] version

### Nexa (Project CLI)
- [ ] serve
- [ ] make:model
- [ ] make:controller
- [ ] make:service
- [ ] route:list
- [ ] migration

---

## Kontribusi

Kontribusi sangat terbuka.

1. Fork repository
2. Buat branch baru
3. Commit perubahan
4. Buat Pull Request

---

## Lisensi

MIT License
