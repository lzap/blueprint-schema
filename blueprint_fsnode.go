package blueprint

import "github.com/invopop/jsonschema"

// fsnodes: # new feature (was file/dir)
//   - path: "/var/tmp/dir"
//     type: "dir" # file as the default (file, dir)
//     state: "present" # the default
//     mode: 0755 # 0700 default
//     user: "lzap" # 0 default
//     group: "lzap" # 0 default
//     ensure_parents: true # false default
//   - path: "/var/tmp/dir/file1"
//     type: "file"
//     state: "present"
//     mode: 0644
//     user: "lzap"
//     group: "lzap"
//     contents:
//     base64: "LS0t" # useful for binaries
//   - path: "/var/tmp/dir/file2"
//     contents:
//     text: "test" # useful for text
type FSNode struct {
	// Path is the absolute path to the file or directory.
	Path string `json:"path" jsonschema:"required,pattern=^/"`

	// Type is the type of the file system node, one of: file, dir.
	Type FSNodeType `json:"type,omitempty" jsonschema:"default=file,enum=file,enum=dir"`

	// State is the state of the file system node, one of: present, absent.
	State FSNodeState `json:"state,omitempty" jsonschema:"default=present,enum=present,enum=absent"`

	// Mode is the file system node permissions. Defaults to 0644 for files and 0755 for directories.
	Mode int `json:"mode,omitempty"`

	// User is the file system node owner. Defaults to root.
	User string `json:"user,omitempty" jsonschema:"default=root"`

	// Group is the file system node group. Defaults to root.
	Group string `json:"group,omitempty" jsonschema:"default=root"`

	// EnsureParents is a boolean that determines if the parent directories should be created if they do not exist.
	EnsureParent bool `json:"ensure_parents,omitempty" jsonschema:"default=false"`

	// Contents is the file system node contents. When not present, an empty file is created.
	Contents *FSNodeContents `json:"contents,omitempty" jsonschema:"nullable"`
}

type FSNodeContents struct {
	// Base64-encoded file contents. Useful for binaries.
	Base64 string `json:"base64,omitempty"`

	// Plain text file contents.
	Text string `json:"text,omitempty"`
}

// JSONSchemaExtend can be used to extend the generated JSON schema from Go struct tags
func (FSNode) JSONSchemaExtend(s *jsonschema.Schema) {
	// If type is file...
	pIf := jsonschema.NewProperties()
	pIf.Set("type", &jsonschema.Schema{
		Const: "file",
	})
	s.If = &jsonschema.Schema{
		Properties: pIf,
	}

	// ...then require contents
	pThen := jsonschema.NewProperties()
	pThen.Set("contents", &jsonschema.Schema{
		OneOf: []*jsonschema.Schema{
			{
				Required: []string{"base64"},
				Title:    "fsnodes_base64",
			},
			{
				Required: []string{"text"},
				Title:    "fsnodes_text",
			},
		},
	})
	// ...and set default mode for files
	pThen.Set("mode", &jsonschema.Schema{
		Default: 0644,
	})
	s.Then = &jsonschema.Schema{
		Properties: pThen,
	}

	// ...else set default mode for directories
	pElse := jsonschema.NewProperties()
	pElse.Set("mode", &jsonschema.Schema{
		Default: 0755,
	})
	s.Else = &jsonschema.Schema{
		Properties: pElse,
	}
}
