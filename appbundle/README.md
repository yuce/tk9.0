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

## Notes

- It is best to encapsulate `<app_name>` in double quotes. This ensures that the correct 
value for the `CFBundleName` key is set if there are spaces in the app name.
- Currently, the app bundle icon cannot be set via the appbundle command. You may set it 
by setting the value for the `CFBundleIconFile` key in the `Info.plist` file or by using 
the [Get Info method](https://discussions.apple.com/thread/255174964?answerId=259623836022&sortBy=rank#259623836022). 
If setting the `CFBundleIconFile` key, you will need to create a [icns](https://en.wikipedia.org/wiki/Apple_Icon_Image_format) file and place 
it in the `Resources` directory of your App Bundle before setting the key.
- If no valid Apple Developer credential is provided for the `-sign` parameter, the appbundle 
will do a temporary app signing which is not suitable for general distribution of an app.

