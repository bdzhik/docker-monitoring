# docker-monitoring

Docker Monitoring
Этот проект представляет собой систему мониторинга с использованием PostgreSQL, Go (бэкенд), Node.js (фронтенд) и Docker Compose.

📌 Стек технологий
Backend: Go
Database: PostgreSQL
Frontend: Node.js + Nginx
Monitoring Service: Pinger на Go
Containerization: Docker & Docker Compose
🚀 Запуск проекта
1. Установи зависимости
Перед запуском убедись, что у тебя установлены:

Docker
Docker Compose
Проверь установку командой:

docker --version
docker-compose --version

Если Docker и Compose установлены, продолжай.

2. Клонируй репозиторий

git clone https://github.com/ТВОЙ_РЕПО.git
cd docker-monitoring

3. Собери и запусти контейнеры

docker-compose up --build

После запуска контейнеров ты должен увидеть логи PostgreSQL, backend, frontend и pinger.

4. Проверь, что всё работает
✅ Проверка работы бэкенда
Отправь тестовый ping:

curl -X POST http://localhost:8080/ping -H "Content-Type: application/json" \
-d '{"ip":"192.168.1.1", "time":123, "last_active":"2025-02-06T12:13:25Z"}'

Посмотри список всех пингов:


curl -X GET http://localhost:8080/pings

Если ты видишь JSON-ответ с данными пингов, значит всё работает! 🎉

⚙️ Структура проекта
docker-monitoring/
├── backend/        # Бэкенд на Go
├── frontend/       # Фронтенд (Node.js + Nginx)
├── pinger/         # Pinger-сервис на Go
├── postgres/       # Конфигурация PostgreSQL
├── docker-compose.yml  # Docker Compose для запуска
└── README.md       # Документация
🛠️ Основные команды
📌 Остановить контейнеры

docker-compose down

📌 Перезапустить с очисткой

docker-compose down -v
docker-compose up --build

📌 Посмотреть логи

docker logs docker-monitoring-backend-1

📌 Войти в базу данных

docker exec -it docker-monitoring-postgres-1 psql -U user -d user