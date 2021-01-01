package debgen_test

import (
	"log"

	"github.com/mwpcheung/debgo/deb"
	"github.com/mwpcheung/debgo/debgen"
)

func Example_genSourcePackage() {

	pkg := deb.NewPackage("testpkg", "0.0.2", "me <a@me.org>", "Dummy package for doing nothing\n")
	build := debgen.NewBuildParams()
	build.IsRmtemp = false
	debgen.ApplyGoDefaults(pkg)
	spkg := deb.NewSourcePackage(pkg)
	err := build.Init()
	if err != nil {
		log.Fatalf("Error initializing dirs: %v", err)
	}
	spgen := debgen.NewSourcePackageGenerator(spkg, build)
	spgen.ApplyDefaultsPureGo()
	sourcesDestinationDir := pkg.Name + "_" + pkg.Version
	sourceDir := ".."
	sourcesRelativeTo := debgen.GetGoPathElement(sourceDir)
	spgen.OrigFiles, err = debgen.GlobForSources(sourcesRelativeTo, sourceDir, debgen.GlobGoSources, sourcesDestinationDir, []string{build.TmpDir, build.DestDir})
	if err != nil {
		log.Fatalf("Error resolving sources: %v", err)
	}
	err = spgen.GenerateAllDefault()

	if err != nil {
		log.Fatalf("Error building source: %v", err)
	}

	// Output:
	//
}
