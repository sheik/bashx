# bashx
A unix shell that extends bash. The goal is to have a new shell that can be backwards compatible with bash, but also provide new features such as providing channeling mechanisms, new IO redirectiont tecniques, debug modes, windowing, and hopefully serve as a central "IDE" for programmers and power users.

It is currently *very early* in development but can currently handle simple piped commands

    bashx$ ls / | grep -i i
    bin
    initrd.img
    initrd.img.old
    lib
    lib32
    lib64
    media
    sbin
    swapfile
    vmlinuz
    vmlinuz.old
    bashx$ head -c 20 /dev/urandom | base64
    +yGFhqlX+2lnuRrfYVOGs+eqEeg=

Parsing is very rudementary right now but the basics are working.
