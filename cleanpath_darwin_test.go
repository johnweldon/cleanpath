package cleanpath_test

// testCase1 is the darwin (mac osx) test case using host specific path format.
var testCase1 = map[string]string{
	"/a/b/c:/a/b/c:/a/b/c/":             "/a/b/c",
	"../a/b:/a/b:../a/b":                "/a/b",
	"/a:/b:/c:/b:/a/:a":                 "/a:/b:/c",
	"/has/path with spaces/bin:/a/b:/a": "/has/path with spaces/bin:/a/b:/a",
}
