<h1 align="center">Assignment 21 - Compute Services</h1>
<h2 align="center">Resume Materi</h2>
<hr>
<ul>
    <li>Pengertian Compute Services</li>
        <p>Compute Services adalah gabungan pemanfaatan teknologi komputer dan pengembangan berbasis Internet. Suatu cara atau metode yang memudahkan pengguna untuk mengakses informasi tanpa mengetahui apa yang ada didalamnya, ahli dengannya, atau memiliki kendali terhadap infrastruktur teknologi yang membantunya</p>
    <li>Software Deployment</li>
        <p>Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi yang telah dikerjakan oleh para orang-orang yang ahli di bidang programmer. Cara penyebarannya pun sangat beragam, tergantung dari jenis aplikasinya. Jika memilih aplikasi Web, maka akan di hosting pada server. Sedangkan jika aplikasi mobile, akan terdapat dua deployment. Pertama adalah deployment untuk aplikasi ke Playstore atau Appstore, dan yang kedua adalah deployment API (backend) ke server.</p>
    <li>Strategi Deployment</li>
        <p>Strategi dalam deployment :</p>
        <p>a.	Big-Bang deployment strategy (atau sering disebut Replace/Recreate deployment strategy) : Yang dimana sifatnya menimpa dan mengganti (mereplace) aplikasi yang aktif secara langsung.</p>
        <p>b.	Rollout deployment strategy : Dengan metode ini, kita melakukan deployment secara bertahap per-server yang hidup. Dan jika satu server saja langsung error, kita dapat langsung rollback tanpa melanjut deploy kesemua server.</p>
        <p>c.	Blue/Green deployment strategy : Konsepnya cukup sederhana, pertama-tama, kita akan membuat satu environment yang serupa dengan yang sedang aktif/live, kemudian kita pun melakukan switching request ke environment baru tersebut.</p>
        <p>d.	A/B deployment strategy : Strategi ini biasanya lebih fokus pada user experience dan layout(UX/UI). Biasanya A/B deployment lebih ke user sentris. Setengah user akan menerima fitur/layout A dan setengah lagi mendapat fitur/layout B, sehingga setiap user bisa mendapatkan tampilan yang beda.</p>
        <p>e.	Canary deployment strategy : Strategi ini lebih advance dari semua metode release tersebut diatas. Prinsip kerjanya mirip seperti Rollout Deployment, tetapi bedanya, jika pada Rollout Deployment, ketika aplikasi di deploy pada satu server, maka server tersebut akan langsung kebagian request dari user sama rata dengan server lainnya.</p>
</ul>
<br>
