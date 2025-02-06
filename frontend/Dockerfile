# Используем Node.js для сборки фронтенда
FROM node:18 AS build
WORKDIR /app
COPY . .
RUN npm install && npm run build

# Запускаем через Nginx
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
