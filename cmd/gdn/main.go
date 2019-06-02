package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"code.cloudfoundry.org/guardian/guardiancmd"
	"github.com/jessevdk/go-flags"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

// version is overwritten at compile time by passing
// -ldflags -X main.version=<version>
var version = "dev"

func main() {
	configFilePath := flag.String("config", "", "config file path")
	printVersion := flag.Bool("v", false, "print version")
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	cmd := &guardiancmd.GdnCommand{}

	parser := flags.NewParser(cmd, flags.Default)
	parser.NamespaceDelimiter = "-"

	if *configFilePath != "" {
		iniParser := flags.NewIniParser(parser)
		must(iniParser.ParseFile(*configFilePath))
	}

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		"gdn",
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	_, err = parser.Parse()
	mustNot(err)
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var mustNot = must
