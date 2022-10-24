<h1 align="center">Assignment 16 - ORM and Code Structure</h1>
<h2 align="center">Resume Materi</h2>
<hr>
<ul>
    <li>Pengertian ORM (Keuntungan dan Kekurangan)</li>
        <p>ORM adalah teknik pemrograman untuk mengubah data antara tipe sistem yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek</p>
        <p>Keuntungan ORM</p>
        <p>- Query yang duplikat atau berulang semakin berkurang</p>
        <p>- Secara otomatis mengambil data ke objek</p>
        <p>- Cara sederhana jika Anda ingin menyaring data sebelum menyimpannya di database</p>
        <p>- Beberapa memiliki feature cache query</p>
        <br>
        <p>Kekurangan ORM</p>
        <p>- Lapisan dalam kode bertambah dan biayai proses overhead</p>
        <p>- Memuat data relationship yang tidak perlu</p>
        <p>- Query yang kompleks dapat lama ditulis dengan ORM (> 10 tabel bergabung)</p>
        <p>- Function SQL tertentu yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada function khusus</p>
    <li>Database Migration</li>
        <p>Database Migration adalah Cara memperbarui versi basis data agar sesuai dengan perubahan versi aplikasi serta perubahan dapat ditingkatkan ke versi terbaru atau pengembalian ke versi sebelumnya</p>
        <p>Mengapa Database Migration ? </p>
        <p>- Kesederhanaan update basis data</p>
        <p>- Kesederhanaan rollback basis data</p>
        <p>- Melacak perubahan pada struktur database</p>
        <p>- Riwayat struktur basis data ditulis pada kode</p>
        <p>- Selalu kompatibel dengan perubahan versi aplikasi</p>
    <li>Code Structure menggunakan MVC</li>
        <p>MVC adalah kependekan dari Model, View, Controller, dalam MVC setiap bagian dari model, view dan controller memiliki kode dengan tujuan tertentu, dan tujuan tersebut berbeda</p>
        <p>Mengapa perlu struktur ?</p>
        <p>Untuk mengarsipkan aplikasi dalam bagian yang berbeda beda, terapkan pemisahan, lebih sedikit konflik pada pembuatan versi serta mempermudah dalam menangani konflik atau error yang terjadi</p>
</ul>
<br>

<h2>API CRUD and Code Structuring</h2>
<h3>Problem 1 = API Crud Using Database</h3>
<h4>Code</h4>
<p align="center">
    <img src="screenshots/problem1_1.png">
    <br>
    <img src="screenshots/problem1_2.png">
    <br>
    <img src="screenshots/problem1_3.png">
    <br>
    <img src="screenshots/problem1_4.png">
    <br>
    <img src="screenshots/problem1_5.png">
    <br>
</p>
<h4>Output</h4>
<p>POST /users (Create New Users)</p>
<p align="center">
    <img src="screenshots/problem1_6.png">
    <br>
    <img src="screenshots/problem1_7.png">
    <br>
</p>
<p>PUT /users/:id (Update User id = 10)</p>
<p align="center">
    <img src="screenshots/problem1_8.png">
    <br>
</p>
<p>GET /users/:id (Get User by id = 10)</p>
<p align="center">
    <img src="screenshots/problem1_9.png">
    <br>
</p>
<p>DELETE /users/:id (Delete User by id = 10)</p>
<p align="center">
    <img src="screenshots/problem1_10.png">
    <br>
</p>
<p>GET /users (Get All User)</p>
<p align="center">
    <img src="screenshots/problem1_11.png">
    <br>
</p>
<br>

<h3>Problem 2 = Structuring Project with Layered Architecture</h3>
<h4>Code</h4>
<h5>Code Structure</h5>
<p align="center">
    <img src="screenshots/problem2_1.png">
    <br>
</p>
<h5>Config</h5>
<p align="center">
    <img src="screenshots/problem2_2.png">
    <br>
</p>
<h5>Controller</h5>
<p>User Controller</p>
<br>
<p align="center">
    <img src="screenshots/problem2_3.png">
    <br>
    <img src="screenshots/problem2_4.png">
    <br>
</p>
<p>Book Controller</p>
<br>
<p align="center">
    <img src="screenshots/problem2_5.png">
    <br>
    <img src="screenshots/problem2_6.png">
    <br>
</p>
<h5>Database</h5>
<p>Book</p>
<br>
<p align="center">
    <img src="screenshots/problem2_7.png">
    <br>
</p>
<p>User</p>
<br>
<p align="center">
    <img src="screenshots/problem2_8.png">
    <br>
</p>
<h5>Model</h5>
<p>Book</p>
<br>
<p align="center">
    <img src="screenshots/problem2_9.png">
    <br>
</p>
<p>User</p>
<br>
<p align="center">
    <img src="screenshots/problem2_10.png">
    <br>
</p>
<h5>Routes</h5>
<p align="center">
    <img src="screenshots/problem2_11.png">
    <br>
</p>
<h5>Main</h5>
<p align="center">
    <img src="screenshots/problem2_12.png">
    <br>
</p>
<h4>Output</h4>
<h5>User</h5>
<p>POST /users (Create New Users)</p>
<p align="center">
    <img src="screenshots/problem2_13.png">
    <br>
</p>
<p>PUT /users/:id (Update User id = 13)</p>
<p align="center">
    <img src="screenshots/problem2_14.png">
    <br>
</p>
<p>GET /users/:id (Get User by id = 13)</p>
<p align="center">
    <img src="screenshots/problem2_15.png">
    <br>
</p>
<p>DELETE /users/:id (Delete User by id = 13)</p>
<p align="center">
    <img src="screenshots/problem2_16.png">
    <br>
</p>
<p>GET /users (Get All User)</p>
<p align="center">
    <img src="screenshots/problem2_17.png">
    <br>
</p>
<br>
<h5>Book</h5>
<p>POST /books (Create New Book)</p>
<p align="center">
    <img src="screenshots/problem2_18.png">
    <br>
    <img src="screenshots/problem2_19.png">
    <br>
</p>
<p>PUT /books/:id (Update Book id = 5)</p>
<p align="center">
    <img src="screenshots/problem2_20.png">
    <br>
</p>
<p>GET /books/:id (Get Book by id = 5)</p>
<p align="center">
    <img src="screenshots/problem2_21.png">
    <br>
</p>
<p>DELETE /books/:id (Delete Book by id = 4)</p>
<p align="center">
    <img src="screenshots/problem2_22.png">
    <br>
</p>
<p>GET /books (Get All Book)</p>
<p align="center">
    <img src="screenshots/problem2_23.png">
    <br>
</p>
<br>
