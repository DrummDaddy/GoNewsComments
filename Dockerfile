FROM golang:1.22-buster AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Загружаем зависимости и создаем исполняемый файл
RUN go mod init comment-censorship-service && \
    go mod tidy && \
    go build -o comment-censorship-service main.go

# Используем минимальный образ для финального контейнера
FROM debian:buster-slim

# Устанавливаем необходимые сертификаты
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем исполняемый файл из контейнера сборки
COPY --from=builder /app/comment-censorship-service .

# Слушаем на 8080 порту
EXPOSE 8080

# Указываем команду для запуска приложения
CMD ["./comment-censorship-service"]