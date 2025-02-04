# เริ่มต้นหยิบ golang image มาเป็น base image
FROM golang:latest

# ทำการกำหนด path /app เป็น directory เริ่มต้น
WORKDIR /app

# Copy go.mod และ go.sum ไฟล์เข้ามา
COPY go.mod go.sum ./

# Download dependencies ทั้งหมด
RUN go mod download

# Copy the source จาก directory ปัจจุบัน สู่ working directory ภายใน container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8000 ออกมาภายนอก
EXPOSE 8000

# ทำการ run command ผ่าน binary file ที่ build มาได้
CMD ["./main"]