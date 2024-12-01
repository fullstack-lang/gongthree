package gongthree

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongthree/dist/ng-github.com-fullstack-lang-gongthree
var NgDistNg embed.FS
