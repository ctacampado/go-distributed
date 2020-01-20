UAUTH_DIR = build/uauth
JWTAUTH_DIR = build/jwtauth

all: build-uauth build-jwtauth
clean: clean-uauth clean-jwtauth
clean-uauth:
	cd ${UAUTH_DIR} && $(MAKE) clean
clean-jwtauth:
	cd ${JWTAUTH_DIR} && $(MAKE) clean
build-uauth:
	cd ${UAUTH_DIR} && $(MAKE) all
build-jwtauth:
	cd ${JWTAUTH_DIR} && $(MAKE) all