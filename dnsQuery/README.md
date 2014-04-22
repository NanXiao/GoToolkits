This project is used for DNS query.

Build:

	go build dnsquery.go

Usage:

	dnsquery [domain_name] [domain_name] ......

Example:

	dnsquery google.com
	
	output:
	The IP addresses of google.com are:
        74.125.128.113
        74.125.128.138
        74.125.128.139
        74.125.128.100
        74.125.128.101
        74.125.128.102
        2404:6800:4005:c00::64

	