#!/usr/bin/env bash

PKGDIR=qrcodewriter-`cat current_pkg_release`_amd64

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
mv control ${PKGDIR}/DEBIAN/
mv preinst ${PKGDIR}/DEBIAN/

echo "Building binary from source"
cd ../src
go build -o ../__debian/${PKGDIR}/opt/bin/qrcw .
strip ../__debian/${PKGDIR}/opt/bin/qrcw
sudo chown 0:0 ../__debian/${PKGDIR}/opt/bin/qrcw

echo "Binary built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
