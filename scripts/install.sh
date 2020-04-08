#!/bin/sh

freebsdH="f77989e2d243dc32ff1ab5aa91b48e2abe72a3daba86b781e7905fd4d6a9bd1a"
openbsdH="0251e6cecb77aa03ac02d4d67ad9d4cd164d199e1a8a3ef4dd8b711fa43a7298"
linuxH="99c35fd048d02fc7136271a7ab2cbe582810d426e2b2bf9687ede6f2b8011960"
netbsdH="f686dae688db4dd1a4733ae0be06b588d47f5d8fac6e23dc11f8175d5dcff53c"
dragonflyH="c7dbf94b379f9b85956cf09fa015e6eab4ee40dd60fe747cde1c2f15cb172ad2"
darwinH="81db0127149b96558ed1d0e556ebe1a653a2e38a0f212fc23257bb9bf0ab2ead"

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
    url="https://archive.org/download/grus-v0.3.0/grus-v0.3.0-$os-$cpu"
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
    echo
}

echo "Grus v0.3.0"
echo
earlyCheck
getURL
printURL
