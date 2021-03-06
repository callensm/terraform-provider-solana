---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solana_epoch Data Source - terraform-provider-solana"
subcategory: ""
description: |-
  (JSON RPC) https://docs.solana.com/developing/clients/jsonrpc-api#getepochinfo Provides all of the relevant information about the current epoch.
---

# solana_epoch (Data Source)

[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getepochinfo) Provides all of the relevant information about the current epoch.

## Example Usage

```terraform
data "solana_epoch" "e" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **absolute_slot** (Number) The current absolute slot in the epoch.
- **block_height** (Number) The current block height.
- **epoch** (Number) The current epoch count.
- **slot_index** (Number) The current slot relative to the start of the current epoch.
- **slots_in_epoch** (Number) The number of slots in this epoch.


