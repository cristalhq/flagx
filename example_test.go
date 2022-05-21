package flagx_test

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/cristalhq/flagx"
)

func ExampleFlagSet() {
	args := []string{"-t", "20s"}

	var d time.Duration
	fset := flagx.NewFlagSet("testing", os.Stderr)
	fset.Duration(&d, "timeout", "t", 10*time.Second, "just a timeout")

	err := fset.Parse(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(d)

	// Output: 20s
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
