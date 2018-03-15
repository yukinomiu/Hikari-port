#!/bin/bash
home=/Users/yukinomiu/go/src/hikari-port
version=`cat $home/script/version.txt`
app_home=$home/command
target=$home/target
config=$home/config

# clean
echo 'clean target directory...'
rm $target/*

function build() {
    CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -ldflags "-s -w" -o $target/$3
}

# darwin
echo 'build darwin executable files...'
cd $app_home
build darwin amd64 hikari-port-darwin-x64-$version
build darwin 386 hikari-port-darwin-x86-$version

# linux
echo 'build linux executable files...'
cd $app_home
build linux amd64 hikari-port-linux-x64-$version
build linux 386 hikari-port-linux-x86-$version

# windows
echo 'build windows executable files...'
cd $app_home
build windows amd64 hikari-port-windows-x64-$version.exe
build windows 386 hikari-port-windows-x86-$version.exe

# copy config
echo 'copy config files...'
cp $config/* $target/

echo 'finished'
