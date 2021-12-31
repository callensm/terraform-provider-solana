# Random Keypair
resource "solana_keypair" "key" {
  output_path = "${path.module}/id.json"
  random      = true
}

# Grind Keypair
resource "solana_keypair" "key" {
  output_path = "${path.module}/id.json"

  grind {
    prefix = "abc"
  }
}
