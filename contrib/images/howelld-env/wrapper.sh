#!/usr/bin/env sh

BINARY=/howelld/${BINARY:-howelld}
ID=${ID:-0}
LOG=${LOG:-howelld.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'howelld'"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export howelldHOME="/howelld/node${ID}/howelld"

if [ -d "$(dirname "${howelldHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${howelldHOME}" "$@" | tee "${howelldHOME}/${LOG}"
else
  "${BINARY}" --home "${howelldHOME}" "$@"
fi
