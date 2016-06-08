provider "cloudflare" {
    email = "${var.cloudflare_email}"
    token = "${var.cloudflare_token}"
}

provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_ssh_key" "default" {
  name = "redirect-server"
  public_key = "${file(var.sshkey)}"
}

resource "digitalocean_droplet" "redirect-server" {
  image = "coreos-beta"
  name = "redirect-server"
  region = "nyc3"
  size = "512mb"
  private_networking = true
  ssh_keys = ["${digitalocean_ssh_key.default.id}"]
  user_data = "${file("cloud-config/redirect-server.yml")}"
}

resource "digitalocean_floating_ip" "redirect-server" {
    droplet_id = "${digitalocean_droplet.redirect-server.id}"
    region = "nyc3"
}

resource "cloudflare_record" "redirect-server-ipv4" {
    domain = "${var.domain}"
    name = "${var.domain_name}"
    value = "${digitalocean_floating_ip.redirect-server.ip_address}"
    type = "A"
}
