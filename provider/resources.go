package twilio

import (
	"fmt"
	"path/filepath"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	twilio "github.com/twilio/terraform-provider-twilio/twilio"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	"github.com/nellisauction/pulumi-twilio/provider/pkg/version"
)

const (
	mainPkg = "twilio"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-twilio/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(twilio.Provider())

	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "twilio",
		Version:           version.Version,
		DisplayName:       "Twilio",
		Publisher:         "NellisAuction",
		PluginDownloadURL: "github://api.github.com/nellisauction/pulumi-twilio",
		Description:       "A Pulumi package for managing Twilio resources.",
		Keywords:          []string{"pulumi", "twilio", "category/infrastructure"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/nellisauction/pulumi-twilio",
		Repository:        "https://github.com/nellisauction/pulumi-twilio",
		GitHubOrg:         "twilio",
		Config: map[string]*tfbridge.SchemaInfo{
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TWILIO_API_KEY", "TWILIO_ACCOUNT_SID"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TWILIO_API_SECRET", "TWILIO_AUTH_TOKEN"},
				},
			},
			"account_sid": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TWILIO_SUBACCOUNT_SID", "TWILIO_ACCOUNT_SID"},
				},
			},
			"edge": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TWILIO_EDGE"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TWILIO_REGION"},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@nellisauction/pulumi-twilio",
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{RespectSchemaVersion: true}
			i.PyProject.Enabled = true
			return i
		})(),
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/nellisauction/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences:    map[string]string{"Pulumi": "3.*"},
			Namespaces:           map[string]string{mainPkg: "Twilio"},
		},
		MetadataInfo:                   tfbridge.NewProviderMetadata(metadata),
		EnableZeroDefaultSchemaVersion: true,
		EnableAccurateBridgePreview:    true,
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("twilio_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
