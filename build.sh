#!/bin/env bash

usage() {
	printf "Usage: $0 [windows][release][clean] \n\twindows: build for windows, leave empty for building for linux\n\trelease: build with release flags\n\tclean: clean the output directory\n"
	exit 1
}

project_name="snake_clone"
out_dir="out"
out_name="$project_name"
linker_flags=""

windows_user="ngkil"

windows_flag=false
release_flag=false
clean_flag=false

# get parameters and set the flags

while [ "$1" != "" ]; do
	case $1 in
	windows)
		windows_flag=true
		;;
	release)
		release_flag=true
		;;
	clean)
		clean_flag=true
		;;
	*)
		usage
		;;
	esac
	shift
done

if [ "$release_flag" == "true" ]; then
	out_name="$out_name-release"
	linker_flags="$linker_flags -s -w"
	if [ "$windows_flag" == "true" ]; then
		linker_flags="$linker_flags -H windowsgui"
	fi
fi

# check if parameter one exist and is equal to "windows"
if [ "$windows_flag" == "true" ]; then
	windows_out_dir="/mnt/c/Users/$windows_user/Desktop/$project_name"
	if [ ! -d "$windows_out_dir" ]; then
		mkdir $windows_out_dir
	fi
	out_name="$out_name.exe"

	if [ "$clean_flag" == "true" ]; then
		echo "Cleaning previous file"
		rm $out_dir/$out_name 2>/dev/null
	fi

	echo "Building for windows"
	set -x
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "$linker_flags" -o $out_dir/$out_name ./src
	mv $out_dir/$out_name $windows_out_dir
	if [ "$release_flag" == "false" ]; then
		printf "$out_name\npause\n" >$windows_out_dir/debug_run.bat
	fi
	set +x

else
	if [ "$clean_flag" == "true" ]; then
		echo "Cleaning previous file"
		rm $out_dir/$out_name 2>/dev/null
	fi

	echo "Building for linux"
	set -x
	go build -ldflags "$linker_flags" -o $out_dir/$out_name ./src
	set +x
fi

printf "Done Building\n---------------------------------\n\n\n"
