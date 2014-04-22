This project is used for testing whether the remote TCP port is open or not.

Build:

	go build tcp_keep_alive.go

Usage:

	Usage: tcp_keep_alive ip_addr port [interval(seconds, default:0)] [count(run times, default:1)]

Example:

	tcp_keep_alive google.com 80  

    tcp_keep_alive google.com 80 1 10
	
	

	