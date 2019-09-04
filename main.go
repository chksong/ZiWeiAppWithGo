package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
	"runtime"
)

var (
	canvas *widgets.QWidget
	scene  *widgets.QGraphicsScene
	view   *widgets.QGraphicsView
)

func resizeEvent(e *gui.QResizeEvent) {
	view.FitInView(scene.ItemsBoundingRect(), core.Qt__KeepAspectRatio)
}

func main()  {
	widgets.NewQApplication(len(os.Args),os.Args)
//	var window = widgets.NewQMainWindow(nil,0)
//	window.SetWindowTitle("test")
//

	//var rect = core.NewQRect4(0,0,680,480)
	//window.SetGeometry(rect)
	//window.SetMinimumSize2(640,480)
	//
	//scene  = widgets.NewQGraphicsScene(nil)
	//view   = widgets.NewQGraphicsView(nil)
	//view.ConnectResizeEvent(resizeEvent)
	//
	//var color = gui.NewQColor3(255, 0, 255, 255)
	//var pen   = gui.NewQPen3(color)
	//
	//scene.AddLine2(0,0,scene.Width()/2,scene.Height()/2,pen)
	//view.SetScene(scene)
	//view.Show()


	var dlg  = NewDrawDlg(nil,0)
	timer := core.NewQTimer(nil)
	timer.ConnectTimeout(dlg.Animate)
	timer.Start(20)
	dlg.Show()




//	window.SetCentralWidget(view)
//	window.Show()
	widgets.QApplication_Exec()
}

func NewCanVas()  *widgets.QWidget {
	canvas = widgets.NewQWidget(nil,0)
	canvas.SetMinimumSize2(640,480)


	scene =  widgets.NewQGraphicsScene(nil)
	view = widgets.NewQGraphicsView(nil)


	var font = gui.NewQFont2("Meiryo", 20, 2, false)
	if runtime.GOARCH == "js" || runtime.GOARCH == "wasm" {
		scene.AddText("Hello World", font)
	} else {
		scene.AddText("Hello 世界hhhhhhhhh", font)
	}

	var color = gui.NewQColor3(255, 0, 255, 255)
	var pen = gui.NewQPen3(color)

	scene.AddLine2(0,3,scene.Width(),3,pen)
	view.SetScene(scene)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(view,0,0)

	canvas.SetLayout(layout)
	return canvas
}