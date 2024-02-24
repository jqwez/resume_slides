FROM node:latest AS frontend

WORKDIR /app

COPY web /app
RUN npm install
RUN npm run build

FROM golang:latest AS backend

WORKDIR /app

COPY src /app
RUN go build -o server

FROM golang:latest

WORKDIR app
COPY --from=frontend /app/dist /app/static
COPY --from=backend /app/server /app/server

EXPOSE 8000

CMD ["/app/server"]
