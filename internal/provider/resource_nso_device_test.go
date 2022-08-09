
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNsoDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsoDeviceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nso_device.test", "name", "test-device01"),
					resource.TestCheckResourceAttr("nso_device.test", "address", "10.1.1.1"),
					resource.TestCheckResourceAttr("nso_device.test", "port", "22"),
					resource.TestCheckResourceAttr("nso_device.test", "authgroup", "default"),
					resource.TestCheckResourceAttr("nso_device.test", "admin_state", "locked"),
					resource.TestCheckResourceAttr("nso_device.test", "cli_ned_id", "cisco-ios-cli-3.0:cisco-ios-cli-3.0"),
				),
			},
			{
				ResourceName:  "nso_device.test",
				ImportState:   true,
				ImportStateId: "tailf-ncs:devices/device=test-device01",
			},
		},
	})
}

func testAccNsoDeviceConfig_minimum() string {
	return `
	resource "nso_device" "test" {
		name = "test-device01"
	}
	`
}

func testAccNsoDeviceConfig_all() string {
	return `
	resource "nso_device" "test" {
		name = "test-device01"
		address = "10.1.1.1"
		port = 22
		authgroup = "default"
		admin_state = "locked"
		cli_ned_id = "cisco-ios-cli-3.0:cisco-ios-cli-3.0"
	}
	`
}