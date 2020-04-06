#!/bin/sh

openbsdH="25a8dc77cda3d225c85d3f0cb318d01c17546c1c4a8c789a318f832ce1948bc3"

earlyCheck(){
    os=`uname`
    os=`echo $os | tr "[:upper:]" "[:lower:"]`

    case $os in
        *openbsd* ) ;;
        *)
            echo "Pre-built binary not available for your os"
            exit 1
            ;;
    esac

    cpu=`uname -m`
    cpu=`echo $cpu | tr "[:upper:]" "[:lower:"]`

    case $cpu in
        *amd*64* | *x86*64* ) ;;
        *)
            echo "Pre-built binary not available for your cpu"
            exit 1
            ;;
    esac
}

getURL(){
    url="https://archive.org/download/grus-v0.1.0/grus-v0.1.0-$os-$cpu"
}

printURL(){
    echo "You can get the Pre-built binary here:"
    echo "$url"
    echo
    echo "Run these commands to install it on your device."
    echo "# curl -L -o /usr/local/bin/grus $url"
    echo "# chmod +x /usr/local/bin/grus"
    echo
    echo "This is sha256 hash for grus built for: $os $cpu"
    case $os in
        *openbsd* )
            echo "$openbsdH"
            ;;
    esac
    echo
    echo "Verify the hash by running sha256 on grus binary."
    echo "$ sha256 /usr/local/bin/grus"
}

echo "Grus v0.1.0"
echo
earlyCheck
getURL
printURL
