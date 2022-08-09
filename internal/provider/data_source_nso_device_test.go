
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNsoDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNsoDeviceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nso_device.test", "address", "10.1.1.1"),
					resource.TestCheckResourceAttr("data.nso_device.test", "port", "22"),
					resource.TestCheckResourceAttr("data.nso_device.test", "authgroup", "default"),
					resource.TestCheckResourceAttr("data.nso_device.test", "admin_state", "locked"),
					resource.TestCheckResourceAttr("data.nso_device.test", "cli_ned_id", "cisco-ios-cli-3.0:cisco-ios-cli-3.0"),
				),
			},
		},
	})
}

const testAccDataSourceNsoDeviceConfig = `

resource "nso_device" "test" {
  name = "test-device01"
  address = "10.1.1.1"
  port = 22
  authgroup = "default"
  admin_state = "locked"
  cli_ned_id = "cisco-ios-cli-3.0:cisco-ios-cli-3.0"
}

data "nso_device" "test" {
  name = "test-device01"
  depends_on = [nso_device.test]
}
`