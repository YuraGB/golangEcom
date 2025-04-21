# Використовуємо образ Go
FROM golang:1.24.2

# Встановлюємо робочу директорію
WORKDIR /app

# Копіюємо go.mod та go.sum для того, щоб завантажити залежності
COPY go.mod go.sum ./
RUN go mod download

# Копіюємо увесь код у контейнер
COPY . ./

# Встановлюємо air для hot-reload
RUN go install github.com/air-verse/air@latest

# Компілюємо Go програму перед запуском
RUN go build -o /app/tmp/main ./cmd/main.go

# Відкриваємо порт для доступу до сервера
EXPOSE 8081

# Запускаємо air для hot-reload
CMD ["air", "-c", ".air.toml"]
