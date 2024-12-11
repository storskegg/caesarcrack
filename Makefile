EXE=caesarcrack

default: all

all: clean build

clean:
	[[ -d "./bin" ]]; rm -rf ./bin

build:
	go build -ldflags "-w -s" -a -o ./bin/$(EXE)

deflate:
	upx -9 -k ./bin/$(EXE)
