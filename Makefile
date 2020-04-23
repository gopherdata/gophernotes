BIN_DIR=build
BIN=$(BIN_DIR)/gophernotes
BUILD_OPTS=-v -buildmode exe -trimpath -o ${BIN}

.PHONY: all
all:
	[ -d ${BIN_DIR} ] || mkdir -p ${BIN_DIR}
	go build ${BUILD_OPTS}

.PHONY: clean
clean:
	if [ -f ${BIN} ] ; then rm -v ${BIN} ; fi

.PHONY: clean_notebooks
clean_notebooks:
	rm -v Untitled*.ipynb
