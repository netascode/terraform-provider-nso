---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nso_device_config Data Source - terraform-provider-nso"
subcategory: "Device"
description: |-
  Retrieves a config part of an NSO device.
---

# nso_device_config (Data Source)

Retrieves a config part of an NSO device.

## Example Usage

```terraform
data "nso_device_config" "example" {
  device = "c1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device` (String) An NSO device name.

### Optional

- `instance` (String) An instance name from the provider configuration.
- `path` (String) A RESTCONF/YANG config path, e.g. `tailf-ned-cisco-ios:access-list/access-list=1`.

### Read-Only

- `attributes` (Map of String) Map of key-value pairs which represents the attributes and its values.
- `id` (String) The RESTCONF path of the retrieved configuration.

