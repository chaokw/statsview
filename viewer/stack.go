package viewer

import (
	"encoding/json"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	// VCStack is the name of StackViewer
	VCStack = "stack memory"
	VOFFHEAP = "off heap memory"
)

// StackViewer collects the stack-stats metrics via `runtime.ReadMemStats()`
type StackViewer struct {
	smgr  *StatsMgr
	graph *charts.Line
}

// NewStackViewer returns the StackViewer instance
// Series: StackSys / StackInuse / MSpanSys / MSpanInuse
func NewStackViewer() Viewer {
	graph := newBasicView(VCStack)
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Stack Memory"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Size", AxisLabel: &opts.AxisLabel{Formatter: "{value} MB"}}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)
	graph.AddSeries("StackSys", []opts.LineData{}).
		AddSeries("StackInuse", []opts.LineData{}).
		SetSeriesOptions(
			charts.WithLabelOpts(
				opts.Label{
					Show: true,
				}),
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.8,
				}),
			charts.WithLineChartOpts(
				opts.LineChart{
					Stack: "stack",
				}),
		)

	return &StackViewer{graph: graph}
}

func (vr *StackViewer) SetStatsMgr(smgr *StatsMgr) {
	vr.smgr = smgr
}

func (vr *StackViewer) Name() string {
	return VCStack
}

func (vr *StackViewer) View() *charts.Line {
	return vr.graph
}

func (vr *StackViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	vr.smgr.Tick()

	metrics := Metrics{
		Values: []float64{
			fixedPrecision(float64(memstats.Stats.StackSys)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.StackInuse)/1024/1024, 2),
		},
		Time: memstats.T,
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}



type OffHeapViewer struct {
	smgr  *StatsMgr
	graph *charts.Line
}

func NewOffHeapViewer() Viewer {
	graph := newBasicView(VOFFHEAP)
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "OFF HEAP Memory"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Size", AxisLabel: &opts.AxisLabel{Formatter: "{value} MB"}}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)
	graph.AddSeries("MCacheInuse", []opts.LineData{}).
		AddSeries("MCacheSys", []opts.LineData{}).
		AddSeries("MSpanSys", []opts.LineData{}).
		AddSeries("MSpanInuse", []opts.LineData{}).
		SetSeriesOptions(
			charts.WithLabelOpts(
				opts.Label{
					Show: true,
				}),
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.8,
				}),
			charts.WithLineChartOpts(
				opts.LineChart{
					Stack: "stack",
				}),
		)

	return &OffHeapViewer{graph: graph}
}

func (vr *OffHeapViewer) SetStatsMgr(smgr *StatsMgr) {
	vr.smgr = smgr
}

func (vr *OffHeapViewer) Name() string {
	return VOFFHEAP
}

func (vr *OffHeapViewer) View() *charts.Line {
	return vr.graph
}

func (vr *OffHeapViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	vr.smgr.Tick()

	metrics := Metrics{
		Values: []float64{
			fixedPrecision(float64(memstats.Stats.MSpanInuse)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.MCacheSys)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.MSpanSys)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.MSpanInuse)/1024/1024, 2),
		},
		Time: memstats.T,
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}
