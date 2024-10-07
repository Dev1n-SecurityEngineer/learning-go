package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Parse command line argument
	logDir := flag.String("log-directory", "/var/log", "Directory containing logs to archive")
	perm := flag.Int("directory-permission", 0755, "Directory permission")
	flag.Parse()

	// Create archive directory
	archiveDir := filepath.Join(*logDir, "archive")
	if err := os.MkdirAll(archiveDir, os.FileMode(*perm)); err != nil { // Create directory with read/write/execute permissions for owner and read/execute permissions for group and others
		log.Fatalf("Error creating archive directory %v: %v", archiveDir, err) // Log error message and exit program
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
		file, err := os.Open(path) // Open file
		if err != nil {
			return err
		}
		defer file.Close()

		// Create tar header
		header, err := tar.FileInfoHeader(info, info.Name()) // Create header from file info
		if err != nil {
			return err
		}
		header.Name = filepath.Join(filepath.Base(*logDir), info.Name()) // Set file name in tar archive

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
