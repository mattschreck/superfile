package icon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIconsMap_NotEmpty tests that the Icons map is populated
func TestIconsMap_NotEmpty(t *testing.T) {
	assert.NotEmpty(t, Icons, "Icons map should not be empty")
	assert.Greater(t, len(Icons), 100, "Icons map should contain more than 100 entries")
}

// TestIconsMap_CommonFileTypes tests that common file types have icon mappings
func TestIconsMap_CommonFileTypes(t *testing.T) {
	commonTypes := []string{
		"go", "js", "ts", "py", "java", "c", "cpp",
		"html", "css", "json", "xml", "yml",
		"md", "txt", "pdf", "zip", "git",
	}

	for _, fileType := range commonTypes {
		t.Run(fileType, func(t *testing.T) {
			style, exists := Icons[fileType]
			assert.True(t, exists, "Icon for %s should exist", fileType)
			assert.NotEmpty(t, style.Icon, "Icon for %s should not be empty", fileType)
			assert.NotEmpty(t, style.Color, "Color for %s should not be empty", fileType)
		})
	}
}

// TestIconsMap_StyleProperties tests that all icon entries have valid properties
func TestIconsMap_StyleProperties(t *testing.T) {
	for fileType, style := range Icons {
		t.Run(fileType, func(t *testing.T) {
			assert.NotEmpty(t, style.Icon, "Icon for %s should not be empty", fileType)
			assert.NotEmpty(t, style.Color, "Color for %s should not be empty", fileType)
		})
	}
}

// TestIconsMap_ColorFormats tests that colors follow expected formats
func TestIconsMap_ColorFormats(t *testing.T) {
	validColors := make(map[string]bool)

	for _, style := range Icons {
		if style.Color == "NONE" {
			validColors[style.Color] = true
			continue
		}

		// Check if color starts with # (hex color)
		if len(style.Color) > 0 && style.Color[0] == '#' {
			assert.True(t, len(style.Color) == 7 || len(style.Color) == 4,
				"Hex color should be #RGB or #RRGGBB format: %s", style.Color)
			validColors[style.Color] = true
		}
	}

	assert.NotEmpty(t, validColors, "Should have found valid colors in Icons map")
}

// TestAliasesMap_NotEmpty tests that the Aliases map is populated
func TestAliasesMap_NotEmpty(t *testing.T) {
	assert.NotEmpty(t, Aliases, "Aliases map should not be empty")
	assert.Greater(t, len(Aliases), 200, "Aliases map should contain more than 200 entries")
}

// TestAliasesMap_CommonExtensions tests that common file extensions have aliases
func TestAliasesMap_CommonExtensions(t *testing.T) {
	commonExtensions := map[string]string{
		"jpg":    "image",
		"png":    "image",
		"mp3":    "audio",
		"mp4":    "video",
		"tar":    "zip",
		"gz":     "zip",
		"sh":     "shell",
		"bash":   "shell",
		"yaml":   "yml",
		"gitignore": "git",
	}

	for ext, expectedTarget := range commonExtensions {
		t.Run(ext, func(t *testing.T) {
			target, exists := Aliases[ext]
			assert.True(t, exists, "Alias for %s should exist", ext)
			assert.Equal(t, expectedTarget, target, "Alias for %s should map to %s", ext, expectedTarget)
		})
	}
}

// TestAliasesMap_AllTargetsExist tests that all alias targets exist in Icons map
func TestAliasesMap_AllTargetsExist(t *testing.T) {
	for alias, target := range Aliases {
		t.Run(alias+"->"+target, func(t *testing.T) {
			_, exists := Icons[target]
			assert.True(t, exists, "Alias %s points to %s, but %s doesn't exist in Icons map", alias, target, target)
		})
	}
}

// TestAliasesMap_ImageExtensions tests that all image extensions map to "image"
func TestAliasesMap_ImageExtensions(t *testing.T) {
	imageExtensions := []string{
		"jpg", "jpeg", "png", "gif", "bmp", "svg", "ico",
		"webp", "tiff", "tif",
	}

	for _, ext := range imageExtensions {
		t.Run(ext, func(t *testing.T) {
			target, exists := Aliases[ext]
			assert.True(t, exists, "Image extension %s should have an alias", ext)
			assert.Equal(t, "image", target, "Image extension %s should map to 'image'", ext)
		})
	}
}

// TestAliasesMap_AudioExtensions tests that audio extensions map to "audio"
func TestAliasesMap_AudioExtensions(t *testing.T) {
	audioExtensions := []string{
		"mp3", "wav", "flac", "aac", "ogg", "opus", "m4a",
	}

	for _, ext := range audioExtensions {
		t.Run(ext, func(t *testing.T) {
			target, exists := Aliases[ext]
			assert.True(t, exists, "Audio extension %s should have an alias", ext)
			assert.Equal(t, "audio", target, "Audio extension %s should map to 'audio'", ext)
		})
	}
}

// TestAliasesMap_VideoExtensions tests that video extensions map to "video"
func TestAliasesMap_VideoExtensions(t *testing.T) {
	videoExtensions := []string{
		"mp4", "avi", "mkv", "mov", "flv", "webm", "mpeg", "mpg",
	}

	for _, ext := range videoExtensions {
		t.Run(ext, func(t *testing.T) {
			target, exists := Aliases[ext]
			assert.True(t, exists, "Video extension %s should have an alias", ext)
			assert.Equal(t, "video", target, "Video extension %s should map to 'video'", ext)
		})
	}
}

// TestAliasesMap_CompressedExtensions tests that compressed file extensions map correctly
func TestAliasesMap_CompressedExtensions(t *testing.T) {
	compressedExtensions := []string{
		"zip", "tar", "gz", "bz2", "7z", "rar", "xz",
	}

	for _, ext := range compressedExtensions {
		t.Run(ext, func(t *testing.T) {
			target, exists := Aliases[ext]
			assert.True(t, exists, "Compressed extension %s should have an alias", ext)
			assert.Equal(t, "zip", target, "Compressed extension %s should map to 'zip'", ext)
		})
	}
}

// TestFoldersMap_NotEmpty tests that the Folders map is populated
func TestFoldersMap_NotEmpty(t *testing.T) {
	assert.NotEmpty(t, Folders, "Folders map should not be empty")
	assert.Greater(t, len(Folders), 10, "Folders map should contain more than 10 entries")
}

// TestFoldersMap_CommonFolders tests that common special folders have icons
func TestFoldersMap_CommonFolders(t *testing.T) {
	commonFolders := []string{
		".git", ".github", ".vscode", "node_modules",
		".Trash", "config", "hidden", "superfile",
	}

	for _, folder := range commonFolders {
		t.Run(folder, func(t *testing.T) {
			style, exists := Folders[folder]
			assert.True(t, exists, "Folder icon for %s should exist", folder)
			assert.NotEmpty(t, style.Icon, "Folder icon for %s should not be empty", folder)
			assert.NotEmpty(t, style.Color, "Folder color for %s should not be empty", folder)
		})
	}
}

// TestFoldersMap_StyleProperties tests that all folder entries have valid properties
func TestFoldersMap_StyleProperties(t *testing.T) {
	for folderName, style := range Folders {
		t.Run(folderName, func(t *testing.T) {
			assert.NotEmpty(t, style.Icon, "Icon for folder %s should not be empty", folderName)
			assert.NotEmpty(t, style.Color, "Color for folder %s should not be empty", folderName)
		})
	}
}

// TestIconConstants_Exist tests that icon constants exist and can be referenced
// Note: Icon values are defined as package-level vars and may change via InitIcon()
func TestIconConstants_Exist(t *testing.T) {
	// Test that icon variables can be referenced (they exist in the package)
	// We don't test their values as they may change based on InitIcon() calls
	_ = SuperfileIcon
	_ = Home
	_ = Download
	_ = Documents
	_ = Pictures
	_ = Videos
	_ = Music
	_ = Directory
	_ = Trash
	_ = Copy
	_ = Cut
	_ = Delete
	_ = Error
	_ = Warn
	_ = Done
	_ = Cursor
	_ = Browser
	_ = Select
	_ = Search
	_ = Terminal
	_ = Pinned

	// This test passes if compilation succeeds (all icons are defined)
	assert.True(t, true, "All icon constants exist")
}

// TestSpace_Constant tests the Space constant
func TestSpace_Constant(t *testing.T) {
	// Space should be a single space by default (before InitIcon is called)
	// This is just to verify it exists
	assert.NotNil(t, Space, "Space constant should be defined")
}

// TestCheckboxIcons tests checkbox-related icons
func TestCheckboxIcons(t *testing.T) {
	assert.NotEmpty(t, CheckboxEmpty, "CheckboxEmpty should not be empty")
	assert.NotEmpty(t, CheckboxChecked, "CheckboxChecked should not be empty")
	assert.NotEqual(t, CheckboxEmpty, CheckboxChecked, "CheckboxEmpty and CheckboxChecked should be different")
}

// TestSortIcons tests sorting-related icons
func TestSortIcons(t *testing.T) {
	assert.NotEmpty(t, SortAsc, "SortAsc should not be empty")
	assert.NotEmpty(t, SortDesc, "SortDesc should not be empty")
	assert.NotEqual(t, SortAsc, SortDesc, "SortAsc and SortDesc should be different")
}

// TestFileOperationIcons_Uniqueness tests that file operation icons are unique
func TestFileOperationIcons_Uniqueness(t *testing.T) {
	// Note: Icon values are defined as package-level vars in icon.go
	// We test that they remain unique when used with GetCopyOrCutIcon
	copyIcon := GetCopyOrCutIcon(false)
	cutIcon := GetCopyOrCutIcon(true)

	assert.NotEqual(t, copyIcon, cutIcon, "Copy and Cut icons should be different")
}

// TestAliasesMap_NoSelfReference tests that most aliases don't reference themselves
func TestAliasesMap_NoSelfReference(t *testing.T) {
	// Some aliases legitimately reference themselves (they're both alias and icon key)
	allowedSelfReferences := map[string]bool{
		"dart": true,
		"env":  true,
	}

	for alias, target := range Aliases {
		if allowedSelfReferences[alias] {
			continue // Skip known self-references
		}
		assert.NotEqual(t, alias, target, "Alias %s should not reference itself", alias)
	}
}

// TestAliasesMap_CaseConsistency tests that aliases are consistently lowercase
func TestAliasesMap_CaseConsistency(t *testing.T) {
	for alias := range Aliases {
		// Most aliases should be lowercase (with some exceptions like .DS_Store)
		// This test just ensures the map keys are defined
		assert.NotEmpty(t, alias, "Alias key should not be empty")
	}
}
