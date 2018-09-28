variable "gcp_project" {}
variable "gcp_credentials" {
    default = "../.google-credentials"
}
variable "gcp_region" {}
variable "gcp_zone" {}

provider "google" {
    version = "~> 1.18"
    credentials = "${file(var.gcp_credentials)}"
    project     = "${var.gcp_project}"
    region      = "${var.gcp_region}"
    zone        = "${var.gcp_zone}"
}

resource "google_compute_address" "bastion-address" {
    name = "bastion-address"
}

resource "google_compute_address" "webserver-address" {
    name = "webserver-address"
}