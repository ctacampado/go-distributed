TARGET = login_svc
BIN = ./bin
CMD_BLD = go build
FLAGS = -o ${TARGET}
FILES = cmd/login/main.go

all: build
clean:
	rm ${BIN}/${TARGET}
build:
	${CMD_BLD} ${FLAGS} ${FILES}; mv ${TARGET} ${BIN}
