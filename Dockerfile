FROM  golang:1.17.5-alpine

# RUN apk add --no-cache git
# RUN go get github.com/golang/dep/cmd/dep

WORKDIR /app
# COPY Gopkg.lock Gopkg.toml /go/src/crowdfund/

COPY . .

RUN go get -u github.com/gin-gonic/gin 
RUN go get -u github.com/joho/godotenv
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql
RUN go build -o ./bin/crowdfund-api

EXPOSE 8080

CMD ./bin/crowdfund-api