package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mwpcheung/debgo/cmdutils"
	"github.com/mwpcheung/debgo/deb"
	"github.com/mwpcheung/debgo/debgen"
)

func main() {
	name := "debgen-deb"
	log.SetPrefix("[" + name + "] ")
	//set to empty strings because they're being overridden
	pkg := deb.NewPackage("", "", "", "")
	build := debgen.NewBuildParams()
	debgen.ApplyGoDefaults(pkg)
	fs := cmdutils.InitFlags(name, pkg, build)

	var binDir string
	var resourcesDir string
	fs.StringVar(&binDir, "binaries", "", "directory containing binaries for each architecture. Directory names should end with the architecture")
	fs.StringVar(&pkg.Architecture, "arch", "any", "Architectures [any,386,armhf,amd64,all]")
	fs.StringVar(&resourcesDir, "resources", "", "directory containing resources for this platform")
	err := cmdutils.ParseFlags(name, pkg, fs)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = build.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	//log.Printf("Resources: %v", build.Resources)
	// TODO determine this platform
	//err = bpkg.Build(build, debgen.GenBinaryArtifact)
	artifacts, err := deb.NewDebWriters(pkg)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for arch, artifact := range artifacts {
		dgen := debgen.NewDebGenerator(artifact, build)
		err = filepath.Walk(resourcesDir, func(path string, info os.FileInfo, err2 error) error {
			if info != nil && !info.IsDir() {
				rel, err := filepath.Rel(resourcesDir, path)
				if err == nil {
					dgen.OrigFiles[rel] = path
				}
				return err
			}
			return nil
		})
		if err != nil {
			log.Fatalf("%v", err)
		}

		archBinDir := filepath.Join(binDir, string(arch))

		err = filepath.Walk(archBinDir, func(path string, info os.FileInfo, err2 error) error {
			if info != nil && !info.IsDir() {
				rel, err := filepath.Rel(binDir, path)
				if err == nil {
					dgen.OrigFiles[rel] = path
				}
				return err
			}
			return nil
		})
		err = dgen.GenerateAllDefault()
		if err != nil {
			log.Fatalf("Error building for '%s': %v", arch, err)
		}
	}
}
