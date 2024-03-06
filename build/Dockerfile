FROM golang:1.20 as build

ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

RUN apt update && apt upgrade -y && apt install -y protobuf-compiler

COPY . ${CODE_DIR}

RUN go mod tidy
RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN make gen
RUN make build

# --- build distroless docker image ---
FROM scratch

LABEL ORGANIZATION="survey"
LABEL SERVICE="survey"
LABEL MAINTAINERS="lixoi@list.ru"

ENV CODE_DIR "/go/src/"
ENV BIN_FILE "bin/survey"
ENV BIN "/opt/survey/survey-app"

COPY --from=build ${CODE_DIR}${BIN_FILE} ${BIN}
ENV CONFIG_FILE /opt/survey/config.json
COPY ./cmd/survey/config.json ${CONFIG_FILE}

COPY --from=build ${CODE_DIR}certs /certs
COPY --from=build ${CODE_DIR}migrations /migrations

ENTRYPOINT ["/opt/survey/survey-app", "--config=/opt/survey/config.json"]
