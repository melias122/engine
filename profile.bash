#!/bin/sh
/usr/bin/time -f '%Uu %Ss %er %MkB %C' go test -profile -run=Profile2080 -cpuprofile cpu.pprof
