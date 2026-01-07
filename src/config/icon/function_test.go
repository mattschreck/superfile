package icon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInitIcon_WithNerdFont tests InitIcon when nerdfont is enabled
func TestInitIcon_WithNerdFont(t *testing.T) {
	// Save original values
	originalSpace := Space
	originalSuperfileIcon := SuperfileIcon
	originalHome := Home
	originalCursor := Cursor
	defer func() {
		Space = originalSpace
		SuperfileIcon = originalSuperfileIcon
		Home = originalHome
		Cursor = originalCursor
	}()

	// Test with nerdfont enabled
	InitIcon(true, "#FF0000")

	// When nerdfont is enabled, icons should remain as Unicode characters
	assert.NotEqual(t, "", Space, "Space should not be empty with nerdfont")
	assert.NotEqual(t, "", SuperfileIcon, "SuperfileIcon should not be empty with nerdfont")
	assert.NotEqual(t, "", Home, "Home icon should not be empty with nerdfont")
	assert.NotEqual(t, ">", Cursor, "Cursor should be Unicode icon with nerdfont")

	// Check folder icon was set
	folderStyle, exists := Folders["folder"]
	assert.True(t, exists, "folder entry should exist in Folders map")
	assert.Equal(t, "#FF0000", folderStyle.Color, "folder color should match input")
	assert.NotEmpty(t, folderStyle.Icon, "folder icon should not be empty")
}

// TestInitIcon_WithoutNerdFont tests InitIcon when nerdfont is disabled
func TestInitIcon_WithoutNerdFont(t *testing.T) {
	// Save original values
	originalSpace := Space
	originalSuperfileIcon := SuperfileIcon
	originalHome := Home
	originalCursor := Cursor
	originalDirectory := Directory
	defer func() {
		Space = originalSpace
		SuperfileIcon = originalSuperfileIcon
		Home = originalHome
		Cursor = originalCursor
		Directory = originalDirectory
	}()

	// Test with nerdfont disabled
	InitIcon(false, "#00FF00")

	// When nerdfont is disabled, icons should be ASCII characters
	assert.Equal(t, "", Space, "Space should be empty without nerdfont")
	assert.Equal(t, "", SuperfileIcon, "SuperfileIcon should be empty without nerdfont")
	assert.Equal(t, "", Home, "Home icon should be empty without nerdfont")
	assert.Equal(t, ">", Cursor, "Cursor should be '>' without nerdfont")
	assert.Equal(t, "^", SortAsc, "SortAsc should be '^' without nerdfont")
	assert.Equal(t, "v", SortDesc, "SortDesc should be 'v' without nerdfont")
	assert.Equal(t, "B", Browser, "Browser should be 'B' without nerdfont")
	assert.Equal(t, "S", Select, "Select should be 'S' without nerdfont")
	assert.Equal(t, "", Directory, "Directory should be empty without nerdfont")

	// Check folder icon was set
	folderStyle, exists := Folders["folder"]
	assert.True(t, exists, "folder entry should exist in Folders map")
	assert.Equal(t, "#00FF00", folderStyle.Color, "folder color should match input")
	assert.NotEmpty(t, folderStyle.Icon, "folder icon should not be empty")
}

// TestInitIcon_DefaultDirectoryColor tests InitIcon with empty directoryIconColor
func TestInitIcon_DefaultDirectoryColor(t *testing.T) {
	// Test with empty directoryIconColor (should default to "NONE")
	InitIcon(true, "")

	folderStyle, exists := Folders["folder"]
	assert.True(t, exists, "folder entry should exist in Folders map")
	assert.Equal(t, "NONE", folderStyle.Color, "folder color should default to 'NONE'")
}

// TestInitIcon_CustomDirectoryColor tests InitIcon with custom directoryIconColor
func TestInitIcon_CustomDirectoryColor(t *testing.T) {
	tests := []struct {
		name  string
		color string
	}{
		{"Hex color", "#FF5733"},
		{"Named color", "blue"},
		{"RGB color", "rgb(255, 87, 51)"},
		{"Special value", "NONE"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitIcon(true, tt.color)

			folderStyle, exists := Folders["folder"]
			assert.True(t, exists, "folder entry should exist in Folders map")
			assert.Equal(t, tt.color, folderStyle.Color, "folder color should match input")
		})
	}
}

// TestGetCopyOrCutIcon_Cut tests GetCopyOrCutIcon when cut is true
func TestGetCopyOrCutIcon_Cut(t *testing.T) {
	result := GetCopyOrCutIcon(true)
	assert.Equal(t, Cut, result, "Should return Cut icon when cut is true")
}

// TestGetCopyOrCutIcon_Copy tests GetCopyOrCutIcon when cut is false
func TestGetCopyOrCutIcon_Copy(t *testing.T) {
	result := GetCopyOrCutIcon(false)
	assert.Equal(t, Copy, result, "Should return Copy icon when cut is false")
}

// TestGetCopyOrCutIcon_Consistency tests that function returns correct icons consistently
func TestGetCopyOrCutIcon_Consistency(t *testing.T) {
	// Test multiple calls with same input return same result
	result1 := GetCopyOrCutIcon(true)
	result2 := GetCopyOrCutIcon(true)
	assert.Equal(t, result1, result2, "Multiple calls with cut=true should return same icon")

	result3 := GetCopyOrCutIcon(false)
	result4 := GetCopyOrCutIcon(false)
	assert.Equal(t, result3, result4, "Multiple calls with cut=false should return same icon")

	// Cut and Copy icons should be different
	assert.NotEqual(t, result1, result3, "Cut and Copy icons should be different")
}

// TestInitIcon_FileOperationIcons tests that file operation icons behavior with InitIcon
func TestInitIcon_FileOperationIcons(t *testing.T) {
	t.Run("File operations without nerdfont are cleared", func(t *testing.T) {
		// Save originals
		origCompress := CompressFile
		origExtract := ExtractFile
		origCopy := Copy
		origCut := Cut
		origDelete := Delete
		defer func() {
			CompressFile = origCompress
			ExtractFile = origExtract
			Copy = origCopy
			Cut = origCut
			Delete = origDelete
		}()

		// Set some values first
		CompressFile = "test"
		ExtractFile = "test"
		Copy = "test"
		Cut = "test"
		Delete = "test"

		// InitIcon with nerdfont=false should clear them
		InitIcon(false, "")

		// Without nerdfont, file operation icons should be empty
		assert.Equal(t, "", CompressFile, "CompressFile should be empty")
		assert.Equal(t, "", ExtractFile, "ExtractFile should be empty")
		assert.Equal(t, "", Copy, "Copy should be empty")
		assert.Equal(t, "", Cut, "Cut should be empty")
		assert.Equal(t, "", Delete, "Delete should be empty")
	})
}

// TestInitIcon_UIElementIcons tests that UI element icons behavior with InitIcon
func TestInitIcon_UIElementIcons(t *testing.T) {
	t.Run("UI elements without nerdfont are cleared", func(t *testing.T) {
		// Save originals
		origError := Error
		origWarn := Warn
		origDone := Done
		origInOperation := InOperation
		origSearch := Search
		defer func() {
			Error = origError
			Warn = origWarn
			Done = origDone
			InOperation = origInOperation
			Search = origSearch
		}()

		// Set some values first
		Error = "test"
		Warn = "test"
		Done = "test"
		InOperation = "test"
		Search = "test"

		// InitIcon with nerdfont=false should clear them
		InitIcon(false, "")

		// Without nerdfont, these should be empty
		assert.Equal(t, "", Error, "Error icon should be empty")
		assert.Equal(t, "", Warn, "Warn icon should be empty")
		assert.Equal(t, "", Done, "Done icon should be empty")
		assert.Equal(t, "", InOperation, "InOperation icon should be empty")
		assert.Equal(t, "", Search, "Search icon should be empty")
	})
}
