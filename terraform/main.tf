variable "ssh_public_key" {}

variable "gcp_project" {}
variable "gcp_credentials" {
    default = ".google-credentials"
}
variable "gcp_region" {}
variable "gcp_zone" {}

variable "base_image" {
    default = "projects/centos-cloud/global/images/centos-7-v20180911"
}
variable "web_server_tag" {
    default = "web-server"
}
variable "bastion_tag" {
    default = "bastion"
}

provider "google" {
    version = "~> 1.18"
    credentials = "${file(var.gcp_credentials)}"
    project     = "${var.gcp_project}"
    region      = "${var.gcp_region}"
    zone        = "${var.gcp_zone}"
}

data "google_compute_address" "bastion-address" {
  name = "bastion-address"
}

data "google_compute_address" "webserver-address" {
  name = "webserver-address"
}

resource "google_compute_firewall" "allow-access-to-bastion" {
    name = "allow-access-to-bastion"
    network = "default"
    direction = "INGRESS"
    source_ranges = ["0.0.0.0/0"]
    allow {
        protocol = "tcp"
        ports = ["22"]
    }
    target_tags = ["${var.bastion_tag}"]
}

resource "google_compute_instance" "bastion" {
    name = "bastion-host"
    machine_type = "f1-micro"

    boot_disk {
        initialize_params {
            image = "${var.base_image}"
        }
    }

    network_interface {
        network = "default"

        access_config {
            nat_ip = "${data.google_compute_address.bastion-address.address}"
        }
    }
    tags = ["${var.bastion_tag}"]

    metadata {
        ssh-keys = "${var.ssh_public_key}"
    }
}

resource "google_compute_firewall" "allow-access-to-webserver" {
    name = "allow-access-to-webserver"
    network = "default"
    direction = "INGRESS"
    source_ranges = ["0.0.0.0/0"]
    allow {
        protocol = "tcp"
        ports = ["80"]
    }
    target_tags = ["${var.web_server_tag}"]
}

resource "google_compute_instance" "web-server" {
    name = "web-server"
    machine_type = "f1-micro"

    boot_disk {
        initialize_params {
            image = "${var.base_image}"
        }
    }

    network_interface {
        network = "default"

        access_config {
            nat_ip = "${data.google_compute_address.webserver-address.address}"
        }
    }
    tags = ["${var.web_server_tag}"]

    metadata {
        ssh-keys = "${var.ssh_public_key}"
    }
}

resource "google_compute_instance_group" "webservers" {
    name = "webservers"
    instances = [
        "${google_compute_instance.web-server.self_link}"
    ]

    named_port {
        name = "http"
        port = "80"
    }
}

resource "google_compute_health_check" "webservers" {
    name = "webservers-healthcheck"

    http_health_check {}
}

resource "google_compute_backend_service" "webservers" {
    name = "webservers-backend"

    backend {
        group = "${google_compute_instance_group.webservers.self_link}"
    }

    health_checks = [
        "${google_compute_health_check.webservers.self_link}"
    ]
}

resource "google_compute_url_map" "webservers" {
    name = "webservers-url-map"

    default_service = "${google_compute_backend_service.webservers.self_link}"
}

resource "google_compute_target_http_proxy" "webservers" {
    name = "webservers-target-proxy"
    url_map = "${google_compute_url_map.webservers.self_link}"
}


resource "google_compute_global_forwarding_rule" "webservers" {
    name = "webservers-global-forwarding-rule"
    port_range = "80"
    target = "${google_compute_target_http_proxy.webservers.self_link}"
}

