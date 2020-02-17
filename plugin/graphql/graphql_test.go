package graphql_test

import (
	"fmt"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	plugin_go "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity/command"
	"io/ioutil"
	"os"
	"testing"
)

var (
	protoContent = `syntax = "proto3";
package test;

import "github.com/opsee/protobuf/opseeproto/opsee.proto";

option (opseeproto.graphql) = true;

message Person {
    string name = 1;
    int32 id = 2;
    string email = 3;
    PhoneNumber phone = 4;
}

message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
}

enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
}`
	protoDessertContent = `syntax = "proto3";

import "github.com/gogo/protobuf/gogoproto/gogo.proto";
import "github.com/opsee/protobuf/opseeproto/opsee.proto";

package flavortown.dessert;

option (opseeproto.graphql) = true;


// A delicious dessert dish on the menu
message Dessert {
  // The name of the dish
  string name = 1;
  // How sweet is the dish, an integer between 0 and 10
  int32 sweetness = 2;
}`
)

func TestGenerate(t *testing.T) {
	file, cleanup := createTestFile(t, protoDessertContent)
	defer cleanup()

	var msg *descriptor.DescriptorProto
	fd, _ := descriptor.ForMessage(msg)

	fd.Name = &file

	// params := "plugins=graphql+bitflags,Mopsee/protobuf/opsee.proto=github.com/opsee/protobuf/opseeproto,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	params := "plugins=graphql,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:dessert"
	request := &plugin_go.CodeGeneratorRequest{
		FileToGenerate: []string{file},
		Parameter:      &params,
		ProtoFile:      []*descriptor.FileDescriptorProto{fd},
	}

	response := command.Generate(request)

	fmt.Println(response)

	for _, generatedFile := range response.File{
		fmt.Printf("\n-------------------------------------\n")
		fmt.Printf("%+v", *generatedFile.Content)
		fmt.Printf("-------------------------------------\n\n")
	}


	fmt.Println("why this?")
}

func TestGenerateDessert(t *testing.T) {
	dessert := &TestDessert{
		Name:                 "Test dessert",
		Sweetness:            5,
	}

	fd, md := descriptor.ForMessage(dessert)

	fmt.Printf("fd: %v, md: %v", fd, md)

	// we aren't interested if generator can find its dependency, just whether the plugin generated its code correct.
	fd.Dependency = make([]string, 0)

	params := "plugins=graphql,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:dessert"
	request := &plugin_go.CodeGeneratorRequest{
		FileToGenerate: []string{*fd.Name},
		Parameter:      &params,
		ProtoFile:      []*descriptor.FileDescriptorProto{fd},
	}


	response := command.Generate(request)

	fmt.Println(response)

	for _, generatedFile := range response.File{
		fmt.Printf("\n-------------------------------------\n")
		fmt.Printf("%+v", *generatedFile.Content)
		fmt.Printf("-------------------------------------\n\n")
	}


	fmt.Println("why this?")
}

/*

var GraphQLDessertType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLDessertType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "Dessert",
		Description: "A delicious dessert dish on the menu",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"name": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The name of the dish",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Dessert)
						if ok {
							return obj.Name, nil
						}
						inter, ok := p.Source.(DessertGetter)
						if ok {
							face := inter.GetDessert()
							if face == nil {
								return nil, nil
							}
							return face.Name, nil
						}
						return nil, fmt.Errorf("field name not resolved")
					},
				},
				"sweetness": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "How sweet is the dish, an integer between 0 and 10",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Dessert)
						if ok {
							return obj.Sweetness, nil
						}
						inter, ok := p.Source.(DessertGetter)
						if ok {
							face := inter.GetDessert()
							if face == nil {
								return nil, nil
							}
							return face.Sweetness, nil
						}
						return nil, fmt.Errorf("field sweetness not resolved")
					},
				},
			}
		}),
	})
}
 */
func createTestFile(t *testing.T, data interface{}) (string, func()) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %v", err)
	}

	_, err = fmt.Fprint(tmpFile, data)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
		_ = os.Remove(tmpFile.Name())
	}

	return tmpFile.Name(), func() {
		// Remember to clean up the file afterwards
		err := os.Remove(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to cleanup test file: %v", err)
		}
	}
}
