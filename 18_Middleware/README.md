<h1 align="center">Assignment 17 - Middleware</h1>
<h2 align="center">Resume Materi</h2>
<hr>
<ul>
    <li>Pengertian Middleware</li>
        <p>middleware adalah entitas yang terhubung ke pemrosesan request/response server, middleware berisi sebuah blok code yang akan diproses sebelum ataupun sesudah http request di proses</p>
        <p>Echo #Pre() : dieksekusi sebelum route memproses request</p>
        <p>Echo #Use() : dieksekusi setelah route memproses request dan memiliki akses penuh ke echo.Context API</p>
    <li>Logger Middleware</li>
        <p>Digunakan untuk mencatatat informasi HTTP request, history atau jejak, serta untuk sumber data yang digunakan analisis</p>
    <li>Auth Middleware</li>
        <p>Authentication digunakan untuk mengamankan sebuah data atau mengidentifikasi pengguna</p>
        <p><b>Basic Authentication</b></p>
        <p>Basic Authentication adalah salah satu teknik authentication, metode ini membutuhkan informasi nama pengguna dan kata sandi untuk dimasukkan ke dalam header permintaan, username dan kata sandi dalam request melalui header di encode menggunakan base64</p>
        <p>Header Request : 'Authorization: Basic ' + base64encode('username:password')</p>
        <p>Sedangkan untuk JWT Authentication Middleware menggunakan token JWT untuk melakukan Authentication</p>
        <p>JWT terdiri dari 3 bagian : Header, Payload, Verify Signature</p>
        <p>Format Header Request JWT Auth : 'Authorization: Bearer ' + Token JWT</p>
</ul>
<br>

<h2>Logging and JWT Authentication</h2>
<h3>Logging Middleware</h3>
<h4>Code</h4>
<p align="center">
    <img src="screenshots/code_2.png">
    <br>
</p>
<h3>JWT Middleware</h3>
<h4>Secret Key</h4>
<p align="center">
    <img src="screenshots/code_1.png">
    <br>
</p>
<h4>Models (Add User Response Token JWT)</h4>
<p align="center">
    <img src="screenshots/code_4.png">
    <br>
</p>
<h4>Login Controller (Create Token JWT)</h4>
<p align="center">
    <img src="screenshots/code_5.png">
    <br>
</p>
<h4>Code Middleware</h4>
<p align="center">
    <img src="screenshots/code_3.png">
    <br>
</p>
<h4>Route (Group JWT Auth Middleware)</h4>
<p align="center">
    <img src="screenshots/code_6.png">
    <br>
</p>
<h4>Main (Implementasi Log Middleware)</h4>
<p align="center">
    <img src="screenshots/code_7.png">
    <br>
</p>

<h4>Output</h4>
<p>POST /login (Create Token JWT)</p>
<p align="center">
    <img src="screenshots/output_1.png">
    <br>
</p>
<p>GET All User /users (dengan token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjU0ODg0OTEsIm5hbWUiOiJEZWRlbiIsInVzZXJJZCI6OH0.2TYpfKxYcVczQAvryaqDHkwWgdREchx6jmcjK3eG_pU")</p>
<p align="center">
    <img src="screenshots/output_2.png">
    <br>
</p>
<p>GET /users/:id Not Authenticated</p>
<p align="center">
    <img src="screenshots/output_3.png">
    <br>
</p>
<p>GET /books Not Authenticated</p>
<p align="center">
    <img src="screenshots/output_4.png">
    <br>
</p>
<p>GET /books/:id Not Authenticated</p>
<p align="center">
    <img src="screenshots/output_5.png">
    <br>
</p>
<p>Logger Middleware</p>
<p align="center">
    <img src="screenshots/output_6.png">
    <br>
</p>