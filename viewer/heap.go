package viewer

import (
	"encoding/json"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	// VHeap is the name of HeapViewer
	VHeap    = "heap memory"
	VHeapOBJ = "heap objects"
)

// HeapViewer collects the heap-stats metrics via `runtime.ReadMemStats()`
type HeapViewer struct {
	smgr  *StatsMgr
	graph *charts.Line
}

type HeapObjectsViewer struct {
	smgr  *StatsMgr
	graph *charts.Line
}

// NewHeapViewer returns the HeapViewer instance
// Series: Alloc / Inuse / Sys / Idle
func NewHeapViewer() Viewer {
	graph := newBasicView(VHeap)
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Heap Memory"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Size", AxisLabel: &opts.AxisLabel{Formatter: "{value} MB"}}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)
	graph.AddSeries("HeapAlloc", []opts.LineData{}).
		AddSeries("HeapInuse", []opts.LineData{}).
		AddSeries("HeapSys", []opts.LineData{}).
		AddSeries("HeapIdle", []opts.LineData{}).
		AddSeries("HeapReleased", []opts.LineData{}).
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

	return &HeapViewer{graph: graph}
}

func NewHeapObjectsViewer() Viewer {
	graph := newBasicView(VHeapOBJ)
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Heap Objects"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Number"}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)
	graph.AddSeries("HeapObjects", []opts.LineData{}).
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

	return &HeapObjectsViewer{graph: graph}
}

func (vr *HeapViewer) SetStatsMgr(smgr *StatsMgr) {
	vr.smgr = smgr
}
func (vr *HeapViewer) Name() string {
	return VHeap
}

func (vr *HeapViewer) View() *charts.Line {
	return vr.graph
}

func (vr *HeapViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	vr.smgr.Tick()

	metrics := Metrics{
		Values: []float64{
			fixedPrecision(float64(memstats.Stats.HeapAlloc)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.HeapInuse)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.HeapSys)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.HeapIdle)/1024/1024, 2),
			fixedPrecision(float64(memstats.Stats.HeapReleased)/1024/1024, 2),
		},
		Time: memstats.T,
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}

func (vr *HeapObjectsViewer) SetStatsMgr(smgr *StatsMgr) {
	vr.smgr = smgr
}
func (vr *HeapObjectsViewer) Name() string {
	return VHeapOBJ
}

func (vr *HeapObjectsViewer) View() *charts.Line {
	return vr.graph
}

func (vr *HeapObjectsViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	vr.smgr.Tick()

	metrics := Metrics{
		Values: []float64{
			float64(memstats.Stats.HeapObjects),
		},
		Time: memstats.T,
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}
