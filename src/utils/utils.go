package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

func LoadJSONConfig(path string, v interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// Read and print the content of the config file before decoding
	_, readErr := io.ReadAll(configFile)
	if readErr != nil {
		return readErr
	}
	//LogInfo("Content of Config File:", path, "\n", string(content))

	// Reset the read pointer to the start of the file
	configFile.Seek(0, 0)

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func MakeTextCenter(text string, width int) string {
	textLen := len(text)
	if textLen >= width {
		return text
	}
	spaceNum := width - textLen
	leftSpaceNum := spaceNum / 2
	rightSpaceNum := spaceNum - leftSpaceNum
	return strings.Repeat(" ", leftSpaceNum) + text + strings.Repeat(" ", rightSpaceNum)
}

func HandleCtrlC(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range c {
			f()
		}
	}()
}

func GetAssemblyFolder() string {

	absPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		LogError("Error getting absolute path of assembly file: ", err)
		// return Home folder
		return os.Getenv("HOME")
	}
	return filepath.Dir(absPath)
}

func GetWorkingFolder() string {
	wd, err := os.Getwd()
	if err != nil {
		LogError("Error getting working directory: ", err)
		// return Home folder
		return os.Getenv("HOME")
	}
	return wd
}

func ReplaceFirstInFile(regexPattern string, newString string, filePath string) error {
	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Create a regular expression pattern from the provided regexPattern
	regex := regexp.MustCompile(regexPattern)

	// Find the first match in the content
	indexes := regex.FindStringIndex(string(content))
	if indexes == nil {
		return fmt.Errorf("no match found for regex pattern: %s,  in file: %s", regexPattern, filePath)
	}

	// Replace the matched substring with the newString
	replacedContent := regex.ReplaceAllString(string(content), newString)

	// Write the modified content back to the file
	err = os.WriteFile(filePath, []byte(replacedContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func IsFolderExist(folderPath string) bool {
	_, err := os.Stat(folderPath)
	return err == nil
}

func AskForYes(title string) bool {
	var input string
	fmt.Print(title + "(Y/n):")
	fmt.Scanln(&input)
	input = strings.ToUpper(input)
	return (input != "N" && input != "NO")
}
