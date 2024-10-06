package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Parse command line argument
	logDir := flag.String("log-directory", "/var/log", "Directory containing logs to archive")
	flag.Parse()

	// Create archive directory
	archiveDir := filepath.Join(*logDir, "archive")
	err := os.MkdirAll(archiveDir, 0755) // Create directory with read/write/execute permissions for owner and read/execute permissions for group and others
	if err != nil {
		fmt.Printf("Error creating archive directory: %v\n", err) // Print error message
		return
	}

	// Create tar.gz file
	tarGzPath := filepath.Join(archiveDir, fmt.Sprintf("logs-%s.tar.gz", time.Now().Format("20060102")))
	tarGzFile, err := os.Create(tarGzPath)
	if err != nil {
		fmt.Printf("Error creating tar.gz file: %v\n", err)
		return
	}
	defer tarGzFile.Close() // Close file when function returns

	gw := gzip.NewWriter(tarGzFile) // Create gzip writer
	defer gw.Close()                // Close writer when function returns

	tw := tar.NewWriter(gw) // Create tar writer
	defer tw.Close()        // Close writer when function returns

	// Walk through log directory and add files to tar.gz
	err = filepath.Walk(*logDir, func(path string, info os.FileInfo, err error) error { // Walk through directory and call function for each file
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create tar header
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = filepath.Join(filepath.Base(*logDir), info.Name())

		// Write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// Write file data
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error archiving logs: %v\n", err)
		return
	}

	fmt.Printf("Logs archived successfully to %s\n", tarGzPath)
}
