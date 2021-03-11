FROM registry.cn-shenzhen.aliyuncs.com/tofu/dt-golang-base:open
ENV GOPROXY https://goproxy.cn,direct
ENV ES_HOST 127.0.0.1
ENV ELASTIC_APM_SERVER_URL http://127.0.0.1:8210
ENV ELASTIC_APM_LOG_FILE stdout
ENV ELASTIC_APM_LOG_LEVEL debug
ENV ELASTIC_APM_TRANSACTION_SAMPLE_RATE 1
WORKDIR /home
COPY go.mod .
RUN go mod download
ADD . /home
RUN GOOS=linux GOARCH=amd64 go build -o halfway_demo.bin ./cmd
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENTRYPOINT ["/home/halfway_demo.bin"]