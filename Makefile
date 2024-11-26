VERSION := $(or $(AppVersion), "v1.1.2")
COMMIT := $(or $(shell git rev-parse --short HEAD), "unknown")
BUILDDATE := $(shell date +%Y-%m-%d)
PWD := $(shell pwd)

LDFLAGS := -X 'main.AppVersion=$(VERSION)' -X 'main.CommitHash=$(COMMIT)' -X 'main.BuildDate=$(BUILDDATE)'

PLATFORMS = linux darwin windows
ARCHITECTURES = amd64 arm64 arm

all: build

download:
	go mod tidy
	go mod download

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

run: build
	./build/stikky

install: build
	cp stikky /usr/local/bin/stikky
	cp man/stikky.1 /usr/local/share/man/man1/stikky.1

build:
	GOOS=darwin go build -ldflags="$(LDFLAGS)" -o build/stikky .

# Create distribution packaging
dist: pkg_macos pkg_windows

# Packaging MacOS
pkg_macos: pkg_macos_amd64 pkg_macos_arm64

# Build process for MacOS amd64
build_macos_amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS) -H=darwin" -o build/macos/stikky .

pkg_macos_amd64: build_macos_amd64
	mkdir -p build/macos/stikky.app/Contents/MacOS
	mkdir -p build/macos/stikky.app/Contents/Resources
	cp Info.plist build/macos/stikky.app/Contents/
	cp build/macos/stikky build/macos/stikky.app/Contents/MacOS/
	cp res/assets/icon.icns build/macos/stikky.app/Contents/Resources/
	chmod +x build/macos/stikky.app
	rm -rf build/macos/stikky
	ln -s /Applications $(PWD)/build/macos/Applications
	hdiutil create -volname "Stikky Installer" -srcfolder build/macos/ -ov -format UDZO build/Stikky_amd64.dmg
	rm -rf build/macos


# Build process for MacOS arm64
build_macos_arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS) -H=darwin" -o build/macos/stikky .

pkg_macos_arm64: build_macos_arm64
	mkdir -p build/macos/stikky.app/Contents/MacOS
	mkdir -p build/macos/stikky.app/Contents/Resources
	cp Info.plist build/macos/stikky.app/Contents/
	cp build/macos/stikky build/macos/stikky.app/Contents/MacOS/
	cp res/assets/icon.icns build/macos/stikky.app/Contents/Resources/
	chmod +x build/macos/stikky.app
	rm -rf build/macos/stikky
	ln -s /Applications $(PWD)/build/macos/Applications
	hdiutil create -volname "Stikky Installer" -srcfolder build/macos/ -ov -format UDZO build/Stikky_arm64.dmg
	rm -rf build/macos

# Packaging Windows
build_windows:
	GOOS=windows go build -ldflags="$(LDFLAGS) -H=windowsgui" -o build/stikky.exe .

# Packaging windows amd64
pkg_windows: build_windows
	tar -zcvf build/stikky-windows.tar.gz build/stikky.exe
	rm -rf build/stikky.exe

# Cleaning up
clean:
	rm -rf stikky build

# Developer dependent work
generate_icns:
	mkdir -p res/assets/icon.iconset
	cp res/assets/icon.png res/assets/icon.iconset/icon_512x512.png
	magick res/assets/icon.png -resize 16x16 res/assets/icon.iconset/icon_16x16.png
	magick res/assets/icon.png -resize 32x32 res/assets/icon.iconset/icon_32x32.png
	magick res/assets/icon.png -resize 48x48 res/assets/icon.iconset/icon_48x48.png
	magick res/assets/icon.png -resize 64x64 res/assets/icon.iconset/icon_64x64.png
	magick res/assets/icon.png -resize 128x128 res/assets/icon.iconset/icon_128x128.png
	magick res/assets/icon.png -resize 256x256 res/assets/icon.iconset/icon_256x256.png
	iconutil -c icns res/assets/icon.iconset/
	rm -rf res/assets/icon.iconset