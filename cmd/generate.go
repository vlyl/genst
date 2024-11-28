package cmd

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates templates/.gitignore.tmpl templates/.github
var templates embed.FS

var rootCmd = &cobra.Command{
	Use:   "genst",
	Short: "Generate a new Golang HTTP server project",
	Long: `Genst is a CLI tool that generates a new Golang HTTP server project with 
common components and best practices pre-configured.`,
}

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	RunE:  runNew,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func runNew(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create project structure
	dirs := []string{
		"cmd/server",
		"internal/api",
		"internal/config",
		"internal/middleware",
		"internal/model",
		"internal/service",
		"pkg/logger",
		"pkg/database",
		"config",
		"scripts",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectName, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Generate files from templates
	templateFiles := map[string]string{
		"templates/scripts/start.sh":               "scripts/start.sh",
		"templates/Makefile":                       "Makefile",
		"templates/config/config.yaml":             "config/config.yaml",
		"templates/server/main.go.tmpl":            "cmd/server/main.go",
		"templates/internal/api/router.go.tmpl":    "internal/api/router.go",
		"templates/internal/config/config.go.tmpl": "internal/config/config.go",
		"templates/pkg/logger/logger.go.tmpl":      "pkg/logger/logger.go",
		"templates/pkg/database/db.go.tmpl":        "pkg/database/db.go",
		"templates/README.md":                      "README.md",
		"templates/go.mod.tmpl":                    "go.mod",
		"templates/.gitignore.tmpl":                ".gitignore",
		"templates/.github/workflows/ci.yml.tmpl":  ".github/workflows/ci.yml",
	}

	data := struct {
		ProjectName string
		ModuleName  string
	}{
		ProjectName: projectName,
		ModuleName:  projectName,
	}

	for tmpl, dest := range templateFiles {
		if err := generateFile(tmpl, filepath.Join(projectName, dest), data); err != nil {
			return fmt.Errorf("failed to generate %s: %w", dest, err)
		}
	}

	// Initialize git repository
	if err := initGitRepo(projectName); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	fmt.Printf("Successfully created project %s\n", projectName)
	return nil
}

func generateFile(tmplPath, destPath string, data interface{}) error {
	content, err := templates.ReadFile(tmplPath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(content))
	if err != nil {
		return err
	}

	// Create destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if strings.HasSuffix(destPath, ".sh") {
		if err := os.Chmod(destPath, 0755); err != nil {
			return err
		}
	}

	return tmpl.Execute(f, data)
}

func initGitRepo(projectPath string) error {
	// Change to project directory
	if err := os.Chdir(projectPath); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}
	defer os.Chdir("..")

	// Initialize git repository
	if err := exec.Command("git", "init").Run(); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	return nil
}
