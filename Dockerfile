# Stage 1: Build Go Back End
FROM golang:1.16-alpine as backend-build
WORKDIR /app
COPY cmd /app/cmd
COPY config /app/config
COPY internal /app/internal
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
LABEL maintainer="fabiancdng <contact@fabiancdng.com>"
RUN go build -v -o goshortrr /app/cmd/goshortrr/main.go

# Stage 2: Build React Front End
FROM node:16-alpine as frontend-build
WORKDIR /app
COPY web /app
RUN npm install
RUN npm run build

# Stage 3: Hosting (only Alpine running the Go binary)

FROM alpine:latest as final
WORKDIR /app
COPY --from=backend-build /app .
COPY --from=frontend-build /app/build ./web/build

EXPOSE 4000
CMD [ "/app/goshortrr" ]