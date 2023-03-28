package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRunCommand(t *testing.T) {
	// Set up a temporary directory for the test
	tmpDir, err := ioutil.TempDir("", "test-run-command")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a file to use as the command argument
	cmdFile := filepath.Join(tmpDir, "cmd.txt")
	err = ioutil.WriteFile(cmdFile, []byte("hello world"), 0644)
	if err != nil {
		t.Fatalf("failed to create command file: %v", err)
	}

	// Call the function with a valid command and path
	err = RunCommand("cat", cmdFile)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Call the function with an invalid command
	err = RunCommand("invalid-command", cmdFile)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	// path, err := exec.Command("which", "git").Output()
	// if err != nil {
	// 	// handle error
	// 	t.Errorf("git not found: %v", err)
	// }

	// fmt.Println(exec.Command(strings.TrimSpace(string(path)), "clone", "https://github.com/shashank-priyadarshi/upgraded-disco.git").CombinedOutput())

	// Call git clone command
	// err = RunCommand("git clone https://github.com/shashank-priyadarshi/upgraded-disco.git", cmdFile)
	// if err != nil {
	// 	t.Errorf("unexpected error: %v", err)
	// }
}
