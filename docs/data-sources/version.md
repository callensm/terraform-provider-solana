---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solana_version Data Source - terraform-provider-solana"
subcategory: ""
description: |-
  (JSON RPC) https://docs.solana.com/developing/clients/jsonrpc-api#getversion Provides the current Solana software version running on the network node.
---

# solana_version (Data Source)

[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getversion) Provides the current Solana software version running on the network node.

## Example Usage

```terraform
data "solana_version" "node" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **feature_set** (Number) A unique identifier of the current software's feature set.
- **solana_core** (String) Software version of `solana-core` running on the node.


