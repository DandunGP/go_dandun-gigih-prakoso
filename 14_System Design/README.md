<h1 align="center">Assignment 13 - System Design</h1>
<h2 align="center">Resume Materi</h2>
<hr>

<ul>
    <li>Pengertian Diagram</li>
        <p>Diagram adalah representasi simbolis dari informasi menggunakan teknik visualisasi</p>
        <p>Jenis Jenis Diagram dalam System antara lain : Flowchart, Use Case Diagram (ringkasan detail pengguna sistem dan interaksi pengguna dengan sistem), Entity Relationship Diagram (Jenis Flowchart yang menggambarkan bagaimana tiap entitas saling berhubungan satu sama lain dalam suatu sistem)</p>
    <li>Pengertian System Design</li>
        <p>Karakteristik utama Sistem Terdistribusi : </p>
        <p>- Scalability adalah kemampuan system, process, atau network untuk tumbuh dan mengelola peningkatan permintaan</p>
        <p>- Reliability adalah kemungkinan suatu sistem akan gagal dalam periode tertentu</p>
        <p>- Availability adalah waktu sistem tetap beroperasi untuk melakukan fungsi yang diperlukan dalam periode tertentu</p>
        <p>- Efficiency adalah daya guna sistem yang mampu memenuhi fungsi yang diperlukan</p>
        <p>- Serviceability and Manageability adalah kesederhanaan dan kecepatan di mana sistem dapat diperbaiki atau dipelihara</p>
    <li>Pengertian Job/Work Queue, Load Balancing, Monolithic and Microservices</li>
        <p>Dalam perangkat lunak sistem, Job Queue adalah struktur data yang dikelola oleh perangkat lunak penjadwal pekerjaan yang berisi pekerjaan untuk dijalankan</p>
        <p>Work Queue adalah Framework untuk membangun aplikasi master-worker besar yang menjangkau ribuan mesin yang diambil dari cluster, cloud, dan grid</p>
        <p>Load Balancing adalah komponen penting lainnya dari setiap sistem terdistribusi karena membantu untuk menyebarkan lalu lintas di sekelompok server untuk meningkatkan daya respon dan ketersediaan aplikasi, situs web, atau basis data</p>
        <p>Aplikasi monolithic memiliki single basis kode dengan banyak modul didalamnya</p>
        <p>Aplikasi Microservices adalah aplikasi dibuat dan terdiri dari banyak layanan yang digabungkan serta dapat digunakan secara independen</p>
</ul>
<br>

<h2>Problem 1 - System Pencatatan Pengeluaran Harian</h2>
<p>ER Diagram</p>
<p align="center">
    <img src="screenshots/problem1_erd1.png">
    <br>
</p>
<br>
<p align="center">
    <img src="screenshots/problem1_erd2.png">
    <br>
</p>
<br>
<p>Use Case Diagram</p>
<p align="center">
    <img src="screenshots/problem1_usecase.png">
    <br>
</p>

<h2>Problem 2 - Query</h2>
<p>Query</p>
<p>MySQL = SELECT * FROM user;</p>
<p>Redis = GET user (key) atau HGETALL user</p>
<p>neo4j = RETURN *; atau (MATCH (u:user) RETURN u;)</p>
<p>cassandra = SELECT * FROM user;</p>