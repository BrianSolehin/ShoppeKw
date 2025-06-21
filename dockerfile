# 1. Ambil image Golang sebagai base
FROM golang:1.23.3

# 2. Set working directory
WORKDIR /app

# 3. Copy semua file dari project lokal ke dalam container
COPY . .

# 4. Download dependency
RUN go mod tidy

# 5. Build aplikasi jadi file binary
RUN go build -o main .

# 6. Expose port (ganti jika aplikasimu pakai port lain)
EXPOSE 8080

# 7. Jalankan aplikasinya
CMD ["./main"]
