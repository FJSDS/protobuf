package typedeclimport

import subpkg "github.com/FJSDS/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
