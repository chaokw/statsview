# 🚀 Statsview

Statsview is a real-time Golang runtime stats visualization profiler. It is built top on another open-source project, [go-echarts](https://github.com/go-echarts/go-echarts), which helps statsview to show its graphs on the browser.

<a href="https://github.com/go-echarts/statsview/pulls">
    <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome">
</a>
<a href="https://goreportcard.com/report/github.com/go-echarts/statsview">
    <img src="https://goreportcard.com/badge/github.com/go-echarts/statsview" alt="Go Report Card">
</a>
<a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="MIT License">
</a>
<a href="https://pkg.go.dev/github.com/go-echarts/statsview">
    <img src="https://godoc.org/github.com/go-echarts/statsview?status.svg" alt="GoDoc">
 </a>

## 🔰 Installation

```shell
$ go get -u github.com/chaokw/statsview/...
```

## 📝 Usage

Statsview is quite simple to use and all static assets have been packaged into the project which makes it possible to run offline. It's worth pointing out that statsview has integrated the standard `net/http/pprof` hence statsview will be the only profiler you need.

```golang
package main

import (
    "time"
    "github.com/go-echarts/statsview"
)

func main() {
	s := statsview.New()

	// Start() runs a HTTP server at `localhost:60023` by default.
	go s.Start()

	// Stop() will shutdown the http server gracefully
	// s.Stop()

	// busy working....
	time.Sleep(time.Minute)
}

// Visit your browser at http://localhost:60023
// Or debug as always via http://localhost:60023/debug/pprof, http://localhost:60023/debug/pprof/heap, ...
```

## ⚙️ Configuration

Statsview gets a variety of configurations for the users. Everyone could customize their favorite charts style.

```golang
// WithInterval sets the interval(in Millisecond) of collecting and pulling metrics
// default -> 2000
WithInterval(interval int)

// WithMaxPoints sets the maximum points of each chart series
// default -> 30000
WithMaxPoints(n int)

// WithTemplate sets the rendered template which fetching stats from the server and
// handling the metrics data
WithTemplate(t string)

// WithAddr sets the listening address and link address
// default -> "localhost:60023"
WithAddr(addr string)

// WithLinkAddr sets the html link address
// default -> "localhost:60023"
WithLinkAddr(addr string)

// WithTimeFormat sets the time format for the line-chart Y-axis label
// default -> "15:04:05"
WithTimeFormat(s string)

// WithTheme sets the theme of the charts
// default -> Westeros
//
// Optional:
// * ThemeWesteros
// * ThemeMacarons
WithTheme(theme Theme)
```

#### Set the options

```golang
import (
    "github.com/chaokw/statsview"
    "github.com/chaokw/statsview/viewer"
)

// set configurations before calling `statsview.New()` method
viewer.SetConfiguration(viewer.WithAddr(":60023"), viewer.WithLinkAddr("10.182.105.147:60023"))

s := statsview.New()
go s.Start()
```

## 🗂 Viewers

Viewer is the abstraction of a Graph which in charge of collecting metrics from Runtime. Statsview provides some default viewers as below.

* `HeapViewer`
* `HeapObjectsViewer`
* `StackViewer`
* `OffHeapViewer`
* `GoroutinesViewer`
* `GCNumViewer`
* `GCSizeViewer`
* `GCCPUFractionViewer`

Viewer wraps a go-echarts [*charts.Line](https://github.com/go-echarts/go-echarts/blob/master/charts/line.go) instance that means all options/features on it could be used. To be honest, I think that is the most charming thing about this project.

## 🔖 Snapshot

![Macarons](https://github.com/chaokw/statsview/blob/master/images/statsview1.png)

![Macarons](https://github.com/chaokw/statsview/blob/master/images/statsview2.png)


## 📄 License

MIT [©chenjiandongx](https://github.com/chenjiandongx)
