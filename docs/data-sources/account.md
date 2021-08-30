---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solana_account Data Source - terraform-provider-solana"
subcategory: ""
description: |-
  Provides all information associated with the account of the provided public key
---

# solana_account (Data Source)

Provides all information associated with the account of the provided public key

## Example Usage

```terraform
data "solana_account" "acc" {
  public_key = "11111111111111111111111111111111"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **public_key** (String) Base-58 encoded public key of the account to query

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **executable** (Boolean) Indicates whether the account contains a program
- **lamports** (Number) Number of lamports assigned to the account
- **owner** (String) Base-58 encoded public key of the program the account is assigned to
- **rent_epoch** (Number) The epoch at which this account will next owe rent

