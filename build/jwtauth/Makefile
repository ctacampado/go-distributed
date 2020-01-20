TARGET = jwtauth
ROOTDIR = ../../
BIN = ${ROOTDIR}/bin
CMD_BLD = go build
FLAGS = -o ${TARGET}
FILES = ${ROOTDIR}/cmd/jwt-auth-server/main.go

all: build
clean:
	rm ${BIN}/${TARGET}
build:
	${CMD_BLD} ${FLAGS} ${FILES}; mv ${TARGET} ${BIN}
