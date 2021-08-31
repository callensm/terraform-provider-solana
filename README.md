# Terraform Solana Provider

[![Release](https://img.shields.io/github/v/release/callensm/terraform-provider-solana)](https://github.com/callensm/terraform-provider-solana/releases)
[![Installs](https://img.shields.io/badge/dynamic/json?logo=terraform&label=installs&query=$.data.attributes.downloads&url=https%3A%2F%2Fregistry.terraform.io%2Fv2%2Fproviders%2F713)](https://registry.terraform.io/providers/callensm/solana)
[![Registry](https://img.shields.io/badge/registry-doc%40latest-lightgrey?logo=terraform)](https://registry.terraform.io/providers/callensm/solana/latest/docs)
[![License](https://img.shields.io/badge/license-MPLv2.0-blue.svg)](https://github.com/callensm/terraform-provider-solana/blob/main/LICENSE)

[Registry Page](https://registry.terraform.io/providers/callensm/solana/latest)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) 1.16.x (for building from source)

## Example Usage

Full provider documentation can be found on the Terraform registry at the link found above.

```hcl
terraform {
  required_providers {
    solana = {
      source  = "callensm/solana"
      version = "<LATEST_VERSION>"
    }
  }
}

provider "solana" {
  cluster = "testnet"
}

data "solana_address_signatures" "sigs" {
  address = "11111111111111111111111111111111"

  search_options {
    limit = 1
  }
}

data "solana_signature_status" "sig" {
  signature                  = data.solana_address_signatures.sigs.results.0.signature
  search_transaction_history = true
}
```

## License

[MPL v2.0](./LICENSE)
