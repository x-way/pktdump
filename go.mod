module github.com/x-way/pktdump

go 1.14

require (
	github.com/florianl/go-nflog/v2 v2.0.1-0.20201019183757-4cdb5aa8b03d
	github.com/google/gopacket v1.1.19
	github.com/x-way/iptables-tracer v0.0.0-20201029065626-a73947133ca7
	golang.org/x/sys v0.0.0-20201022201747-fb209a7c41cd
)

replace github.com/florianl/go-nflog/v2 => ./go-nflog
