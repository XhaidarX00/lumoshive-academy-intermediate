# Voucher System API

Voucher System API adalah sebuah REST API yang dikembangkan menggunakan Golang dengan framework **Gin** dan database **PostgreSQL**. API ini menyediakan berbagai endpoint untuk mengelola voucher, termasuk pembuatan, pengeditan, validasi, redeem, dan pelacakan riwayat voucher.

## Teknologi yang Digunakan

- **Bahasa Pemrograman**: Golang
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Logger**: Zap logger
- **JSON**: Format input/output data
- **Testing**: Unit testing dengan target rata-rata coverage 50%

## Fitur Utama

1. **Pengelolaan Voucher**:

   - Membuat voucher baru.
   - Menghapus voucher.
   - Mengedit detail voucher.
   - Melihat daftar voucher dengan filter.
   - Menampilkan daftar voucher yang dapat diredeem berdasarkan poin pengguna.

2. **Redeem Voucher**:

   - Melakukan redeem voucher berdasarkan poin pengguna.
   - Validasi poin pengguna sebelum redeem.

3. **Validasi Voucher**:

   - Mengecek validitas voucher berdasarkan syarat dan ketentuan.
   - Menampilkan informasi manfaat voucher (diskon, gratis ongkir, dll.).

4. **Penggunaan Voucher**:

   - Menggunakan voucher pada transaksi.
   - Menyimpan riwayat penggunaan voucher.

5. **Riwayat dan Pelacakan**:

   - Melihat riwayat redeem voucher oleh pengguna.
   - Melihat riwayat penggunaan voucher oleh pengguna.
   - Melihat daftar pengguna yang menggunakan voucher tertentu.

6. **Automigrasi dan Seeder**:
   - Mendukung automigrasi database.
   - Seeder data untuk testing dan pengembangan.

## Endpoint API

### Voucher Management

| Method | Endpoint           | Deskripsi                                         |
| ------ | ------------------ | ------------------------------------------------- |
| POST   | `/vouchers/create` | Membuat voucher baru.                             |
| DELETE | `/vouchers/:id`    | Menghapus voucher berdasarkan `voucher_id`.       |
| PUT    | `/vouchers/:id`    | Mengedit detail voucher berdasarkan `voucher_id`. |
| GET    | `/vouchers/`       | Melihat daftar semua voucher dengan filter.       |

### Redeem Points

| Method | Endpoint                  | Deskripsi                                           |
| ------ | ------------------------- | --------------------------------------------------- |
| GET    | `/vouchers/redeem-points` | Melihat daftar voucher yang bisa diredeem.          |
| POST   | `/vouchers/redeem`        | Melakukan redeem voucher berdasarkan poin pengguna. |

### Voucher Validation and Usage

| Method | Endpoint                      | Deskripsi                                          |
| ------ | ----------------------------- | -------------------------------------------------- |
| GET    | `/vouchers/:user_id`          | Menampilkan daftar voucher yang dimiliki pengguna. |
| GET    | `/vouchers/:user_id/validate` | Mengecek validitas voucher.                        |
| POST   | `/vouchers/`                  | Menggunakan voucher pada transaksi.                |

### History

| Method | Endpoint                                   | Deskripsi                                           |
| ------ | ------------------------------------------ | --------------------------------------------------- |
| GET    | `/vouchers/redeem-history/:user_id`        | Melihat riwayat redeem voucher pengguna.            |
| GET    | `/vouchers/usage-history/:user_id`         | Melihat riwayat penggunaan voucher pengguna.        |
| GET    | `/vouchers/users-by-voucher/:voucher_code` | Melihat pengguna yang menggunakan voucher tertentu. |

## Instalasi dan Penggunaan

### Prasyarat

1. **Golang** versi 1.19 atau lebih baru.
2. **PostgreSQL** terinstal di sistem Anda.

### Langkah Instalasi

1. **Clone Repository**:

   ```bash
   git clone https://github.com/Nameless-ID/project-sistem-voucher-golang-homework-team4
   cd voucher-system
   ```

2. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

3. **Konfigurasi Database**:

   - Ubah file `.env` untuk menyesuaikan konfigurasi database Anda:
     `env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_NAME=db_voucher
    DB_ConnectTimeOut=10
    App_Debug=true
     `

4. **Jalankan Server**:

   ```bash
   go run main.go
   ```

   Server akan berjalan di `http://localhost:8080`.

## Testing

Jalankan unit test dengan perintah berikut:

```bash
go test ./... -cover
```

## Catatan Tambahan

- Pastikan semua input tervalidasi menggunakan `binding tag` untuk menghindari kesalahan data.
- Gunakan logger untuk mencatat semua operasi penting dalam aplikasi.

## Kontribusi

Kami menerima kontribusi dari siapa saja. Silakan buat **pull request** atau ajukan **issue** di repository ini.
