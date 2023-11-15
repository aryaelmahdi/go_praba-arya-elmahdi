# Middleware

1. Middleware adalah perangkat lunak yang berjalan di antara server web (atau aplikasi web) dan endpoint. Middleware digunakan untuk melakukan berbagai tugas di antara permintaan HTTP yang diterima oleh server dan respons yang dikirimkan kembali ke klien

2. Ada banyak contoh middleware pihak ketiga seperti Negroni, Echo, Interpose, Alice, dan dapat dibuat sendiri. Echo menyediakan middleware seperti JWT, Basic Auth, Proxy, Redirect, CORS, Logger dll.

3. Echo menyediakan #Pre() dan #Use() yang dimana #Pre() akan di eksekusi sebelum router memproses request, sementara #Use() akan dieksekusi setelah router memproses request dan memiliki akses penuh ke echo.Context API