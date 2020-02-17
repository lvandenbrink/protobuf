package graphql_test

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/opsee/protobuf/plugin/graphql"
	"github.com/opsee/protobuf/test/testdata"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func Test_Generate(t *testing.T) {
	person := &testdata.Dessert{
		Name:      "test",
		Sweetness: 2,
	}
	fd, md := descriptor.ForMessage(person)

	fmt.Println(fd)
	fmt.Println(md)
	//
	//fmt.Println(fd.GetOptions())
	//if fd.Options == nil {
	//	fd.Options = &descriptor.FileOptions{}
	//}

	//fd.Options.XXX_InternalExtensions.p.extensionMap[64040] = &proto.ExtensionDesc{
	//	ExtendedType:  (*descriptor.FileOptions)(nil),
	//	ExtensionType: (*bool)(nil),
	//	Field:         64040,
	//	Name:          "opseeproto.graphql",
	//	Tag:           "varint,64040,opt,name=graphql",
	//}

	fileDescriptor := &generator.FileDescriptor{
		FileDescriptorProto: fd,
	}

	graphQL := graphql.NewGraphQL()
	graphQL.Generate(fileDescriptor)

	fmt.Println(fileDescriptor.String())


}

type mockFileDescriptorProto struct {
	descriptor.FileOptions
}

func (*mockFileDescriptorProto) ExtensionRangeArray() []proto.ExtensionRange {
	return []proto.ExtensionRange{
		{Start: 100, End: 199},
	}
}
func (*mockFileDescriptorProto) GetExtensions() *[]byte {
	extensions := make([]byte, 0)
	return &extensions
}

const (
	bProto = `
syntax = "proto3";
package test.alpha;
option go_package = "package/alpha";
import "beta/b.proto";
message M { test.beta.M field = 1; }`

	aProto = `
syntax = "proto3";
package test.beta;
// no go_package option
message M {}`
)

func TestT(t *testing.T) {
	workdir, err := ioutil.TempDir("", "proto-test")
	if err != nil {
		t.Fatal(err)
	}
	// defer os.RemoveAll(workdir)

	for _, dir := range []string{"alpha", "beta", "out"} {
		if err := os.MkdirAll(filepath.Join(workdir, dir), 0777); err != nil {
			t.Fatal(err)
		}
	}
	if err := ioutil.WriteFile(filepath.Join(workdir, "alpha", "a.proto"), []byte(aProto), 0666); err != nil {
		t.Fatal(err)
	}

	protoc(t, []string{
		"-I" + workdir,
		"--go_out=" + ":" + filepath.Join(workdir, "out"),
		filepath.Join(workdir, "alpha", "a.proto"),
	})
}

func protoc(t *testing.T, args []string) {
	cmd := exec.Command("protoc", "--plugin=protoc-gen-go="+os.Args[0])
	cmd.Args = append(cmd.Args, args...)
	// We set the RUN_AS_PROTOC_GEN_GO environment variable to indicate that
	// the subprocess should act as a proto compiler rather than a test.
	cmd.Env = append(os.Environ(), "RUN_AS_PROTOC_GEN_GO=1")
	out, err := cmd.CombinedOutput()
	if len(out) > 0 || err != nil {
		t.Log("RUNNING: ", strings.Join(cmd.Args, " "))
	}
	if len(out) > 0 {
		t.Log(string(out))
	}
	if err != nil {
		t.Fatalf("protoc: %v", err)
	}
}
