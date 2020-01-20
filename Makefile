UAUTH_DIR = build/uauth
JWTAUTH_DIR = build/jwtauth

all: build-uauth build-jwtauth
clean:
	cd ${UAUTH_DIR} && $(MAKE) clean
	cd ${JWTAUTH_DIR} && $(MAKE) clean
build-uauth:
	cd ${UAUTH_DIR} && $(MAKE) all
build-jwtauth:
	cd ${JWTAUTH_DIR} && $(MAKE) all