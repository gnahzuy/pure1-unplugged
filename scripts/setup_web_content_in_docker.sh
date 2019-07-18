#!/usr/bin/env bash
# Sets up the the pure1-unplugged web content in Docker

set -ex
BASIC_SCRIPT_DIR="$(dirname "$0")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Gets the current absolute directory (bash wizardry)

BUILDER_IMG="pure-angular-builder:1.0"
REPO_PATH="/pure1-unplugged"

${BASIC_SCRIPT_DIR}/print_console_label.sh "Installing Web Content Dependencies in Docker"
docker run -v "${SCRIPT_DIR}/../":"${REPO_PATH}/" -u $(id -u):$(id -g)  --net=host -w "${REPO_PATH}/" ${BUILDER_IMG} "${REPO_PATH}/scripts/setup_web_content.sh"
