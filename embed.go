package grape

import "embed"

// StaticFS 静态资源
//go:embed static
var StaticFS embed.FS

// Helloworld Helloworld
//go:embed static/helloworld.txt
var Helloworld string
