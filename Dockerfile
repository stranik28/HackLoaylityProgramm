# Используем официальный образ Go как базовый образ
FROM golang:1.22

# Установка переменной окружения для указания рабочей директории внутри контейнера
WORKDIR /build

# Копируем файлы go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы проекта в рабочую директорию контейнера
COPY . .

# Сборка бинарника
RUN GOOS=linux go build -o app ./cmd/server

EXPOSE 8080

CMD ["./app"]

