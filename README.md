# SOLID Principles dalam Go

SOLID adalah lima prinsip desain dalam pemrograman yang membantu membuat kode lebih mudah dipelihara, fleksibel, dan skalabel. Prinsip ini membantu dalam pembuatan arsitektur perangkat lunak yang lebih baik.

## 1. Single Responsibility Principle (SRP)
Setiap class, struct, atau modul harus memiliki hanya satu alasan untuk berubah, yaitu satu tanggung jawab spesifik. Dengan prinsip ini, kode menjadi lebih modular dan lebih mudah dikelola.

## 2. Open/Closed Principle (OCP)
Kode harus terbuka untuk ekstensi tetapi tertutup untuk modifikasi. Artinya, kita bisa menambahkan fitur baru tanpa harus mengubah kode yang sudah ada, biasanya dengan menggunakan interface atau pola desain yang fleksibel.

## 3. Liskov Substitution Principle (LSP)
Subkelas atau implementasi dari suatu interface harus bisa menggantikan superclass atau interface tersebut tanpa mengubah perilaku program. Dengan kata lain, tipe turunan harus bisa digunakan sebagai pengganti tipe induknya tanpa menyebabkan kesalahan.

## 4. Interface Segregation Principle (ISP)
Jangan buat interface yang terlalu besar dan memaksa implementasi mendukung metode yang tidak diperlukan. Sebaiknya, pecah interface menjadi beberapa bagian kecil yang lebih spesifik agar implementasi hanya perlu menggunakan metode yang relevan.

## 5. Dependency Inversion Principle (DIP)
Kode harus bergantung pada abstraksi (interface) daripada implementasi konkret. Dengan cara ini, dependensi dalam sistem menjadi lebih fleksibel dan mudah untuk diuji serta diganti jika diperlukan.

Dengan menerapkan prinsip SOLID dalam Go, kode akan lebih bersih, terstruktur, dan lebih mudah diperluas di masa depan. 🚀
