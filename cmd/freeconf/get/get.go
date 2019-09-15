package get

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Run() {
	dir := flag.String("dir", ".", "Output directory for YANG files.")

	flag.Parse()

	cmd := exec.Command("go", "list", "-m", "-f={{.Dir}}", "all")
	cmd.Stderr = os.Stderr
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("could not list modules. %s", err)
	}
	lines := bufio.NewReader(&out)
	for {
		line, err := lines.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("could not read output. %s", err)
			}
		}
		yangDir := filepath.Join(strings.Trim(line, "\n"), "yang")
		if err := syncDir(yangDir, *dir); err != nil {
			log.Fatalf("Could not sync yang directory %s", err)
		}
	}
}

func syncDir(src, dest string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		// doesn't exist, completely normal
		return nil
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yang") {
			if err := copyFile(filepath.Join(src, f.Name()), dest); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, destDir string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	destFname := filepath.Join(destDir, filepath.Base(src))
	log.Print(destFname)
	destFile, err := os.Create(destFname)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(srcFile, destFile)
	return err
}
