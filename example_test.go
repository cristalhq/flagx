package flagx_test

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/cristalhq/flagx"
)

func ExampleFlagSet() {
	var d time.Duration
	fset := flagx.NewFlagSet("testing", os.Stderr)
	fset.Duration(&d, "timeout", "t", 10*time.Second, "just a timeout")

	err := fset.Parse([]string{"-t", "20s"})
	if err != nil {
		panic(err)
	}

	if d != 20*time.Second {
		panic(fmt.Sprintf("got %v want %v", d, 20*time.Second))
	}
}

func ExampleTextVar() {
	fs := flagx.NewFlagSet("ExampleTextVar", os.Stdout)

	var ip net.IP
	fs.Text(&ip, "ip", "", net.IPv4(192, 168, 0, 100), "`IP address` to parse")

	err := fs.Parse([]string{"-ip", "127.0.0.1"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("{ip: %v}\n\n", ip)

	// Output:
	// {ip: 127.0.0.1}
}
