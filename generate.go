package main

import "google.golang.org/protobuf/compiler/protogen"

func generateFile(p *protogen.Plugin, f *protogen.File) error {
	// Skip generating file if there is no message.
	if len(f.Messages) == 0 {
		return nil
	}

	filename := f.GeneratedFilenamePrefix + "_my_plugin.pb.go"
	g := p.NewGeneratedFile(filename, f.GoImportPath)
	g.P("// Code generated by protoc-gen-my-plugin. DO NOT EDIT.")
	g.P()
	g.P("package ", f.GoPackageName)
	g.P()

	// generate factory functions
	for _, m := range f.Messages {
		msgName := m.GoIdent.GoName
		g.P("// New creates a new instance of ", msgName, ".")
		g.P("func (", msgName, ") New() *", msgName, " {")
		g.P("return &", msgName, "{}")
		g.P("}")
	}

	return nil
}
