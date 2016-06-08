variable "cloudflare_email" {
  type = "string"
  description = "CloudFlare account email."
}

variable "cloudflare_token" {
  type = "string"
  description = "CloudFlare API global token. See: https://www.cloudflare.com/a/account/my-account"
}

variable "do_token" {
  type = "string"
  description = "Digital Ocean API token read/write. See: https://cloud.digitalocean.com/settings/api/tokens"
}

variable "domain" {
  type = "string"
  description = "The domain to add the record to."
}
variable "domain_name" {
  type = "string"
  description = "The name of the record."
}

variable "sshkey" {
  type = "string"
  description = "Public key file path."
  default = "certs/ssh-key.pub"
}
