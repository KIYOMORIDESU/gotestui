package view

import (
	"strings"

	"github.com/KIYOMORIDESU/gotestui/collector"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Color(ta collector.TestAction) tcell.Color {
	switch ta {
	case collector.FAIL:
		return tcell.ColorRed
	case collector.PASS:
		return tcell.ColorGreen
	case collector.SKIP:
		return tcell.ColorGray
	default:
		return tcell.ColorWhite
	}
}

func CreateApplication(tes []*collector.TestEventForView, results *collector.Results) *tview.Application {
	app := tview.NewApplication()

	flex := CreateTestCaseView(tes)
	app.SetRoot(flex, true)

	return app
}

func CreateTestCaseView(tes []*collector.TestEventForView) *tview.Flex {
	flex := tview.NewFlex()

	// right side flex
	logViewer := tview.NewTextView()
	logViewer.SetBorder(true).SetTitle("log")

	// left side flex
	leftFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	searchInputField := tview.NewInputField().SetLabel("TestCaseName:")
	searchInputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter: // Enterを入力したとき
			// textView.SetText(textView.GetText(true) + inputField.GetText() + "\n")
			// textViewに入力されている内容と、inputFieldに入力されている内容を取得し、textViewのテキストエリアに表示する

      searchWord := searchInputField.GetText()
      _ = searchWord
			return nil             // Enterをdefaultのキーアクションへは入力しない
		}
		return event // Enter以外はdefaultのキーアクションへ入力
	})
	searchFlex := tview.NewFlex()
	searchFlex.AddItem(searchInputField, 0, 1, false).SetBorder(true)

	leftFlex.AddItem(searchFlex, 5, 1, false)

	table := tview.NewTable()
	table.SetBorder(true).SetTitle("TestCases")
	for index, te := range tes {
		table.SetCell(index, 0, tview.NewTableCell(te.TestName).SetTextColor(Color(te.State)))
		table.SetCell(index, 1, tview.NewTableCell(string(te.State)).SetTextColor(Color(te.State)))
	}
	table.SetSelectable(true, false).Select(0, 0).SetFixed(1, 1)
	table.SetSelectedFunc(func(row, col int) {
		logViewer.SetText(strings.Join(tes[row].Outputs, ""))
	})
	leftFlex.AddItem(table, 0, 4, true)

	flex.AddItem(leftFlex, 0, 1, true)
	flex.AddItem(logViewer, 0, 1, false)
	return flex
}
