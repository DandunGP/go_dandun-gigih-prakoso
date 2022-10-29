<h1 align="center">Assignment 20 - Docker</h1>
<h2 align="center">Resume Materi</h2>
<hr>
<ul>
    <li>Pengertian Docker</li>
        <p>Docker adalah layanan yang menyediakan kemampuan untuk mengemas dan menjalankan sebuah aplikasi dalam sebuah lingkungan terisolasi yang disebut dengan container</p>
    <li>Container vs Virtual Machine</li>
        <p>Container adalah  paket atau aplikasi yang mengandalkan isolasi virtual untuk menjalankan aplikasi yang dapat menjalankan sistem operasi kernel secara simultan tanpa memerlukan mesin virtual (VMs)</p>
        <p>Virtual Machine adalah sebuah emulasi dari sebuah sistem komputer. Virtual Machine dapat membuat kita membagi resource hardware dari satu hardware fisik menjadi beberapa sistem komputer</p>
    <li>Syntax Docker</li>
        <p>a.	FROM : mendapatkan gambar dari register docker</p>
        <p>b.	RUN : Menjalankan perintah bash saat membangun container</p>
        <p>c.	ENV : Mengatur variable di dalam container</p>
        <p>d.	ADD : Menyalin file dengan beberapa proses lain</p>
        <p>e.	COPY : Menyalin filenya</p>
        <p>f.	WORKDIR : Mengatur direktori file yang berfungsi</p>
        <p>g.	ENTRYPOINT : Menjalankan perintah saat selesai membangun wadah</p>
        <p>h.	CMD : Menjalankan perintah tetapi bisa ditimpa</p>
</ul>
<br>

<h2>Docker</h2>
<h3>Dockerfile</h3>
<p align="center">
    <img src="screenshots/1.png">
    <br>
</p>
<h3>Docker Compose</h3>
<p align="center">
    <img src="screenshots/2.png">
    <br>
    <img src="screenshots/3.png">
    <br>
</p>
<h3>Docker Build Image</h3>
<p align="center">
    <img src="screenshots/4.png">
    <br>
</p>
<h3>Docker Compose Build and Deploy</h3>
<p align="center">
    <img src="screenshots/5.png">
    <br>
</p>
<h3>Docker image list</h3>
<p align="center">
    <img src="screenshots/6.png">
    <br>
</p>
<h3>Docker container list</h3>
<p align="center">
    <img src="screenshots/7.png">
    <br>
</p>
<h3>Docker container push</h3>
<p align="center">
    <img src="screenshots/8.png">
    <br>
</p>