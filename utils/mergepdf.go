package utils

import (
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePdfFile(recvFiles map[int]string) {
	inFiles := []string{}
	for _, filename := range recvFiles {
		inFiles = append(inFiles, "./temp/"+filename)
	}
	pdfcpu.MergeCreateFile(inFiles, "./output/merged.pdf", nil)
}