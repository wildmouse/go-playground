// Code generated by goa v3.0.6, DO NOT EDIT.
//
// files HTTP client CLI support package
//
// Command:
// $ goa gen github.com/wild-mouse/go-playground/design/goa_files

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	filesc "github.com/wild-mouse/go-playground/gen/http/files/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `files add
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` files add --body '{
      "file": "Neque ipsum et saepe."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		filesFlags = flag.NewFlagSet("files", flag.ContinueOnError)

		filesAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		filesAddBodyFlag = filesAddFlags.String("body", "REQUIRED", "")
	)
	filesFlags.Usage = filesUsage
	filesAddFlags.Usage = filesAddUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "files":
			svcf = filesFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "files":
			switch epn {
			case "add":
				epf = filesAddFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "files":
			c := filesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = filesc.BuildAddPayload(*filesAddBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// filesUsage displays the usage of the files command and its subcommands.
func filesUsage() {
	fmt.Fprintf(os.Stderr, `Service is the files service interface.
Usage:
    %s [globalflags] files COMMAND [flags]

COMMAND:
    add: Add implements add.

Additional help:
    %s files COMMAND --help
`, os.Args[0], os.Args[0])
}
func filesAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] files add -body JSON

Add implements add.
    -body JSON: 

Example:
    `+os.Args[0]+` files add --body '{
      "file": "Neque ipsum et saepe."
   }'
`, os.Args[0])
}
