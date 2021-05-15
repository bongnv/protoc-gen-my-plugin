package main

import (
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	log.Println("protoc-gen-my-plugin is called")
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			if err := generateFile(plugin, file); err != nil {
				return err
			}
		}

		return nil
	})
}
