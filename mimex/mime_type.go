/*
 * @PackageName: mimex
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 14:40
 */

package mimex

import (
	"path/filepath"
	"strings"
)

const (
	MimeTypeJPEG = "image/jpeg"
	MimeTypePNG  = "image/png"
	MimeTypeGIF  = "image/gif"
	MimeTypeSVG  = "image/svg+xml"
	MimeTypePDF  = "application/pdf"
	// ".m3u8":
	MimeTypeM3U8 = "application/x-mpegURL"
	// ".ts":
	MimeTypeTS = "video/mp2ts"
)

var (
	fileTypeExts = map[string]string{
		MimeTypeJPEG: ".jpg",
		MimeTypePNG:  ".png",
		MimeTypeGIF:  ".gif",
		MimeTypePDF:  ".pdf",
		MimeTypeSVG:  ".svg",
		MimeTypeM3U8: ".m3u8",
		MimeTypeTS:   ".ts",
	}
	fileTypeTypes = map[string]string{
		".jpg":  MimeTypeJPEG,
		".jpeg": MimeTypeJPEG,
		".png":  MimeTypePNG,
		".gif":  MimeTypeGIF,
		".pdf":  MimeTypePDF,
		".svg":  MimeTypeSVG,
		".m3u8": MimeTypeM3U8,
		".ts":   MimeTypeTS,
	}
)

// MimeType .
func MimeType(name string) string {
	return fileTypeTypes[filepath.Ext(name)]
}

// ImageMimeType .
func ImageMimeType(name string) (string, string) {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".jpg", ".jpeg":
		return ext, MimeTypeJPEG
	case ".png":
		return ext, MimeTypePNG
	}
	return ext, ""
}
