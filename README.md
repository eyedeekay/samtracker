samtracker
==========


Oh neat, that was easy.

This is an HTTP Bittorrent tracker with automatic port-forwarding to I2P, based
on [vvampirius/retracker](https://github.com/vvampirius/retracker). It is very
minimal, memory safe, and doesn't really do much other than let people announce
torrents, which is all it's supposed to do.

Installation
------------

If you have a go environment set up, you can easily install it from source using

        go get -u github.com/eyedeekay/samtracker

Or, if you want to use a binary, you can download the tar.gz file:

        wget https://github.com/eyedeekay/samtracker/releases/download/0.0.01/samtracker.tar.gz
        mkdir tmp && cd tmp
        tar xvf ../samtracker.tar.gz
        sudo make install
