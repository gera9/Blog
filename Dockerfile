FROM golang:1.25-alpine AS builder

WORKDIR /build

RUN apk update && apk add tzdata

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -o blog

# Production stage
FROM scratch AS production

WORKDIR /prod

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /build/blog .

CMD [ "/prod/blog" ]