---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solana_token_registry_entry Data Source - terraform-provider-solana"
subcategory: ""
description: |-
  Provides the metadata for a Solana token registry entry based on the given minting address.
---

# solana_token_registry_entry (Data Source)

Provides the metadata for a Solana token registry entry based on the given minting address.

## Example Usage

```terraform
# Basic Usage
data "solana_token_registry_entry" "entry" {
  mint_address = "abc123"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **mint_address** (String) The public key of the mint account address as a base-58 string.

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **authority** (String) The public key of the registration authority as a base-58 string.
- **initialized** (Boolean) Flag that indicates in the token entry has been initialized.
- **name** (String) The name of the token.

