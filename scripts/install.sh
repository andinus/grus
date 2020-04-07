#!/bin/sh

freebsdH="3393230e78fd495dad4193dfdc7aafab11c007de058790aa0dbfd2e7ddd84689"
openbsdH="01b9fddd004b588a07f8fbec756cfc2cc7c07a51d765969e7a79af34dc93bee5"
linuxH="b12957117a839a741c85401ca0305db9d48b623c6e8754b81002cb3c588d9bf3"
netbsdH="67e0efe184b0327700ba67ddc662ba3174a7484de7d8223581367d3fdc21e5b8"
dragonflyH="d7501082472fd0b4b6686d2ec6667bf6965087e5e4fe8c8b04372123e51df5b3"
darwinH="4b96dbd72f0816fa9e89a9c2a249df4a7cb3a043d5a87fb8e6d948f17846b56d"

earlyCheck(){
    os=`uname`
    os=`echo $os | tr "[:upper:]" "[:lower:"]`

    case $os in
	# Not sure about uname output on DragonFly BSD.
        *openbsd* | *linux* | *freebsd* | *netbsd* | *dragonfly* | *dragonflybsd* | *darwin* ) ;;
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
    url="https://archive.org/download/grus-v0.2.0/grus-v0.2.0-$os-$cpu"
}

printURL(){
    echo "You can get the pre-built binary here:"
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
	*netbsd* )
            echo "$netbsdH"
            ;;
	*dragonflybsd* | *dragonfly* )
            echo "$dragonflyH"
            ;;
	*darwin* )
            echo "$darwinH"
            ;;
        *freebsd* )
            echo "$freebsdH"
            ;;
        *linux* )
            echo "$linuxH"
            ;;
    esac
    echo
    echo "Verify the hash by running sha256 on grus binary."
    echo "$ sha256 /usr/local/bin/grus"
}

echo "Grus v0.2.0"
echo
earlyCheck
getURL
printURL
