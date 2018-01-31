package poppler

/*
#cgo pkg-config: --cflags-only-I poppler-glib
#include <stdio.h>
#include <stdlib.h>
#include <poppler.h>
*/
import "C"
import (
	"fmt"
	"path/filepath"
	"time"
)

// PageLayout types
type PageLayout int

const (
	PageLayoutUnset          PageLayout = C.POPPLER_PAGE_LAYOUT_UNSET            // No specific layout set
	PageLayoutSinglePage                = C.POPPLER_PAGE_LAYOUT_SINGLE_PAGE      // One page at a time
	PageLayoutOneColumn                 = C.POPPLER_PAGE_LAYOUT_ONE_COLUMN       // Pages in one column
	PageLayoutTwoColumnLeft             = C.POPPLER_PAGE_LAYOUT_TWO_COLUMN_LEFT  // Pages in two columns with odd numbered pages on the left
	PageLayoutTwoColumnRight            = C.POPPLER_PAGE_LAYOUT_TWO_COLUMN_RIGHT // Pages in two columns with odd numbered pages on the right
	PageLayoutTwoPageLeft               = C.POPPLER_PAGE_LAYOUT_TWO_PAGE_LEFT    // Two pages at a time with odd numbered pages on the left
	PageLayoutTwoPageRight              = C.POPPLER_PAGE_LAYOUT_TWO_PAGE_RIGHT   // Two pages at a time with odd numbered pages on the right
)

// PageMode
type PageMode int

const (
	PageModeUnset          PageMode = C.POPPLER_PAGE_MODE_UNSET           // No specific mode set
	PageModeNone                    = C.POPPLER_PAGE_MODE_NONE            // Neither document ouline nor thumbnails visible
	PageModeUseOutlines             = C.POPPLER_PAGE_MODE_USE_OUTLINES    // Document outline visible
	PageModeUseThumbs               = C.POPPLER_PAGE_MODE_USE_THUMBS      // Thumbnails visible
	PageModeFullScreen              = C.POPPLER_PAGE_MODE_FULL_SCREEN     // Full-screen mode
	PageModeUseOC                   = C.POPPLER_PAGE_MODE_USE_OC          // Layers panel visible
	PageModeUseAttachments          = C.POPPLER_PAGE_MODE_USE_ATTACHMENTS // Attachments panel visible
)

// FontType are types of fonts
type FontType int

const (
	FontTypeUnknown     FontType = C.POPPLER_FONT_TYPE_UNKNOWN      // Unknown font type
	FontTypeType1                = C.POPPLER_FONT_TYPE_TYPE1        // Type 1 font type
	FontTypeType1C               = C.POPPLER_FONT_TYPE_TYPE1C       // Type 1 font type embedded in Compact Font Format (CFF) font program
	FontTypeType1COT             = C.POPPLER_FONT_TYPE_TYPE1COT     // Type 1 font type embedded in OpenType font program
	FontTypeType3                = C.POPPLER_FONT_TYPE_TYPE3        // A font type that is defined with PDF graphics operators
	FontTypeTrueType             = C.POPPLER_FONT_TYPE_TRUETYPE     // TrueType font type
	FontTypeTrueTypeOT           = C.POPPLER_FONT_TYPE_TRUETYPEOT   // TrueType font type embedded in OpenType font program
	FontTypeCIDType0             = C.POPPLER_FONT_TYPE_CID_TYPE0    // CIDFont type based on Type 1 font technology
	FontTypeCIDType0C            = C.POPPLER_FONT_TYPE_CID_TYPE0C   // CIDFont type based on Type 1 font technology embedded in CFF font program
	FontTypeCIDType0COT          = C.POPPLER_FONT_TYPE_CID_TYPE0COT // CIDFont type based on Type 1 font technology embedded in OpenType font program
	FontTypeCIDType2             = C.POPPLER_FONT_TYPE_CID_TYPE2    // CIDFont type based on TrueType font technology
	FontTypeCIDType2OT           = C.POPPLER_FONT_TYPE_CID_TYPE2OT  // CIDFont type based on TrueType font technology embedded in OpenType font program
)

// ViewerPreferences
type ViewerPreferences int

const (
	ViewerPreferencesUnset           ViewerPreferences = C.POPPLER_VIEWER_PREFERENCES_UNSET             // No preferences set
	ViewerPreferencesHideToolbar                       = C.POPPLER_VIEWER_PREFERENCES_HIDE_TOOLBAR      // Hider toolbars when document is active
	ViewerPreferencesHideMenubar                       = C.POPPLER_VIEWER_PREFERENCES_HIDE_MENUBAR      // Hide menu bar when document is active
	ViewerPreferencesHideWindowUI                      = C.POPPLER_VIEWER_PREFERENCES_HIDE_WINDOWUI     // Hide UI elements in document's window
	ViewerPreferencesFitWindow                         = C.POPPLER_VIEWER_PREFERENCES_FIT_WINDOW        // Resize document's window to fit the size of the first displayed page
	ViewerPreferencesCenterWindow                      = C.POPPLER_VIEWER_PREFERENCES_CENTER_WINDOW     // Position the document's window in the center of the screen
	ViewerPreferencesDisplayDOCTitle                   = C.POPPLER_VIEWER_PREFERENCES_DISPLAY_DOC_TITLE // Display document title in window's title bar
	ViewerPreferencesDirectionRTL                      = C.POPPLER_VIEWER_PREFERENCES_DIRECTION_RTL     // The predominant reading order for text is right to left
)

// Permissions
type Permissions int

const (
	PermissionsOKToPrint               Permissions = C.POPPLER_PERMISSIONS_OK_TO_PRINT                 // Document can be printer
	PermissionsOKToModify                          = C.POPPLER_PERMISSIONS_OK_TO_MODIFY                // Document contents can be modified
	PermissionsOKToCopy                            = C.POPPLER_PERMISSIONS_OK_TO_COPY                  // Document can be copied
	PermissionsOKToAddNotes                        = C.POPPLER_PERMISSIONS_OK_TO_ADD_NOTES             // Annotations can added to the document
	PermissionsOKToFillForm                        = C.POPPLER_PERMISSIONS_OK_TO_FILL_FORM             // Interactive form fields can be filled in
	PermissionsOKToExtractContents                 = C.POPPLER_PERMISSIONS_OK_TO_EXTRACT_CONTENTS      // Extract text and graphics (in support of accessibility to users with disabilities or for other purposes). Since 0.18
	PermissionsOKToAssemble                        = C.POPPLER_PERMISSIONS_OK_TO_ASSEMBLE              // Assemble the document (insert, rotate, or delete pages and create bookmarks or thumbnail images). Since 0.18
	PermissionsOKToPrintHighResolution             = C.POPPLER_PERMISSIONS_OK_TO_PRINT_HIGH_RESOLUTION // Document can be printer at high resolution. Since 0.18
	PermissionsFull                                = C.POPPLER_PERMISSIONS_FULL                        // Document permits all operations
)

type popDocument *C.PopplerDocument

type Document struct {
	doc popDocument
}

// NewFromFile Loads document from file path
func NewFromFile(filename, password string) (*Document, error) {
	fp, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	cFilepath := C.CString(fmt.Sprintf("file://%s", fp))
	// defer C.free(unsafe.Pointer(cFilepath))
	cPassword := C.CString(password)
	// defer C.free(unsafe.Pointer(cPassword))

	var gErr *C.GError
	// defer C.free(unsafe.Pointer(gErr))

	var popDoc popDocument
	if password == "" {
		popDoc = C.poppler_document_new_from_file(cFilepath, nil, &gErr)
	} else {
		// Password protected file
		popDoc = C.poppler_document_new_from_file(cFilepath, cPassword, &gErr)
	}
	if gErr != nil {
		return nil, fmt.Errorf(C.GoString((*C.char)(gErr.message)))
	}

	return &Document{doc: popDoc}, nil
}

// // Destroy cleans up allocated C memory for loaded document
// func (d *Document) Destroy() {
// 	// if d.doc != nil {
// 	// 	C.free(unsafe.Pointer(d.doc))
// 	// }
// }

// GetTitle of the document
func (d *Document) GetTitle() string {
	cTitle := C.poppler_document_get_title(d.doc)
	// defer C.free(unsafe.Pointer(cTitle))
	return C.GoString((*C.char)(cTitle))
}

// GetPDFVersion of the document
func (d *Document) GetPDFVersion() string {
	return C.GoString((*C.char)(C.poppler_document_get_pdf_version_string(d.doc)))
}

// GetTotalPages of the document
func (d *Document) GetTotalPages() int {
	return int(C.poppler_document_get_n_pages(d.doc))
}

// GetAuthor of the document
func (d *Document) GetAuthor() string {
	return C.GoString((*C.char)(C.poppler_document_get_author(d.doc)))
}

// GetSubject of the document
func (d *Document) GetSubject() string {
	return C.GoString((*C.char)(C.poppler_document_get_subject(d.doc)))
}

// GetKeywords for the document
func (d *Document) GetKeywords() string {
	return C.GoString((*C.char)(C.poppler_document_get_keywords(d.doc)))
}

// GetCreator of the document
func (d *Document) GetCreator() string {
	return C.GoString((*C.char)(C.poppler_document_get_creator(d.doc)))
}

// GetProducer of the document
func (d *Document) GetProducer() string {
	return C.GoString((*C.char)(C.poppler_document_get_producer(d.doc)))
}

// GetMetadata for the document
func (d *Document) GetMetadata() string {
	return C.GoString((*C.char)(C.poppler_document_get_metadata(d.doc)))
}

// GetCreatedDate of the document
func (d *Document) GetCreatedDate() time.Time {
	t := int64(C.poppler_document_get_creation_date(d.doc))
	// i, _ := strconv.ParseInt(t, 10, 64)
	return time.Unix(t, 0)
}

// GetModificationDate of the document
func (d *Document) GetModificationDate() time.Time {
	t := int64(C.poppler_document_get_modification_date(d.doc))
	// i, _ := strconv.ParseInt(t, 10, 64)
	return time.Unix(t, 0)
}

// IsLinearized or not
func (d *Document) IsLinearized() bool {
	return int(C.poppler_document_is_linearized(d.doc)) > 0
}
