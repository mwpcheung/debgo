package deb_test

import (
	"log"
	"testing"

	"debgo/deb"
)

func Example_buildDevPackage() {

	pkg := deb.NewPackage("testpkg", "0.0.2", "me", "A package\ntestpkg is a lovel package with many wow")
	buildFunc := func(dpkg *deb.Package) error {
		// Generate files here.
		return nil
	}
	dpkg := deb.NewDevPackage(pkg)
	err := buildFunc(dpkg)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func Test_buildDevPackage(t *testing.T) {

	pkg := deb.NewPackage("testpkg", "0.0.2", "me", "A package\ntestpkg is a lovel package with many wow")
	buildFunc := func(dpkg *deb.Package) error {
		// Generate files here.
		return nil
	}
	dpkg := deb.NewDevPackage(pkg)
	err := buildFunc(dpkg)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
