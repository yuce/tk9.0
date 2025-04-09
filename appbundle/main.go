// Copyright 2025 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command appbundler creates [MacOS Application_bundles].
//
// # Installation
//
//	$ go install modernc.org/tk9.0/appbundle@latest
//
// # Usage
//
//	$ appbundle [-dark] [-name <app_name>] [-version <ver>] [-id <bundle_id>] [-sign <identity>] path-to-binary
//
// If any of the above flags/arguments are missing appbundler will start a GUI
// allowing to fill the options.
//
// [MacOS Application_bundles]: https://en.wikipedia.org/wiki/Bundle_(macOS)#Application_bundles
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"modernc.org/fileutil"
	"modernc.org/opt"
	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

var (
	//go:embed logo.png
	icon []byte

	bin         string // path-to-binary
	dark        bool   // -dark
	goos        = runtime.GOOS
	optID       string // -id
	optIdentity = "-"  // -sign
	optName     string // -name
	optVersion  string // -version
)

func fail(rc int, s string, args ...any) {
	s = fmt.Sprintf("FAIL: "+s, args...)
	fmt.Fprintln(os.Stderr, strings.TrimSpace(s))
	os.Exit(rc)
}

func main() {
	set := opt.NewSet()
	set.Arg("name", false, func(opt, arg string) error { optName = arg; return nil })
	set.Arg("version", false, func(opt, arg string) error { optVersion = arg; return nil })
	set.Arg("id", false, func(opt, arg string) error { optID = arg; return nil })
	set.Arg("sign", false, func(opt, arg string) error { optIdentity = arg; return nil })
	set.Opt("dark", func(opt string) error { dark = true; return nil })
	if err := set.Parse(os.Args[1:], func(arg string) error {
		if bin != "" {
			return fmt.Errorf("multiple path-to-binary arguments, expected at most one")
		}

		bin = arg
		return nil
	}); err != nil {
		fail(2, "err=%v", err)
	}

	if optID != "" && (optIdentity != "" || goos != "darwin") && optName != "" && optVersion != "" && bin != "" {
		if _, err := run(optID, optIdentity, optName, optVersion, bin); err != nil {
			fail(1, "err=%v", err)
		}

		return
	}

	switch {
	case dark:
		ActivateTheme("azure dark")
	default:
		ActivateTheme("azure light")
	}

	var check func()
	sty := Opts{Padx("2m"), Pady("2m")}

	nmEntry := TEntry(Textvariable(optName))
	Tooltip(nmEntry, "Name of your application")
	Grid(TLabel(Txt("Name")), Row(0), Column(0), Sticky("e"), sty)
	Grid(nmEntry, Row(0), Column(1), sty)
	Bind(nmEntry, "<FocusOut>", Command(func() { check() }))

	verEntry := TEntry(Textvariable(optVersion))
	Tooltip(verEntry, "Version number")
	Grid(TLabel(Txt("Version")), Row(1), Column(0), Sticky("e"), sty)
	Grid(verEntry, Row(1), Column(1), sty)
	Bind(verEntry, "<FocusOut>", Command(func() { check() }))

	idEntry := TEntry(Textvariable(optID))
	Tooltip(idEntry, "Unique bundle identifier")
	Grid(TLabel(Txt("ID")), Row(2), Column(0), Sticky("e"), sty)
	Grid(idEntry, Row(2), Column(1), sty)
	Bind(idEntry, "<FocusOut>", Command(func() { check() }))

	identityEntry := TEntry(Textvariable(optIdentity))
	Tooltip(identityEntry, "Signing identity, used only on MacOS")
	Grid(TLabel(Txt("Identity")), Row(3), Column(0), Sticky("e"), sty)
	Grid(identityEntry, Row(3), Column(1), sty)
	Bind(identityEntry, "<FocusOut>", Command(func() { check() }))
	if goos != "darwin" {
		identityEntry.Configure(State("readonly"))
	}

	binEntry := TEntry(Textvariable(bin))
	Tooltip(binEntry, "Path of your compiled Go binary")
	Grid(TLabel(Txt("Binary")), Row(4), Column(0), Sticky("e"), sty)
	Grid(binEntry, Row(4), Column(1), sty)
	Grid(Tooltip(TButton(Txt("Browse"), Command(func() {
		if r := GetOpenFile(
			Title("Choose compiled Go binary"),
			Parent(App),
			Filetypes([]FileType{{TypeName: "executable Go file", Extensions: []string{"*"}}}),
		); len(r) != 0 {
			binEntry.Configure(Textvariable(r[0]))
			Focus(binEntry)
		}
	})), "Open the file chooser dialog"), Row(4), Column(2), sty)
	Bind(binEntry, "<FocusOut>", Command(func() { check() }))

	bundle := TButton(Txt("Bundle"), Style("Accent.TButton"), Command(func() {
		s, err := run(idEntry.Textvariable(), identityEntry.Textvariable(), nmEntry.Textvariable(), verEntry.Textvariable(), binEntry.Textvariable())
		if err != nil {
			MessageBox(Detail(fmt.Sprintf("Error creating applicaiton bundle:\n\n%v", err)), Icon("error"), Parent(App), Title("Fail"), Type("ok"))
			return
		}

		MessageBox(Detail(fmt.Sprintf("%s\n\nClick OK to quit", s)), Icon("info"), Parent(App), Title("Success"), Type("ok"))
		Destroy(App)
	}))
	Bind(bundle, "<Enter>", Command(func() { Focus(bundle) }))

	check = func() {
		was := bundle.State()
		new := "disabled"

		defer func() {
			if new != was {
				bundle.Configure(State(new))
			}
		}()

		if idEntry.Textvariable() != "" && (identityEntry.Textvariable() != "" || goos != "darwin") &&
			nmEntry.Textvariable() != "" && verEntry.Textvariable() != "" && binEntry.Textvariable() != "" {
			new = "normal"
		}
	}
	check()

	Grid(Tooltip(bundle, "Create the application bundle"), Row(6), Column(1), Sticky("e"), sty)
	Grid(Tooltip(TButton(Txt("Quit"), Command(func() { Destroy(App) })), "Quit the application"), Row(6), Column(2), sty)

	s := "uknown version"
	if bi, _ := debug.ReadBuildInfo(); bi != nil {
		s = fmt.Sprintf("%s\n%s", bi.Main.Path, bi.Main.Version)
	}
	Grid(Tooltip(Label(Image(NewPhoto(Data(icon)))), s), Row(0), Column(2), Rowspan(4))
	App.SetResizable(false, false)
	App.Wait()
}

func run(appID, codeSignIdentity, appName, appVersion, bin string) (r string, err error) {
	absBin, err := filepath.Abs(bin)
	if err != nil {
		return "", err
	}

	fi, err := os.Stat(bin)
	if err != nil {
		return "", err
	}

	if fi.IsDir() {
		return "", fmt.Errorf("%s is a directory", bin)
	}

	// Create the .app directory structure
	appDir := fmt.Sprintf("%s.app", appName)
	contentsDir := filepath.Join(appDir, "Contents")
	macOSDir := filepath.Join(contentsDir, "MacOS")
	resourceDir := filepath.Join(contentsDir, "Resources")
	if err := os.MkdirAll(macOSDir, 0770); err != nil {
		return "", err
	}

	if err := os.MkdirAll(resourceDir, 0770); err != nil {
		return "", err
	}

	// Copy the binary
	if _, err := fileutil.CopyFile(os.DirFS(filepath.Dir(absBin)), filepath.Join(macOSDir, appName), filepath.Base(absBin), nil); err != nil {
		return "", err
	}

	// Create the Info.plist file
	b := bytes.NewBuffer(nil)
	fmt.Fprintf(b, `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>%s</string>
    <key>CFBundleIdentifier</key>
    <string>%s</string>
    <key>CFBundleName</key>
    <string>%[1]s</string>
    <key>CFBundleVersion</key>
    <string>%[3]s</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleShortVersionString</key>
    <string>%[3]s</string>
    <key>NSPrincipalClass</key>
    <string>NSApplication</string>
</dict>
</plist>
`, appName, appID, appVersion,
	)
	if err := os.WriteFile(filepath.Join(contentsDir, "Info.plist"), b.Bytes(), 0660); err != nil {
		return "", err
	}

	if goos == "darwin" {
		out, err := exec.Command("codesign", "--force", "--deep", "--sign", codeSignIdentity, appDir).CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("err=%v\nout=%s", err, out)
		}
	}

	app := filepath.Join(macOSDir, appName)
	if fi, err = os.Stat(app); err != nil {
		return "", err
	}

	// Make the binary executable
	if err := os.Chmod(app, fi.Mode()|0110); err != nil {
		return "", err
	}

	return fmt.Sprintf("Sucessfully created %s", appDir), nil
}
