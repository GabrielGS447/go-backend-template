FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go-backend-template

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /go-backend-template /go-backend-template
EXPOSE 8080
USER nonroot:nonroot
ENV GO_ENV=production \
PORT=8080 \
MONGODB_URI=mongodb://mongodb \
DATABASE=go-backend-template
ENTRYPOINT ["/go-backend-template"]
