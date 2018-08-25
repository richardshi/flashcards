package components

import (
	"github.com/gizak/termui"
	"github.com/richardshi/flashcards/config"
)

// WordList is the definition of WordList component
type WordList struct {
	List         *termui.List
	Offset       int // starting point for scrolling
	SelectedWord int // position of the highlighted word
}

// CreateWordList is the constructor for the WordList component
func CreateWordList(flashcards *[]config.Flashcard, optionsHeight int) *WordList {
	wordList := getNames(*flashcards)
	list := termui.NewList()
	list.Items = wordList
	list.BorderLabel = "Flashcards"
	list.Height = termui.TermHeight() - optionsHeight
	return &WordList{
		List:         list,
		SelectedWord: list.InnerBounds().Min.Y,
		Offset:       0,
	}
}

// Buffer implements interface termui.Bufferer
func (w *WordList) Buffer() termui.Buffer {
	buf := w.List.Buffer()
	for i, item := range w.List.Items[w.Offset:] {

		y := w.List.InnerBounds().Min.Y + i

		if y > w.List.InnerBounds().Max.Y-1 {
			break
		}

		var cells []termui.Cell
		if y == w.SelectedWord {
			cells = termui.DefaultTxBuilder.Build(
				item, w.List.ItemBgColor, w.List.ItemFgColor)
		} else {
			cells = termui.DefaultTxBuilder.Build(
				item, w.List.ItemFgColor, w.List.ItemBgColor)
		}

		cells = termui.DTrimTxCls(cells, w.List.InnerWidth())

		x := 0
		for _, cell := range cells {
			width := cell.Width()
			buf.Set(w.List.InnerBounds().Min.X+x, y, cell)
			x += width
		}

		// When not at the end of the pane fill it up empty characters
		for x < w.List.InnerBounds().Max.X-1 {
			if y == w.SelectedWord {
				buf.Set(x+1, y,
					termui.Cell{
						Ch: ' ',
						Fg: w.List.ItemBgColor,
						Bg: w.List.ItemFgColor,
					},
				)
			} else {
				buf.Set(
					x+1, y,
					termui.Cell{
						Ch: ' ',
						Fg: w.List.ItemFgColor,
						Bg: w.List.ItemBgColor,
					},
				)
			}
			x++
		}
	}
	return buf
}

// ScrollUp enables us to scroll through the word list when it overflows
func (w *WordList) ScrollUp() (change bool) {
	if w.SelectedWord != w.List.InnerBounds().Min.Y {
		w.SelectedWord--
		change = true
	} else {
		if w.Offset > 0 {
			w.Offset--
			change = true
		}
	}
	return
}

// ScrollDown enables us to scroll through the word list when it overflows
func (w *WordList) ScrollDown() (change bool) {
	if w.SelectedWord < len(w.List.Items) {
		if w.SelectedWord == w.List.InnerBounds().Max.Y-1 {
			if w.Offset < len(w.List.Items)-1 {
				w.Offset++
				change = true
			}
		} else {
			w.SelectedWord++
			change = true
		}
	}
	return
}

// ScrollTop will move the cursor to the top of the list
func (w *WordList) ScrollTop() {
	w.SelectedWord = w.List.InnerBounds().Min.Y
	w.Offset = 0
}

// ScrollBottom will move the cursor to the bottom of the list
func (w *WordList) ScrollBottom() {
	w.SelectedWord = (len(w.List.Items))

	offset := len(w.List.Items) - (w.List.InnerBounds().Max.Y - 1)
	w.Offset = offset

	if offset < 0 {
		w.Offset = 0
	}
}

// GetHeight implements interface termui.GridBufferer
func (w *WordList) GetHeight() int {
	return w.List.Block.GetHeight()
}

// SetWidth implements interface termui.GridBufferer
func (w *WordList) SetWidth(width int) {
	w.List.SetWidth(width)
}

// SetX implements interface termui.GridBufferer
func (w *WordList) SetX(x int) {
	w.List.SetX(x)
}

// SetY implements interface termui.GridBufferer
func (w *WordList) SetY(y int) {
	w.List.SetY(y)
}

func getNames(wordList []config.Flashcard) []string {
	words := make([]string, 0)
	for i := 0; i < len(wordList); i++ {
		words = append(words, wordList[i].Name)
	}
	return words
}
