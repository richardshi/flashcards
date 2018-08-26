package views

import (
	"github.com/gizak/termui"

	"github.com/richardshi/flashcards/components"
	"github.com/richardshi/flashcards/config"
)

const optionsHeight = 3

// View holds the application components
type View struct {
	Flashcard  *components.Flashcard
	Options    *components.Options
	WordList   *components.WordList
	flashcards []config.Flashcard
}

// CreateFlashcardView creates a new application view
func CreateFlashcardView(config *config.Config) *View {
	flashcard := components.CreateFlashcard(&config.Flashcards[0], optionsHeight)
	options := components.CreateOptions(optionsHeight)
	wordList := components.CreateWordList(&config.Flashcards, optionsHeight)

	view := &View{
		Options:    options,
		Flashcard:  flashcard,
		WordList:   wordList,
		flashcards: config.Flashcards,
	}

	return view
}

// UpdateFlashcard sets the text and definition to the currently selected flashcard
func (v *View) UpdateFlashcard() {
	v.Flashcard.Change(&v.flashcards[v.WordList.Offset+v.WordList.SelectedWord-1])
}

// Refresh re-renders the components
func (v *View) Refresh() {
	termui.Render(
		v.Flashcard,
		v.WordList,
		v.Options,
	)
}
