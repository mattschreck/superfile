package variable

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

// TestSetLastDir tests the SetLastDir function
func TestSetLastDir(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Set valid path",
			path:     "/home/user/documents",
			expected: "/home/user/documents",
		},
		{
			name:     "Set empty path",
			path:     "",
			expected: "",
		},
		{
			name:     "Set path with spaces",
			path:     "/home/user/my documents",
			expected: "/home/user/my documents",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLastDir(tt.path)
			assert.Equal(t, tt.expected, LastDir)
		})
	}
}

// TestSetChooserFile tests the SetChooserFile function
func TestSetChooserFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Set valid file path",
			path:     "/tmp/chooser.txt",
			expected: "/tmp/chooser.txt",
		},
		{
			name:     "Set empty path",
			path:     "",
			expected: "",
		},
		{
			name:     "Set path with special characters",
			path:     "/tmp/chooser-file_123.txt",
			expected: "/tmp/chooser-file_123.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetChooserFile(tt.path)
			assert.Equal(t, tt.expected, ChooserFile)
		})
	}
}

// TestUpdateVarFromCliArgs_ConfigFile tests UpdateVarFromCliArgs with config-file argument
func TestUpdateVarFromCliArgs_ConfigFile(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.toml")
	err := os.WriteFile(configFile, []byte("# test config"), 0644)
	require.NoError(t, err)

	// Create CLI command with config-file argument
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config-file",
				Value: configFile,
			},
			&cli.StringFlag{
				Name: "hotkey-file",
			},
			&cli.StringFlag{
				Name: "chooser-file",
			},
			&cli.BoolFlag{
				Name: "fix-hotkeys",
			},
			&cli.BoolFlag{
				Name: "fix-config-file",
			},
			&cli.BoolFlag{
				Name: "print-last-dir",
			},
		},
	}

	// Reset ConfigFile to default before test
	originalConfigFile := ConfigFile
	defer func() { ConfigFile = originalConfigFile }()

	UpdateVarFromCliArgs(cmd)

	assert.Equal(t, configFile, ConfigFile)
}

// TestUpdateVarFromCliArgs_HotkeyFile tests UpdateVarFromCliArgs with hotkey-file argument
func TestUpdateVarFromCliArgs_HotkeyFile(t *testing.T) {
	// Create a temporary hotkey file
	tmpDir := t.TempDir()
	hotkeyFile := filepath.Join(tmpDir, "hotkeys.toml")
	err := os.WriteFile(hotkeyFile, []byte("# test hotkeys"), 0644)
	require.NoError(t, err)

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config-file",
			},
			&cli.StringFlag{
				Name:  "hotkey-file",
				Value: hotkeyFile,
			},
			&cli.StringFlag{
				Name: "chooser-file",
			},
			&cli.BoolFlag{
				Name: "fix-hotkeys",
			},
			&cli.BoolFlag{
				Name: "fix-config-file",
			},
			&cli.BoolFlag{
				Name: "print-last-dir",
			},
		},
	}

	originalHotkeysFile := HotkeysFile
	defer func() { HotkeysFile = originalHotkeysFile }()

	UpdateVarFromCliArgs(cmd)

	assert.Equal(t, hotkeyFile, HotkeysFile)
}

// TestUpdateVarFromCliArgs_ChooserFile tests UpdateVarFromCliArgs with chooser-file argument
func TestUpdateVarFromCliArgs_ChooserFile(t *testing.T) {
	chooserPath := "/tmp/chooser-test.txt"

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config-file",
			},
			&cli.StringFlag{
				Name: "hotkey-file",
			},
			&cli.StringFlag{
				Name:  "chooser-file",
				Value: chooserPath,
			},
			&cli.BoolFlag{
				Name: "fix-hotkeys",
			},
			&cli.BoolFlag{
				Name: "fix-config-file",
			},
			&cli.BoolFlag{
				Name: "print-last-dir",
			},
		},
	}

	originalChooserFile := ChooserFile
	defer func() { ChooserFile = originalChooserFile }()

	UpdateVarFromCliArgs(cmd)

	assert.Equal(t, chooserPath, ChooserFile)
}

// TestUpdateVarFromCliArgs_BoolFlags tests UpdateVarFromCliArgs with boolean flags
func TestUpdateVarFromCliArgs_BoolFlags(t *testing.T) {
	tests := []struct {
		name                string
		fixHotkeys          bool
		fixConfigFile       bool
		printLastDir        bool
		expectedFixHotkeys  bool
		expectedFixConfig   bool
		expectedPrintLastDir bool
	}{
		{
			name:                "All flags true",
			fixHotkeys:          true,
			fixConfigFile:       true,
			printLastDir:        true,
			expectedFixHotkeys:  true,
			expectedFixConfig:   true,
			expectedPrintLastDir: true,
		},
		{
			name:                "All flags false",
			fixHotkeys:          false,
			fixConfigFile:       false,
			printLastDir:        false,
			expectedFixHotkeys:  false,
			expectedFixConfig:   false,
			expectedPrintLastDir: false,
		},
		{
			name:                "Mixed flags",
			fixHotkeys:          true,
			fixConfigFile:       false,
			printLastDir:        true,
			expectedFixHotkeys:  true,
			expectedFixConfig:   false,
			expectedPrintLastDir: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cli.Command{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "config-file",
					},
					&cli.StringFlag{
						Name: "hotkey-file",
					},
					&cli.StringFlag{
						Name: "chooser-file",
					},
					&cli.BoolFlag{
						Name:  "fix-hotkeys",
						Value: tt.fixHotkeys,
					},
					&cli.BoolFlag{
						Name:  "fix-config-file",
						Value: tt.fixConfigFile,
					},
					&cli.BoolFlag{
						Name:  "print-last-dir",
						Value: tt.printLastDir,
					},
				},
			}

			// Reset flags before test
			originalFixHotkeys := FixHotkeys
			originalFixConfigFile := FixConfigFile
			originalPrintLastDir := PrintLastDir
			defer func() {
				FixHotkeys = originalFixHotkeys
				FixConfigFile = originalFixConfigFile
				PrintLastDir = originalPrintLastDir
			}()

			UpdateVarFromCliArgs(cmd)

			assert.Equal(t, tt.expectedFixHotkeys, FixHotkeys)
			assert.Equal(t, tt.expectedFixConfig, FixConfigFile)
			assert.Equal(t, tt.expectedPrintLastDir, PrintLastDir)
		})
	}
}

// TestConstants verifies that important constants are defined correctly
func TestConstants(t *testing.T) {
	t.Run("CurrentVersion is not empty", func(t *testing.T) {
		assert.NotEmpty(t, CurrentVersion)
		assert.Contains(t, CurrentVersion, "v")
	})

	t.Run("LatestVersionURL is valid", func(t *testing.T) {
		assert.NotEmpty(t, LatestVersionURL)
		assert.Contains(t, LatestVersionURL, "github.com")
		assert.Contains(t, LatestVersionURL, "releases")
	})

	t.Run("EmbedConfigFile is defined", func(t *testing.T) {
		assert.NotEmpty(t, EmbedConfigFile)
		assert.Contains(t, EmbedConfigFile, "config.toml")
	})

	t.Run("EmbedHotkeysFile is defined", func(t *testing.T) {
		assert.NotEmpty(t, EmbedHotkeysFile)
		assert.Contains(t, EmbedHotkeysFile, "hotkeys.toml")
	})
}

// TestDirectoryVariables verifies that directory paths are properly initialized
func TestDirectoryVariables(t *testing.T) {
	t.Run("SuperFileMainDir is not empty", func(t *testing.T) {
		assert.NotEmpty(t, SuperFileMainDir)
		assert.Contains(t, SuperFileMainDir, "superfile")
	})

	t.Run("SuperFileCacheDir is not empty", func(t *testing.T) {
		assert.NotEmpty(t, SuperFileCacheDir)
		assert.Contains(t, SuperFileCacheDir, "superfile")
	})

	t.Run("ConfigFile is properly joined", func(t *testing.T) {
		assert.NotEmpty(t, ConfigFile)
		assert.Equal(t, filepath.Join(SuperFileMainDir, "config.toml"), ConfigFile)
	})

	t.Run("HotkeysFile is properly joined", func(t *testing.T) {
		assert.NotEmpty(t, HotkeysFile)
		assert.Equal(t, filepath.Join(SuperFileMainDir, "hotkeys.toml"), HotkeysFile)
	})

	t.Run("LogFile is in StateDir", func(t *testing.T) {
		assert.NotEmpty(t, LogFile)
		assert.Contains(t, LogFile, SuperFileStateDir)
		assert.Contains(t, LogFile, "superfile.log")
	})
}
