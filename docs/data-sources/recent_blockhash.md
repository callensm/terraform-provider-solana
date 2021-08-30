---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solana_recent_blockhash Data Source - terraform-provider-solana"
subcategory: ""
description: |-
  Retrieves a recent block hash from the ledger and the associated cost in lamports per signature on a new transaction for that block
---

# solana_recent_blockhash (Data Source)

Retrieves a recent block hash from the ledger and the associated cost in lamports per signature on a new transaction for that block

## Example Usage

```terraform
data "solana_recent_blockhash" "hash" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **blockhash** (String) Base-58 encoded hash string of the block
- **lamports_per_signature** (Number) The lamports cost per signature of the block

