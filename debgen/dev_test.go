package debgen_test

import (
	"log"

	"github.com/mwpcheung/debgo/deb"
	"github.com/mwpcheung/debgo/debgen"
)

func Example_genDevPackage() {
	pkg := deb.NewPackage("testpkg", "0.0.2", "me", "Dummy package for doing nothing\n")

	ddpkg := deb.NewDevPackage(pkg)
	build := debgen.NewBuildParams()
	build.IsRmtemp = false
	build.Init()
	var err error
	mappedFiles, err := debgen.GlobForGoSources(".", []string{build.TmpDir, build.DestDir})
	if err != nil {
		log.Fatalf("Error building -dev: %v", err)
	}

	err = debgen.GenDevArtifact(ddpkg, build, mappedFiles)
	if err != nil {
		log.Fatalf("Error building -dev: %v", err)
	}

	// Output:
	//
}
