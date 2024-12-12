# Cara Mengaktifkan SMTP di Gmail

SMTP (Simple Mail Transfer Protocol) memungkinkan Anda mengirim email melalui server Gmail menggunakan aplikasi atau perangkat lain. Berikut langkah-langkahnya:

## 1. Aktifkan Akses untuk Aplikasi yang Kurang Aman (Jika Diperlukan)
1. Login ke akun Gmail Anda.
2. Buka [Pengaturan Akses Akun](https://myaccount.google.com/security).
3. Temukan opsi **"Akses aplikasi yang kurang aman"**.
   - Jika opsi ini tidak muncul, lanjutkan ke langkah berikutnya karena fitur ini sudah tergantikan oleh App Password.

## 2. Gunakan Password Aplikasi
1. Login ke akun Gmail Anda.
2. Buka [Halaman Password Aplikasi](https://myaccount.google.com/apppasswords).
3. Pilih **Aplikasi** dan **Perangkat** yang ingin Anda gunakan.
4. Klik **Hasilkan**, lalu salin password aplikasi yang diberikan.
   - Gunakan password ini sebagai pengganti password Gmail Anda di aplikasi atau perangkat.

## 3. Konfigurasi SMTP
Gunakan pengaturan berikut di aplikasi email Anda:

- **SMTP Server**: `smtp.gmail.com`
- **Port**:
  - Untuk SSL: `465`
  - Untuk TLS/STARTTLS: `587`
- **Authentication Required**: Ya
- **Email Address**: Alamat email Gmail Anda
- **Password**: Password aplikasi yang telah Anda hasilkan.

## 4. Uji Koneksi
1. Masukkan pengaturan SMTP di aplikasi email atau perangkat Anda.
2. Kirim email uji coba untuk memastikan bahwa konfigurasi berfungsi dengan benar.

## 5. Catatan Keamanan
- Gunakan **Password Aplikasi** untuk meningkatkan keamanan.
- Jangan bagikan informasi login atau password Anda.

Jika mengalami masalah, pastikan Anda telah mengaktifkan autentikasi dua langkah (2FA) di akun Gmail Anda.
