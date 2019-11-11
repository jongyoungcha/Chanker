WORK_DIR:=$(shell pwd)
BUILD_DIR:=$(WORK_DIR)/build


chanker:
	cd $(WORK_DIR)/cmd/chanker && \
	go build -o $(BUILD_DIR)/chanker







