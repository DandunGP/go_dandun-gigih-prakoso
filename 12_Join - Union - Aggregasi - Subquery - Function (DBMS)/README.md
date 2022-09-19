<h1 align="center">Assignment 11 - Join - Union - Agregasi - Subquery - Function (DBMS)</h1>
<h2 align="center">Resume Materi</h2>
<hr>

<ul>
    <li>Pengertian Join, Union dan Agregasi</li>
        <p>Join adalah Sebuah klausa untuk menggabungkan data dari 2 atau lebih table yang saling berhubungan atau memiliki 1 field yang sama</p>
        <p>Join terdapat beberapa jenis antara lain : INNER JOIN, LEFT JOIN, RIGHT JOIN dan CROSSJOIN</p>
        <p>Union adalah sebuah klausa untuk menggabungkan data dari 2 table atau lebih serta harus memiliki jumlah field yang sama</p>
        <p>Agregasi adalah fungsi dimana nilai dalam beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal</p>
        <p>Agregasi terdiri dari : MIN, MAX, SUM, AVG, COUNT dan HAVING</p>
    <li>Pengertian SubQuery</li>
        <p>SubQuery atau Inner Query adalah Query yang terdapat didalam Query SQL lain</p>
        <p>Sebuah Subquery biasa digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil</p>
        <p>SubQuery dapat digunakan dengan SELECT, INSERT, UPDATE dan DELETE disertai dengan statement operator seperti =, <, >, >=, <=</p>
    <li>Pengertian Function dalam MySQL</li>
        <p>Function didalam MySQL adalah sekumpulan statement yang akan mengembalikan sebuah nilai balik pada saat pemanggilannya</p>
</ul>
<br>

<h2>Insert</h2>
<p>Insert 5 Operators</p>
<p align="center">
    <img src="screenshots/problem_1.png">
</p>
<p>Insert 3 Product Type</p>
<p align="center">
    <img src="screenshots/problem_2.png">
</p>
<p>Insert 2 Product dengan product type id = 1 dan operator id = 3</p>
<p>Insert 3 Product dengan product type id = 2 dan operator id = 1</p>
<p>Insert 3 Product dengan product type id = 3 dan operator id = 4</p>
<p align="center">
    <img src="screenshots/problem_3.png">
</p>
<p>Insert Product Description pada setiap product</p>
<p align="center">
    <img src="screenshots/problem_4.png">
</p>
<p>Insert 3 Payment Method</p>
<p align="center">
    <img src="screenshots/problem_5.png">
</p>
<p>Insert 5 User</p>
<p align="center">
    <img src="screenshots/problem_6.png">
</p>
<p>Insert 3 transaksi di masing masing user</p>
<p align="center">
    <img src="screenshots/problem_7.png">
</p>
<p>Insert 3 Product di masing masing transaksi</p>
<p align="center">
    <img src="screenshots/problem_8.png">
</p>

<h2>Select</h2>
<p>Tampilkan nama user dengan gender M atau Laki-laki</p>
<p align="center">
    <img src="screenshots/problem_9.png">
</p>
<p>Tampilkan product dengan id = 3</p>
<p align="center">
    <img src="screenshots/problem_10.png">
</p>
<p>Tampilkan Data Pelanggan dengan created_at 7 hari kebelakang dan memiliki gender mengandung kata M</p>
<p align="center">
    <img src="screenshots/problem_11.png">
</p>
<p>Hitung jumlah user dengan status gender perempuan</p>
<p align="center">
    <img src="screenshots/problem_12.png">
</p>
<p>Tampilkan data pelanggan dengan urutan berdasarkan gender</p>
<p align="center">
    <img src="screenshots/problem_13.png">
</p>
<p>Tampilkan 5 data pada data product</p>
<p align="center">
    <img src="screenshots/problem_14.png">
</p>
<h2>Update</h2>
<p>Ubah data product id 1 dengan nama 'product dummy'</p>
<p align="center">
    <img src="screenshots/problem_15.png">
</p>
<p>Update qty = 3 pada transaction detail dengan product id 1</p>
<p align="center">
    <img src="screenshots/problem_16.png">
</p>
<h2>Delete</h2>
<p>Delete data pada table product dengan id 1</p>
<p align="center">
    <img src="screenshots/problem_17.png">
</p>
<p>Delete data pada tabel product dengan product type id 1</p>
<p align="center">
    <img src="screenshots/problem_18.png">
</p>
<h2>Join, Union, Sub Query, Function</h2>
<p>Gabungkan data transaksi dari user id 1 dan user id 2</p>
<p align="center">
    <img src="screenshots/problem_19.png">
</p>
<p>Tampilkan jumlah harga transaksi user id 1</p>
<p align="center">
    <img src="screenshots/problem_20.png">
</p>
<p>Tampilkan total transaksi dengan product type 2</p>
<p align="center">
    <img src="screenshots/problem_21.png">
</p>
<p>Tampilkan semua field table product, field name table product type yang saling berhubungan</p>
<p align="center">
    <img src="screenshots/problem_22.png">
</p>
<p>Tampilkan semua field table transaction, field name table product dan field name table user</p>
<p align="center">
    <img src="screenshots/problem_23.png">
</p>
<p>Buat function setelah data transaksi dihapus maka transaksi detail tehapus juga dengan transaksi id yang dimaksud</p>
<p align="center">
    <img src="screenshots/problem_24.png">
</p>
<p>Buat function setelah data transaksi detail terhapus maka data total_qty terupdate berdasarkan qty data transaksi id yang terhapus</p>
<p align="center">
    <img src="screenshots/problem_25.png">
</p>
<p>tampilkan data product yang tidak pernah ada didalam table transaction detail dengan subquery</p>
<p align="center">
    <img src="screenshots/problem_26.png">
</p>
