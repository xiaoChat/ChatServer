FROM golang:1.16.2-buster

WORKDIR /go/src/app

RUN go version \
	# apt source
	&& echo "deb http://mirrors.aliyun.com/debian/ buster main non-free contrib \n \
		deb-src http://mirrors.aliyun.com/debian/ buster main non-free contrib \n \
		deb http://mirrors.aliyun.com/debian-security buster/updates main \n \
		deb-src http://mirrors.aliyun.com/debian-security buster/updates main \n \
		deb http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib \n \
		deb-src http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib \n \
		deb http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib \n \
		deb-src http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib" > /etc/apt/sources.list \
	# go proxy
	&& go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.io,direct \
	# down go protoc buidl module
	&& go get github.com/gogo/protobuf/protoc-gen-gofast \
	&& go get github.com/gogo/protobuf/protoc-gen-gogofaster \
	&& go get github.com/gogo/protobuf/protoc-gen-gogoslick \
	# install protobuf
	# && apt update && apt install unzip \
	# && wget -O protobuf.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.15.5/protoc-3.15.5-linux-x86_64.zip \
	# && unzip protobuf.zip -d /usr/local \
	&& echo "complete"

CMD [ "sh", "-c", "tail -f /dev/null" ]