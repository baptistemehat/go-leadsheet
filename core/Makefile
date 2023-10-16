BINARY_DIR=bin
BINARY_FILE=${BINARY_DIR}/core
DEFAULT_CONFIG_FILE=../default-config.yaml

all: build run

build:
	mkdir -p ${BINARY_DIR}
	go build -o ${BINARY_FILE}

run:
	./${BINARY_FILE} -c ${DEFAULT_CONFIG_FILE}

clean:
	rm ${BINARY_FILE} -rf
	rm latex/tmp -rf