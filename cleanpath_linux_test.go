package cleanpath_test

// testCase1 is the linux test case using host specific path format.
var testCase1 = map[string]string{
	"/a/b/c:/a/b/c:/a/b/c/": "/a/b/c",
	"../a/b:/a/b:../a/b":    "/a/b",
	"/a:/b:/c:/b:/a/:a":     "/a:/b:/c",
}
