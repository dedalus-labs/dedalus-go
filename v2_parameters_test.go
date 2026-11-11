package dedalus_test

import (
	"encoding/json"
	"reflect"
	"testing"

	sdk "github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/dedalus-go/packages/param"
)

func TestV2ParametersKeepWireShape(t *testing.T) {
	cases := []struct {
		name  string
		value any
		want  string
	}{
		{"machine policy", sdk.MachineAutoresizingUpdateParams{
			MachineID: "fixture", CPU: sdk.AutoresizingSwitchOff, Memory: sdk.AutoresizingSwitchOn, MemoryMaxMiB: sdk.Int(16384),
		}, `{"cpu":"off","memory":"on","memory_max_mib":16384}`},
		{"organization policy", sdk.OrganizationAutoresizingUpdateParams{
			CPU: sdk.AutoresizingSwitchOff, Memory: sdk.AutoresizingSwitchOn, MemoryMaxMiB: sdk.Int(16384),
		}, `{"cpu":"off","memory":"on","memory_max_mib":16384}`},
		{"machine policy omitted limits", sdk.MachineAutoresizingUpdateParams{
			MachineID: "fixture", CPU: sdk.AutoresizingSwitchOff,
		}, `{"cpu":"off"}`},
		{"organization policy cleared limits", sdk.OrganizationAutoresizingUpdateParams{
			CPUMaxVCPU: param.Null[float64](), MemoryMaxMiB: param.Null[int64](),
		}, `{"cpu_max_vcpu":null,"memory_max_mib":null}`},
		{"required SSH key", sdk.SSHSessionCreateParams{
			PublicKey: "ssh-ed25519 fixture",
		}, `{"public_key":"ssh-ed25519 fixture"}`},
		{"required zero value", sdk.SSHSessionCreateParams{}, `{"public_key":""}`},
		{"resource casing", sdk.CreateParams{
			MemoryMiB: sdk.Int(4096), StorageGiB: sdk.Int(10), VCPU: sdk.Float(1),
		}, `{"memory_mib":4096,"storage_gib":10,"vcpu":1}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.value)
			if err != nil {
				t.Fatal(err)
			}
			var got, want any
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(tc.want), &want); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("JSON: got %s, want %s", data, tc.want)
			}
		})
	}
}
