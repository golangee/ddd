package parser

// Presentation contains the driver adapters.
type Presentation struct{
	Rest *Rest `("rest" @@)?`
}

type Rest struct{
	Version RestMajorVersion `@@*`
}


type RestMajorVersion struct{
	Version SemVer `@@ "{" "}"`
}