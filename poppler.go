// Package poppler implements routines based on the
// poppler libraries as pulled from xpdf toolchain.
package poppler

/*
#cgo pkg-config: --cflags-only-I poppler-glib
#include <stdio.h>
#include <stdlib.h>
#include <poppler.h>
*/
import "C"

// Error codes returned by Poppler Document
type Error int

const (
	ErrorInvalid    Error = C.POPPLER_ERROR_INVALID     // Generic error when a document operation fails
	ErrorEncrypted        = C.POPPLER_ERROR_ENCRYPTED   // Document is encrypted
	ErrorOpenFile         = C.POPPLER_ERROR_OPEN_FILE   // File could not be opened for writing when saving document
	ErrorBadCatalog       = C.POPPLER_ERROR_BAD_CATALOG // Failed to read the document catalog
	ErrorDamaged          = C.POPPLER_ERROR_DAMAGED     // Document is damaged
)

type Orientation int

const (
	OrientationPortrait   Orientation = C.POPPLER_ORIENTATION_PORTRAIT
	OrientationLandscape              = C.POPPLER_ORIENTATION_LANDSCAPE
	OrientationUpsidedown             = C.POPPLER_ORIENTATION_UPSIDEDOWN
	OrientationSeascape               = C.POPPLER_ORIENTATION_SEASCAPE
)

// PageTransition types
type PageTransition int

const (
	PageTransitionReplace  PageTransition = C.POPPLER_PAGE_TRANSITION_REPLACE  // The new page replace the old one
	PageTransitionSplit                   = C.POPPLER_PAGE_TRANSITION_SPLIT    // Two lines sweep across the screen, revealing the new page
	PageTransitionBlinds                  = C.POPPLER_PAGE_TRANSITION_BLINDS   // Multiple lines, evenly spaced across the screen, synchronously sweep in the same direction to reveal the new page
	PageTransitionBox                     = C.POPPLER_PAGE_TRANSITION_BOX      // A rectangular box sweeps inward from the edges of the page or outward from the center revealing the new page
	PageTransitionWipe                    = C.POPPLER_PAGE_TRANSITION_WIPE     // A single line sweeps across the screen from one edge to the other revealing the new page
	PageTransitionDissolve                = C.POPPLER_PAGE_TRANSITION_DISSOLVE // The old page dissolves gradually to reveal the new one
	PageTransitionGlitter                 = C.POPPLER_PAGE_TRANSITION_GLITTER  // Similar to #POPPLER_PAGE_TRANSITION_DISSOLVE, except that the effect sweeps across the page in a wide band moving from one side of the screen to the other
	PageTransitionFly                     = C.POPPLER_PAGE_TRANSITION_FLY      // Changes are flown out or in to or from a location that is offscreen
	PageTransitionPush                    = C.POPPLER_PAGE_TRANSITION_PUSH     // The old page slides off the screen while the new page slides in
	PageTransitionCover                   = C.POPPLER_PAGE_TRANSITION_COVER    // The new page slides on to the screen covering the old page
	PageTransitionUncover                 = C.POPPLER_PAGE_TRANSITION_UNCOVER  // The old page slides off the screen uncovering the new page
	PageTransitionFade                    = C.POPPLER_PAGE_TRANSITION_FADE     // The new page gradually becomes visible through the old one
)

// PageTransitionAlignment types for #POPPLER_PAGE_TRANSITION_SPLIT and #POPPLER_PAGE_TRANSITION_BLINDS transition types
type PageTransitionAlignment int

const (
	PageTransitionHorizontal PageTransitionAlignment = C.POPPLER_PAGE_TRANSITION_HORIZONTAL // Horizontal dimension
	PageTransitionVertical                           = C.POPPLER_PAGE_TRANSITION_VERTICAL   // Vertical dimension
)

// PageTransitionDirection types for #POPPLER_PAGE_TRANSITION_SPLIT, POPPLER_PAGE_TRANSITION_BOX and POPPLER_PAGE_TRANSITION_FLY transition types
type PageTransitionDirection int

const (
	PageTransitionInward  PageTransitionDirection = C.POPPLER_PAGE_TRANSITION_INWARD  // Inward from the edges of the page
	PageTransitionOutward                         = C.POPPLER_PAGE_TRANSITION_OUTWARD // Outward from the center of the page
)

// SelectionStyle styles
type SelectionStyle int

const (
	SelectionGlyph SelectionStyle = C.POPPLER_SELECTION_GLYPH // Glyph is the minimum unit for selection
	SelectionWord                 = C.POPPLER_SELECTION_WORD  // Word is the minimum unit for selection
	SelectionLine                 = C.POPPLER_SELECTION_LINE  // Line is the minimum unit for selection
)

// PrintFlags for printing (Since 0.16)
type PrintFlags int

const (
	PrintDocument        PrintFlags = C.POPPLER_PRINT_DOCUMENT          // Print main document contents
	PrintMarkupAnnots               = C.POPPLER_PRINT_MARKUP_ANNOTS     // Print document and markup annotations
	PrintStampAnnotsOnly            = C.POPPLER_PRINT_STAMP_ANNOTS_ONLY // Print document and only stamp annotations
	PrintAll                        = C.POPPLER_PRINT_ALL               // Print main document contents and all markup annotations
)

// FindFlags while searching text in a page
type FindFlags int

const (
	FindDefault        FindFlags = C.POPPLER_FIND_DEFAULT
	FindCaseSensitive            = C.POPPLER_FIND_CASE_SENSITIVE
	FindBackwards                = C.POPPLER_FIND_BACKWARDS
	FindWholeWordsOnly           = C.POPPLER_FIND_WHOLE_WORDS_ONLY
)

// Backend codes returned by poppler_get_backend()
type Backend int

const (
	BackendUnknown Backend = C.POPPLER_BACKEND_UNKNOWN
	BackendSplash          = C.POPPLER_BACKEND_SPLASH
	BackendCairo           = C.POPPLER_BACKEND_CAIRO
)

func (b Backend) String() string {
	switch b {
	case BackendSplash:
		return "Splash"
	case BackendCairo:
		return "Cairo"
	}
	return "Unknown"
}

func GetVersion() string {
	// Don't free this value
	return C.GoString(C.poppler_get_version())
}

func GetBackend() Backend {
	return Backend(C.poppler_get_backend())
}
