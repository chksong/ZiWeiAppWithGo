package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type DrawDlg struct {
	widgets.QWidget

	_ func() `constructor:"init"`
	_ func() `slot:"animate"`


	background  *gui.QBrush
	circleBrush *gui.QBrush
	textFont    *gui.QFont
	circlePen   *gui.QPen
	textPen     *gui.QPen

	elapsed int
}

func (dlg *DrawDlg) init()  {
	dlg.SetFixedSize2(400,400)
	dlg.ConnectPaintEvent(dlg.paintEvent)


	gradient := gui.NewQLinearGradient2(core.NewQPointF3(50, -20), core.NewQPointF3(80, 20))
	gradient.SetColorAt(0, gui.NewQColor2(core.Qt__white))
	gradient.SetColorAt(1, gui.NewQColor3(0xa6, 0xce, 0x39, 255))

	dlg.background = gui.NewQBrush3(gui.NewQColor3(64, 32, 64, 255), 1)
	dlg.circleBrush = gui.NewQBrush10(gradient)
	dlg.circlePen = gui.NewQPen3(gui.NewQColor2(core.Qt__black))
	dlg.circlePen.SetWidth(1)
	dlg.textPen = gui.NewQPen3(gui.NewQColor2(core.Qt__white))
	dlg.textFont = gui.NewQFont()
	dlg.textFont.SetPixelSize(50)


	dlg.ConnectAnimate(dlg.animate)
	dlg.ConnectPaintEvent(dlg.paintEvent)
}

func (dlg *DrawDlg) animate()  {
	dlg.elapsed = (dlg.elapsed + 50) % 1000
	dlg.Update()
}

func (dlg *DrawDlg) paintEvent( event *gui.QPaintEvent)  {
	painter := gui.NewQPainter2(dlg)
	painter.SetRenderHint(gui.QPainter__Antialiasing,true)


	painter.FillRect3(event.Rect(), dlg.background)
	painter.Translate3(100, 100)

	painter.Save()
	painter.SetBrush(dlg.circleBrush)
	painter.SetPen(dlg.circlePen)
	painter.Rotate(float64(dlg.elapsed) * float64(0.03))

	r := float64(dlg.elapsed) / 1000
	n := float64(30)
	for i := 0; i < int(n); {
		i++
		painter.Rotate(30)
		factor := (float64(i) + r) / n
		radius := 0 + 120*factor
		circleRadius := 1 + factor*20
		painter.DrawEllipse(core.NewQRectF4(radius, -circleRadius, circleRadius*2, circleRadius*2))
	}
	painter.Restore()

	painter.SetPen(dlg.textPen)
	painter.SetFont(dlg.textFont)
//	painter.DrawText6(core.NewQRect4(-50, -50, 100, 100), int(core.Qt__AlignCenter), "Qt", core.NewQRect())

	painter.DestroyQPainter()
}