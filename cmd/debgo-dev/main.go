package main

import (
	"github.com/laher/debgo-v0.2/cmd"
	"github.com/laher/debgo-v0.2/deb"
	"log"
)

func main() {
	name := "debgo-dev"
	log.SetPrefix("[" + name + "] ")
	//set to empty strings because they're being overridden
	pkg := deb.NewGoPackage("", "", "")

	fs := cmdutils.InitFlags(name, pkg)
	fs.StringVar(&pkg.Architecture, "arch", "all", "Architectures [any,386,armel,amd64,all]")
	ddpkg := deb.NewDevPackage(pkg)
	var sourceDir string
	var sourcesRelativeTo string
	fs.StringVar(&sourceDir, "sources", ".", "source dir")
	fs.StringVar(&sourcesRelativeTo, "sources-relative-to", "", "(optional) sources - relative to")
	err := cmdutils.ParseFlags(name, pkg, fs)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = ddpkg.Build()
	if err != nil {
		log.Fatalf("%v", err)
	}

}