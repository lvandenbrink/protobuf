package graphql_test

import (
	"fmt"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity/command"
	"github.com/opsee/protobuf/opseeproto/types"
	"github.com/opsee/protobuf/plugin/graphql"
	"github.com/opsee/protobuf/plugin/graphql/scalars"
	"github.com/opsee/protobuf/test/testdata"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGenerateDessert(t *testing.T) {
	dessert := &testdata.Dessert{
		// Name
		Name: "Test dessert",
		// Sweetness
		Sweetness: 5,
	}

	fd, md := descriptor.ForMessage(dessert)

	fmt.Printf("fd: %v, md: %v\n\n", fd, md)

	// we aren't interested if generator can find its dependency, just whether the plugin generated its code correct.
	fd.Dependency = make([]string, 0)

	params := "plugins=graphql"
	request := &plugin_go.CodeGeneratorRequest{
		FileToGenerate: []string{*fd.Name},
		Parameter:      &params,
		ProtoFile:      []*descriptor.FileDescriptorProto{fd},
	}

	response := command.Generate(request)

	if assert.NotNil(t, response) && assert.NotNil(t, response.File) {
		assert.Equal(t, 1, len(response.File))
		for _, generatedFile := range response.File {
			// search for the init where the graph type are defined. Use that that is with the pattern GraphQL<name>Type
			re := regexp.MustCompile(`(?s)func\s+init\(\)\s*{\s+(GraphQL.*?)\n}\s*?`)
			match := re.FindStringSubmatch(*generatedFile.Content)

			assert.Equal(t, 2, len(match), "> %v", match)

			fmt.Printf("MATCH: %v\n", match[0])
		}

		//
		//
		//for _, generatedFile := range response.File {
		//	fmt.Printf("\n-------------------------------------\n")
		//	fmt.Printf("%+v", *generatedFile.Content)
		//	fmt.Printf("-------------------------------------\n\n")
		//}
	}
	fmt.Println("why this?")
}


func TestGenerateTimed(t *testing.T) {
	var _ = graphql.NewGraphQL
	var _ = scalars.Timestamp

	timed := &testdata.Timed{
		UpdatedAt: &types.Timestamp{Seconds: 5, Nanos: 500,},
	}

	fd, md := descriptor.ForMessage(timed)

	fmt.Printf("fd: %v, md: %v\n\n", fd, md)


	timestamp := &types.Timestamp{}
	timestampDescriptor,_ := descriptor.ForMessage(timestamp)

	// we aren't interested if generator can find its dependency, just whether the plugin generated its code correct.
	fd.Dependency = make([]string, 0)
	timestampDescriptor.Dependency = make([]string, 0)

	params := "plugins=graphql,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	request := &plugin_go.CodeGeneratorRequest{
		FileToGenerate: []string{*fd.Name},
		Parameter:      &params,
		ProtoFile:      []*descriptor.FileDescriptorProto{timestampDescriptor, fd},
	}

	response := command.Generate(request)

	if assert.NotNil(t, response) && assert.NotNil(t, response.File) {
		assert.Equal(t, 1, len(response.File))
		for _, generatedFile := range response.File {
			// search for the init where the graph type are defined. Use that that is with the pattern GraphQL<name>Type
			re := regexp.MustCompile(`(?s)func\s+init\(\)\s*{\s+(GraphQL.*?)\n}\s*?`)
			match := re.FindStringSubmatch(*generatedFile.Content)

			if !assert.Equal(t, 2, len(match), "> %v", match) {
				fmt.Printf("%v\n", *generatedFile.Content)
				continue
			}

			fmt.Printf("MATCH: \n%v\n", match[0])
		}
	}
	fmt.Println("why this?")
}