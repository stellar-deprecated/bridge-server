#! /usr/bin/env bash
set -e

DIST="dist"
VERSION=$(git describe --always --dirty --tags)
GOARCH=amd64
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

build() {
  NAME=$1
  GOOS=$2
  RELEASE="$NAME-$VERSION-$GOOS-$GOARCH"
  PKG_DIR="$DIST/$RELEASE"

  # do the actual build
  GOOS=$GOOS GOARCH=$GOARCH gb build  -ldflags "-X main.version=$VERSION"

  # make package directory
  rm -rf $PKG_DIR
  mkdir -p $PKG_DIR
  cp bin/$(srcBin $NAME $GOOS) $PKG_DIR/$(destBin $NAME $GOOS)
  cp CHANGELOG.md $PKG_DIR/
  cp LICENSE.txt $PKG_DIR/
  cp ${NAME}_example.cfg $PKG_DIR/
  cp readme_$NAME.md $PKG_DIR/

  # TODO: add platform specific install intstructions

  # zip/tar package directory
  pkg $GOOS $RELEASE
}

srcBin() {
  NAME=$1
  GOOS=$2
  BIN="$NAME-$GOOS-$GOARCH"

  if [ "$GOOS" = "windows" ]; then
    BIN+=".exe"
  fi

  echo $BIN
}

destBin() {
  if [ "$2" = "windows" ]; then
    echo "$NAME.exe"
  else
    echo "$NAME"
  fi
}

pkg() {
  GOOS=$1
  RELEASE=$2

  if [ "$GOOS" = "windows" ]; then
    pushd $DIST
    zip $RELEASE.zip $RELEASE/*
    popd
  else
    tar -czf $DIST/$RELEASE.tar.gz -C $DIST $RELEASE
  fi

  rm -rf $DIST/$RELEASE
}

build bridge darwin
build bridge linux
build bridge windows

build compliance darwin
build compliance linux
build compliance windows
