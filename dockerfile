# Menggunakan image Golang sebagai base
FROM golang:1.22 AS builder

# Set working directory
WORKDIR /app

# Menyalin semua file ke working directory
COPY . .

# Membangun aplikasi dan mencetak log kesalahan
RUN go build -o myapp . || { echo 'Build failed'; exit 1; }

# Stage kedua untuk menjalankan aplikasi
FROM alpine:latest

# Menyalin binary dari stage builder
COPY --from=builder /app/myapp /myapp

# Mengatur port yang digunakan
EXPOSE 8080

# Menjalankan aplikasi
CMD ["/myapp"]
