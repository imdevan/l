# Installation

## Homebrew
```bash
brew install imdevan/l/l
```

## Arch (AUR)
```bash
yay -S l
```

## GitHub Release

Download the latest binary for your platform from the [releases page](https://github.com/imdevan/l/releases).

```bash
# Linux (amd64)
curl -L https://github.com/imdevan/l/releases/latest/download/l-linux-amd64.tar.gz | tar -xz
sudo mv l-linux-amd64 /usr/local/bin/l
```

```bash
# macOS (Apple Silicon)
curl -L https://github.com/imdevan/l/releases/latest/download/l-darwin-arm64.tar.gz | tar -xz
sudo mv l-darwin-arm64 /usr/local/bin/l
```

```bash
# macOS (Intel)
curl -L https://github.com/imdevan/l/releases/latest/download/l-darwin-amd64.tar.gz | tar -xz
sudo mv l-darwin-amd64 /usr/local/bin/l
```

## Manual
```bash
gh repo clone imdevan/l
cd l
just build
sudo just install
```
