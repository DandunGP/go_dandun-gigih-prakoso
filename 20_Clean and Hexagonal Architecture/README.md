<h1 align="center">Assignment 19 - Clean and Hexagonal Architecture</h1>
<h2 align="center">Resume Materi</h2>
<hr>
<ul>
    <li>Structure Code Clean Architecture</li>
        <p>Controller : Berisi Code yang berhubungan langsung ke user interface</p>
        <p>Repository : Berisi code yang berhubungan langsung dengan database</p>
        <p>Usecase : Berisi bisnis logic yang dipakai</p>
    <li>Pengertian Usecase dan Repository</li>
        <p>Usecase merupakan layer yang bertugas sebagai pengontrol yakni menangani bisnis logic pada setiap domain, layer ini juga bertugas memilih repository apa yang akan digunakan</p>
        <p>Repository merupakan layer yang menyimpan database handler. Querying, Inserting, Deleting akan dilakukan pada layer ini. tidak ada business logic disini</p>
    <li>Manfaat Clean Architecture</li>
        <p>Code menjadi lebih rapi dan lebih mudah untuk di maintenance</p>
</ul>
<br>

<h2>Clean Architecture</h2>
<h3>Structure Code</h3>
<p align="center">
    <img src="screenshots/1.png">
    <br>
</p>
<h3>Create Token Use Email and Password User</h3>
<p align="center">
    <img src="screenshots/2.png">
    <br>
</p>
<h3>Create User Use Token (eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImR1ZHVuQGdtYWlsLmNvbSIsImV4cCI6MTY2NjUyNTc2OSwidXNlcklkIjoxM30.Duwa-b6SgegMfW8nI3IzMqjudhO4e8uCj4t-_yIAI6k)</h3>
<p align="center">
    <img src="screenshots/3.png">
    <br>
    <img src="screenshots/4.png">
    <br>
</p>
<h3>Get All User Use Token (eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImR1ZHVuQGdtYWlsLmNvbSIsImV4cCI6MTY2NjUyNTc2OSwidXNlcklkIjoxM30.Duwa-b6SgegMfW8nI3IzMqjudhO4e8uCj4t-_yIAI6k)</h3>
<p align="center">
    <img src="screenshots/5.png">
    <br>
</p>