package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	projectName := getUserInput("Enter project name: ")
	language := selectLanguage()
	projectPath := getUserInput("Enter project path: ")
	selectedTemplates := selectTemplates(language)

	if language == "C#" {
		err := createCSharpProject(projectName, projectPath)
		if err != nil {
			fmt.Printf("Error creating C# project: %s\n", err)
			os.Exit(1)
		}
	}

	if language == "Golang" {
		err := createGolangProject(projectName, projectPath)
		if err != nil {
			fmt.Printf("Error creating C# project: %s\n", err)
			os.Exit(1)
		}
	}

	err := createProject(projectName, language, projectPath, selectedTemplates)
	if err != nil {
		fmt.Printf("Error creating project: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Project created successfully!")
	getUserInput("Press Enter to Close")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}

func selectLanguage() string {
	languages := []string{"PHP", "C#", "Golang"}
	fmt.Println("Select programming language(Use Number):")
	for i, lang := range languages {
		fmt.Printf("%d. %s\n", i+1, lang)
	}

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(languages) {
		fmt.Println("Invalid choice.")
		os.Exit(1)
	}

	return languages[choice-1]
}

func selectTemplates(language string) []string {
	templateDir := fmt.Sprintf("templates/%s", language)
	fmt.Printf("Available templates for %s:\n", language)

	templates, err := os.ReadDir(templateDir)
	if err != nil {
		fmt.Printf("Error reading templates directory: %s\n", err)
		os.Exit(1)
	}

	selectedTemplates := make([]string, 0)
	for i, template := range templates {
		fmt.Printf("%d. %s\n", i+1, template.Name())
	}

	fmt.Print("Enter template numbers (comma-separated): ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	selectedTemplateIndices := strings.Split(userInput, ",")
	for _, indexStr := range selectedTemplateIndices {
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 1 || index > len(templates) {
			fmt.Println("Invalid template number.")
			os.Exit(1)
		}
		selectedTemplates = append(selectedTemplates, templates[index-1].Name())
	}

	return selectedTemplates
}

func createProject(projectName, language, projectPath string, selectedTemplates []string) error {
	projectDir := filepath.Join(projectPath, projectName)
	err := os.MkdirAll(projectDir, 0755)
	if err != nil {
		return err
	}

	templateDir := fmt.Sprintf("templates/%s", language)
	for _, template := range selectedTemplates {
		srcFile, err := os.Open(filepath.Join(templateDir, template))
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(filepath.Join(projectDir, template))
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func createCSharpProject(projectName, projectPath string) error {
	cmd := exec.Command("dotnet", "new", "console", "-n", projectName)
	cmd.Dir = projectPath
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func createGolangProject(projectName, projectPath string) error {
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectPath
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd_tidy := exec.Command("go", "mod", "tidy")
	cmd_tidy.Dir = projectPath + projectName
	err = cmd_tidy.Run()
	if err != nil {
		return err
	}

	return nil
}
