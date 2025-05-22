// Program to access GitHub releases page and print beautified JSON response

package main

import (
	"archive/tar"   // for untar files
	"archive/zip"   // for unzipping files
	"bytes"         // for handling byte buffers
	"compress/gzip" // for gzip compression
	"encoding/json" // for regular expressions
	"fmt"           // for formatted I/O
	"io"            // for reading response body
	"net/http"      // for HTTP requests
	"os"            //	 for file operations
	"path/filepath" // for file path manipulation
	"regexp"        // for unzipping files
	"strings"       // for string manipulation
)

func GetGithubReleasesLatest(repo string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return string(body) // Return the raw JSON response as a string
}

func GetGithubReleaseUrls(meta_data string, regex string) []string {
	// Get URLs of all release files that match the regex
	var data map[string]interface{}
	err := json.Unmarshal([]byte(meta_data), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	// Get the assets array from the JSON data
	assets, ok := data["assets"].([]interface{})
	if !ok {
		fmt.Println("assets not found in JSON data")
	}
	// Iterate over the assets and find all that match the regex
	var matchedUrls []string
	for _, asset := range assets {
		assetMap, ok := asset.(map[string]interface{})
		if !ok {
			continue
		}
		browser_download_url, ok := assetMap["browser_download_url"].(string)
		if !ok {
			continue
		}
		// Check if the asset name matches the regex
		if matched, err := regexp.MatchString(regex, browser_download_url); err == nil && matched {
			// Add the URL of the matching asset to the list
			matchedUrls = append(matchedUrls, browser_download_url)
		}
	}

	if len(matchedUrls) == 0 {
		fmt.Println("no matching assets found")
	}

	return matchedUrls
}

func DownloadToBuffer(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad Status Code: %d", resp.StatusCode)
	}

	var buff bytes.Buffer
	_, err = io.Copy(&buff, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the response body: %v", err)
	}

	return &buff, nil
}

func ExtractFileNameFromURL(url string) string {
	// Extract the file name from the URL
	re := regexp.MustCompile(`[^/]+$`)
	fileName := re.FindString(url)
	if fileName == "" {
		fmt.Println("failed to extract file name from URL")
	}
	return fileName
}

func GzipCheck(compressedData []byte) string {

	if len(compressedData) < 10 {
		fmt.Println("not a gzip file")
		return ""
	}
	// Check the gzip header
	header := compressedData[:10]
	// Check the gzip magic number

	if !bytes.HasPrefix(header, []byte{0x1f, 0x8b}) {
		// Gzip file
		fmt.Println("not a gzip file")
	}

	data := compressedData[10:]
	nullIndex := bytes.IndexByte(data, 0)
	if nullIndex == -1 {
		fmt.Println("not a gzip file name")
		return ""
	}
	// Extract the file name
	fileNameBytes := data[:nullIndex]
	fileName := string(fileNameBytes)
	return fileName
}

func ExtractGz(compressedData []byte, output_directory string) error {

	var name string
	name = GzipCheck(compressedData)

	// Decompress the data
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return err
	}
	defer reader.Close()

	// Create the output file
	outputFile, err := os.Create(output_directory + "/" + name)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}
	defer outputFile.Close()

	// Write the decompressed data to the output file
	_, err = io.Copy(outputFile, reader)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return err
	}
	return nil
}

func ExtractTar(compressedData []byte, output_directory string) error {
	// Decompress the data
	if len(compressedData) < 10 || !bytes.HasPrefix(compressedData, []byte{0x1f, 0x8b}) {
		return nil
	}

	reader := bytes.NewReader(compressedData)

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			fmt.Println("Error reading tar header:", err)
			return err
		}

		target := filepath.Join(output_directory, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			// Create the target directory if it doesn't exist
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}

			outputFile, err := os.Create(target)
			if err != nil {
				return err
			}

			_, err = io.Copy(outputFile, tarReader)
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return err
			}

			outputFile.Close()

			// Apply the file permissions from the tar header
			if err := os.Chmod(target, os.FileMode(header.Mode)); err != nil {
				return err
			}
		default:

			fmt.Println("Unknown type:", header.Typeflag, header.Name)
		}
	}

	return nil
}

func ExtractZip(compressedData []byte, output_directory string) error {
	// Create a new zip reader
	reader := bytes.NewReader(compressedData)
	zipReader, err := zip.NewReader(reader, int64(len(compressedData)))
	if err != nil {
		fmt.Println("Error creating zip reader:", err)
		return err
	}

	// Iterate through the files in the zip archive
	for _, file := range zipReader.File {
		target := filepath.Join(output_directory, file.Name)

		if file.FileInfo().IsDir() {
			// Create the directory if it doesn't exist
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				fmt.Println("Error creating directory:", err)
				return err
			}
			continue
		}

		// Create the target directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(target), os.ModePerm); err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}

		outputFile, err := os.Create(target)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return err
		}

		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("Error opening file in zip:", err)
			return err
		}
		defer fileReader.Close()

		if _, err := io.Copy(outputFile, fileReader); err != nil {
			fmt.Println("Error writing to output file:", err)
			return err
		}

		outputFile.Close()

		// apply the file permissions from the zip header
		if err := os.Chmod(target, os.FileMode(file.Mode())); err != nil {
			fmt.Println("Error setting file permissions:", err)
			return err
		}
	}
	return nil
}

func EnsureDirectoryExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the directory
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}
}

func UnifiedExtractionByteStream(repo string, regex string, output_directory string) error {

	EnsureDirectoryExists(output_directory)
	meta_data := GetGithubReleasesLatest(repo)
	browser_download_url := GetGithubReleaseUrls(meta_data, regex)

	// Print each URL from the list
	for _, url := range browser_download_url {

		fmt.Println("Downloading:", url)

		buffer, err := DownloadToBuffer(url)
		if err != nil {
			fmt.Println("Error Downloading the Buffer:", err)
		}

		binaryName := ExtractFileNameFromURL(url)

		if strings.HasSuffix(binaryName, "tar.gz") {
			err = ExtractTar(buffer.Bytes(), output_directory)
			if err != nil {
				fmt.Println("Error Extracting Tar:", err)
			}
		} else if strings.HasSuffix(binaryName, "zip") {
			err = ExtractZip(buffer.Bytes(), output_directory)
			if err != nil {
				fmt.Println("Error Extracting Zip:", err)
			}
		} else if strings.HasSuffix(binaryName, "gz") {
			err = ExtractGz(buffer.Bytes(), output_directory)
			if err != nil {
				fmt.Println("Error Extracing Gz:", err)
			}

		} else {
			// if not a tar.gz or zip file, save that file as it is
			file, err := os.Create(filepath.Join(output_directory, binaryName))
			if err != nil {
				fmt.Println("Error creating file:", err)
				return err
			}
			defer file.Close()
			_, err = io.Copy(file, bytes.NewReader(buffer.Bytes()))
			if err != nil {
				fmt.Println("Error copying file:", err)
				return err
			}
		}
	}
	return nil
}

// func main() {
// 	UnifiedExtractionByteStream("threathunters-io/laurel", "86_64-glibc.tar.gz", "/home/perumalj/Code/dir_chisel")
// }

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run githubdownload_byte_stream.go <repo> <regex> <output_directory>")
		fmt.Println("For Example: go run githubdownload_byte_stream.go 'threathunters-io/laurel' '86_64-glibc.tar.gz' '/home/perumalj/Code/dir_chisel'")
		return
	}

	repo := os.Args[1]
	regex := os.Args[2]
	output_directory := os.Args[3]

	err := UnifiedExtractionByteStream(repo, regex, output_directory)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
