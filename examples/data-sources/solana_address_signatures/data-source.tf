# Basic Usage
data "solana_address_signatures" "sigs" {
  address = "11111111111111111111111111111111"
}

# Using Custom Search Parameters
data "solana_address_signatures" "sigs" {
  address = "11111111111111111111111111111111"

  search_options {
    limit = 10
    after = "Vote1111111111111111111111111111"
  }
}
