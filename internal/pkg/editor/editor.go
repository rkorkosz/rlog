package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
)

const DefaultEditor = "vim"

func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	exe, err := exec.LookPath(editor)
	if err != nil {
		return err
	}
	cmd := exec.Command(exe, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CaptureFromEditor() ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return nil, err
	}
	filename := file.Name()
	defer os.Remove(filename)
	err = file.Close()
	if err != nil {
		return nil, err
	}
	err = OpenFileInEditor(filename)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(filename)
}
