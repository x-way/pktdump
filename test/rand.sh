#!/bin/bash

TMPFILE="/tmp/$0.$$"
tcprewrite --seed $RANDOM --enet-mac-seed $RANDOM --enet-mac-seed-keep-bytes 3 -i "$1" -o "$TMPFILE"
cp "$TMPFILE" "$1"
rm -f "$TMPFILE"
