DIST_DIR = dist
APP = bonusly-mp
TMP_BUILD := tmp/$(APP)

dist: dist-darwin dist-linux dist-windows

ensure-dist:
	mkdir -p dist tmp

dist-darwin: ensure-dist
	GOOS=darwin GOARCH=amd64 go build -o $(TMP_BUILD)
	tar -C tmp -zcvf $(DIST_DIR)/$(APP)-darwin.gz $(APP)
	rm $(TMP_BUILD)

dist-linux: ensure-dist
	GOOS=linux GOARCH=amd64 go build -o $(TMP_BUILD)
	tar -C tmp -zcvf $(DIST_DIR)/$(APP)-linux.gz $(APP)
	rm $(TMP_BUILD)

dist-windows: ensure-dist
	GOOS=windows GOARCH=amd64 go build -o $(TMP_BUILD)
	tar -C tmp -zcvf $(DIST_DIR)/$(APP)-windows.gz $(APP)
	rm $(TMP_BUILD)

.PHONY: dist
