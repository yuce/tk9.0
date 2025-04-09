![sshot-png](sshot.png)

[![LiberaPay](https://liberapay.com/assets/widgets/donate.svg)](https://liberapay.com/jnml/donate)
[![receives](https://img.shields.io/liberapay/receives/jnml.svg?logo=liberapay)](https://liberapay.com/jnml/donate)
[![patrons](https://img.shields.io/liberapay/patrons/jnml.svg?logo=liberapay)](https://liberapay.com/jnml/donate)

[![Go Reference](https://pkg.go.dev/badge/modernc.org/tk9.0/appbundle.svg)](https://pkg.go.dev/modernc.org/tk9.0/appbundle)

# appbundle

Command appbundle creates [MacOS Application_bundles].

## Installation

     $ go install modernc.org/tk9.0/appbundle@latest

## Usage

     $ appbundle [-dark] [-name <app_name>] [-version <ver>] [-id <bundle_id>] [-sign <identity>] path-to-binary

If any of the above flags/arguments are missing appbundler will start a GUI
allowing to fill the options.

[MacOS Application_bundles]: https://en.wikipedia.org/wiki/Bundle_(macOS)#Application_bundles
