#!/bin/sh -e

SNOWBALL_UPSTREAM="github.com/snowballstem"
SNOWBALL_PATH="$(dirname "${0}")/${SNOWBALL_UPSTREAM}"

cd "$(dirname "${0}")"
BASE_DIR="$(pwd)"

mkdir -p build.d
cd build.d
if ! [ -d snowball ]; then
	git clone https://${SNOWBALL_UPSTREAM}/snowball.git
fi
cd snowball
make
mkdir -p go/algorithms
./snowball algorithms/english.sbl -go -o go/algorithms/english

cd "${BASE_DIR}"
mkdir -vp "${SNOWBALL_PATH}/snowball/"
mv build.d/snowball/go "${SNOWBALL_PATH}/snowball/"
