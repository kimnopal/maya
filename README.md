![✨ Mock Up](https://github.com/kimnopal/maya/assets/88240429/431922f0-4ef3-41ff-85c7-03d1c0bf8975)

# 👋 About Maya: Build your Dream Team!

Finding the right team members for projects or competitions is a common challenge, as we've experienced firsthand. Existing methods are often time-consuming and rely on personal networks, limiting the pool of potential collaborators. To address this, we're working on a user-friendly platform that streamlines the process, making it easy for individuals and teams to connect based on specific skills and roles. Our goal is to enhance efficiency, diversity, and collaboration in team formation, ultimately improving the success rates of various initiatives.

## Mutex Team

Member :

1. Naufal Hakim
2. M. Saujana Shafi Kehaulani
3. Muhammad Nur Bijak Bestari
4. Daffa Randika

## Screenshot

<img src="https://github.com/kimnopal/maya/assets/88240429/3ef6800d-03af-49b9-8925-1667a625de0f" alt="Screenshot 1" width="300"><br><br>

<img src="https://github.com/kimnopal/maya/assets/88240429/dd88e6ae-0bff-472a-a103-9db244effd29" alt="Screenshot 2" width="500"><br><br>

## Penggunaan

1. Install Golang, NodeJS, MySQL
2. Pindah ke direktori backend
3. Jalankan perintah berikut untuk menginstall dependency

```bash
$ go get
```

```bash
$ go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

4. Buat database dengan nama `maya`
5. Jalankan perintah berikut untuk melakukan migration

```bash
$ migrate -database "mysql://root@tcp(localhost:3306)/maya" -path db/migrations up
```

6. Jalankan perintah berikut untuk menjalankan program backend

```bash
$ go run main.go
```

7. Pindah ke direktori frontend
8. Jalankan perintah berikut untuk menginstall dependency

```bash
$ npm install
```

9. Jalankan perintah berikut untuk menjalankan program frontend

```bash
$ npm run dev
```
