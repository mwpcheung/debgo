package deb_test

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/laher/argo/ar"

	"github.com/mwpcheung/debgo/deb"
)

func XTest_parse(t *testing.T) {
	rdr, err := os.Open(filepath.Join(".", "testpkg_0.0.2_amd64.deb"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	pkg, err := deb.DebParseMetadata(rdr)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("Package: %+v", pkg)
}

//Reading an ar archive ...
func XTest_learning_reading_ar(t *testing.T) {
	rdr, err := os.Open(filepath.Join(deb.DistDirDefault, "testpkg_0.0.2_amd64.deb"))
	if err != nil {
		t.Fatalf("%v", err)
	}

	arr, err := ar.NewReader(rdr)
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Iterate through the files in the archive.
	for {
		hdr, err := arr.Next()
		if err == io.EOF {
			// end of ar archive
			break
		}
		if err != nil {
			t.Fatalf("%v", err)
		}
		t.Logf("File %s:\n", hdr.Name)
		if strings.HasSuffix(hdr.Name, ".tar.gz") {
			// TODO

		} else if strings.HasSuffix(hdr.Name, ".tar") {
			// TODO
		} else if hdr.Name == "debian-binary" {
			b, err := ioutil.ReadAll(arr)
			if err != nil {
				t.Fatalf("%v", err)
			}
			t.Logf("Debian binary contents: %s", string(b))
		} else {
			t.Logf("Unsupported file %s:\n", hdr.Name)
		}
		/*
			if _, err := io.Copy(os.Stdout, arr); err != nil {
				t.Fatalf("%v", err)
			}
		*/
	}

}
