# Собираем в гошке
FROM golang:1.20 as build

ENV BIN_FILE /opt/survey/survey-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-a -installsuffix cgo \
        -o ${BIN_FILE} ./cmd/survey/main.go ./cmd/survey/version.go

# На выходе тонкий образ
FROM scratch

LABEL ORGANIZATION="survey"
LABEL SERVICE="survey"
LABEL MAINTAINERS="lixoi@list.ru"

ENV BIN_FILE "/opt/survey/survey-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /opt/survey/config.json
COPY ./cmd/survey/config.json ${CONFIG_FILE}

ENTRYPOINT ["/opt/survey/survey-app", "--config=/opt/survey/config.json"]
#CMD ${BIN_FILE} --config=${CONFIG_FILE}
