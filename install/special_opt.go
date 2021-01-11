package install

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"

	"github.com/romberli/log"
)

// UnTarGz a .tar.gz file.
func UnTarGz(srcFilePath string, destDirPath string) (stdErr string, err error) {
	// Create destination directory
	_, stdErr, err = Mkdir(destDirPath)
	if err != nil {
		return stdErr, err
	}

	log.Info("Start UnTarGzing file ...")
	fr, err := os.Open(srcFilePath)
	if err != nil {
		stdErr = "The src-file is not exits."
		return stdErr, err
	}

	defer fr.Close()

	// Gzip reader
	gr, err := gzip.NewReader(fr)

	// Tar reader
	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		// Check if it is diretory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create file
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			// Write data to file
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			if err != nil {
				stdErr = "Cannot Create destDirpath."
				return stdErr, err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				stdErr = "Cannot Copy."
				return stdErr, err
			}
		}
	}
	return stdErr, err
}
