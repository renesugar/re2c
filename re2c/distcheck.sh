#!/bin/bash

set -e

if [ $# -ne 1 ]
then
    echo "usage: ./distcheck.sh <builddir>"
    exit 1
fi
builddir=$1
srcdir=`pwd`

$srcdir/autogen.sh

# try to be portable on various MAKEs
for make_prog in make bmake
do
    rm -rf $builddir
    mkdir $builddir
    cd $builddir
        $srcdir/configure --enable-docs && \
        $make_prog bootstrap -j5
        $make_prog distcheck -j5
    cd $srcdir
done
