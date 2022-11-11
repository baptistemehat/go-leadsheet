BINARY_FILE=core

all: build

build:
	go build -o ${BINARY_FILE}

run:
	go build -o ${BINARY_FILE}
	./${BINARY_FILE}

clean:
	rm ${BINARY_FILE} -rf
	rm latex/tmp -rf