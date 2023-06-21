all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./qrcw.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.20.5 amd64
	sudo apk update

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/qrcw.exe $(DESTDIR)$(PREFIX)/bin/qrcw
	strip $(DESTDIR)$(PREFIX)/bin/qrcw
