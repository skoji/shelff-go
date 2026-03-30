package shelff

import "os"

// ReadMetadata returns metadata for a PDF.
// If a sidecar exists, returns its content.
// If no sidecar exists, returns a minimal SidecarMetadata with
// dc:title set to the PDF filename (without .pdf extension).
// Unlike ReadSidecar, this never returns nil for an existing PDF.
// Returns ErrPDFNotFound if the PDF does not exist.
func ReadMetadata(pdfPath string) (*SidecarMetadata, error) {
	info, err := os.Stat(pdfPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrPDFNotFound
		}
		return nil, err
	}
	if info.IsDir() {
		return nil, ErrPDFNotFound
	}

	sidecar, err := ReadSidecar(pdfPath)
	if err != nil {
		return nil, err
	}
	if sidecar != nil {
		return sidecar, nil
	}

	return &SidecarMetadata{
		SchemaVersion: SchemaVersion,
		Metadata: DublinCore{
			Title: pdfTitleFromPath(pdfPath),
		},
	}, nil
}
