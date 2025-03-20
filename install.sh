#!/usr/bin/env bash

# Project details
REPO_OWNER="jgfranco17"
REPO_NAME="jocasta"
DEFAULT_VERSION="latest"
INSTALL_PATH="${HOME}/.local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[1;32m'
BLUE='\033[1;36m'
NC='\033[0m'

print_error_message() {
  local message=$1
  echo -e "${RED}[ERROR] ${message}${NC}"
}

print_ok_message() {
  local message=$1
  echo -e "${GREEN}${message}${NC}"
}

get_latest_version() {
  curl --silent "https://api.github.com/repos/${REPO_OWNER}/${REPO_NAME}/releases/${DEFAULT_VERSION}" | \
  jq -r .tag_name
}

download_binary() {
  local version=$1
  local os=$2
  local arch=$3

  url="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${version}/jocasta-${version}-${os}-${arch}.tar.gz"

  curl -L "$url" -o jocasta.tar.gz || {
    print_error_message "Download failed; please check the version and try again."
    exit 1
  }
}

install_binary() {
  sudo tar -xzf jocasta.tar.gz -C "${INSTALL_PATH}" jocasta || {
    print_error_message "Installation failed."
    exit 1
  }
  rm jocasta.tar.gz
  print_ok_message "Jocasta installation complete!"
  print_ok_message "Installed at: ${INSTALL_PATH}"
}

# =============== MAIN SCRIPT ===============

version="${1:-$DEFAULT_VERSION}"

# Detect OS and architecture
case "$(uname -s)" in
  Linux*) os="linux" ;;
  Darwin*) os="darwin" ;;
  *) print_error_message "Unsupported OS"; exit 1 ;;
esac

arch="$(uname -m)"
case "$arch" in
  x86_64) arch="amd64" ;;
  aarch64) arch="arm64" ;;
  *) print_error_message "Unsupported architecture: $arch"; exit 1 ;;
esac

# Resolve latest version if needed
if [ "$version" = "latest" ]; then
  version=$(get_latest_version)
  if [ -z "$version" ]; then
    print_error_message "Unable to fetch the latest version."
    exit 1
  fi
fi

echo -e "Installing Jocasta version ${BLUE}${version}${NC} for ${BLUE}${os}/${arch}${NC}..."

# Download and install
download_binary "$version" "$os" "$arch"
install_binary

echo -e "You can now run ${BLUE}jocasta --version${NC} to verify the installation."
echo -e "To get started, run ${BLUE}jocasta --help${NC} to see available commands."
