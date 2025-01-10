package fileUtils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileToLines(path string) ([]string, error) {
	file, err := os.Open(path) 
  
	
    if err != nil { 
		log.Fatalf("failed to open") 
	} 
	
	defer file.Close() 
  
    scanner := bufio.NewScanner(file) 
  
    scanner.Split(bufio.ScanLines) 
    var text []string 
  
    for scanner.Scan() { 
        text = append(text, scanner.Text()) 
    } 
  

	return text, nil
}