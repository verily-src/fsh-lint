package fsh_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/verily-src/fsh-lint/internal/fsh"
	"github.com/verily-src/fsh-lint/internal/fsh/types"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

//go:embed resources/TestValueSet_Want.json
var ValueSetWant string

//go:embed resources/VerilyTestProfile_Want.json
var ProfileWant string

//go:embed resources/VerilyTestProfileMultilineDescription_Want.json
var ProfileMultilineDescriptionWant string

//go:embed resources/TestCodeSystem_Want.json
var CodeSystemWant string

//go:embed resources/TestInstance_Want.json
var InstanceWant string

//go:embed resources/TestExtension_Want.json
var ExtensionWant string

//go:embed resources/TestValueSet.fsh
var ValueSetFSHData string

//go:embed resources/VerilyTestProfile.fsh
var ProfileFSHData string

//go:embed resources/TestProfileMultilineDescription.fsh
var MultilineFSHProfileData string

//go:embed resources/TestCodeSystem.fsh
var CodeSystemFSHData string

//go:embed resources/TestInstance.fsh
var InstanceFSHData string

//go:embed resources/TestExtension.fsh
var ExtensionFSHData string

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		fshData string
		want    *types.FSHDocument
		wantErr error
	}{
		{
			name:    "valid valueset",
			fshData: ValueSetFSHData,
			want:    parseDocJSON(ValueSetWant, t),
		},
		{
			name:    "valid profile",
			fshData: ProfileFSHData,
			want:    parseDocJSON(ProfileWant, t),
		},
		{
			name:    "multiline profile description",
			fshData: MultilineFSHProfileData,
			want:    parseDocJSON(ProfileMultilineDescriptionWant, t),
		},
		{
			name:    "valid code system",
			fshData: CodeSystemFSHData,
			want:    parseDocJSON(CodeSystemWant, t),
		},
		{
			name:    "valid instance",
			fshData: InstanceFSHData,
			want:    parseDocJSON(InstanceWant, t),
		},
		{
			name:    "valid extension",
			fshData: ExtensionFSHData,
			want:    parseDocJSON(ExtensionWant, t),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fsh.Parse(tt.fshData)
			if got, want := err, tt.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Parse() err = %v, want %v", got, want)
			}

			if diff := cmp.Diff(*got, *tt.want); diff != "" {
				t.Errorf("Parse() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func parseDocJSON(jsonData string, t *testing.T) *types.FSHDocument {
	var fd types.FSHDocument

	err := json.Unmarshal([]byte(jsonData), &fd)
	if err != nil {
		t.Fatalf("Error parsing json: %v", err.Error())
	}
	return &fd
}
