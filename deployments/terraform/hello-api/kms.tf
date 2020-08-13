
resource "google_kms_key_ring" "hello-keyring" {
  name     = var.kms_key_ring_hello
  location = var.region
}

resource "google_kms_crypto_key" "hello-key" {
  name            = var.kms_key_hello
  key_ring        = google_kms_key_ring.hello-keyring.id
  purpose         = "ENCRYPT_DECRYPT"

  lifecycle {
    prevent_destroy = true
  }
}