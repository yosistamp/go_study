package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type Structure struct {
	Category   string      `json:"category"`
	Sections   interface{} `json:"sections"`
	Node       []Structure `json:"node"`
	LastUpdate string      `json:"lastupdate"`
}

func getMdxSections(filePath string) (interface{}, string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []interface{}{}, ""
	}

	re := regexp.MustCompile(`(?s)sections:\s*(\[.*?\])`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) > 1 {
		var sections interface{}
		err := json.Unmarshal([]byte(matches[1]), &sections)
		if err != nil {
			return []interface{}{}, ""
		}

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return sections, ""
		}

		lastUpdate := fileInfo.ModTime().Format("2006年01月02日")
		return sections, lastUpdate
	}

	return []interface{}{}, ""
}

func getFolderStructure(folderPath string, parentCategory string) Structure {
	folderName := filepath.Base(folderPath)
	category := folderName
	if parentCategory != "" {
		category = parentCategory + "-" + folderName
	}

	mdxPath := filepath.Join(folderPath, "contents.mdx")
	sections, lastUpdate := getMdxSections(mdxPath)

	structure := Structure{
		Category:   category,
		Sections:   sections,
		Node:       []Structure{},
		LastUpdate: lastUpdate,
	}

	subFolders, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return structure
	}

	for _, subFolder := range subFolders {
		if subFolder.IsDir() {
			subFolderPath := filepath.Join(folderPath, subFolder.Name())
			structure.Node = append(structure.Node, getFolderStructure(subFolderPath, category))
		}
	}

	return structure
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <RootFolder>")
		return
	}

	rootFolder := os.Args[1]
	subFolders, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		fmt.Println("Error reading root folder:", err)
		return
	}

	var result []Structure
	for _, subFolder := range subFolders {
		if subFolder.IsDir() {
			subFolderPath := filepath.Join(rootFolder, subFolder.Name())
			result = append(result, getFolderStructure(subFolderPath, ""))
		}
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	outputPath := filepath.Join(".", "folder_structure.json")
	err = ioutil.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON output written to", outputPath)
}
