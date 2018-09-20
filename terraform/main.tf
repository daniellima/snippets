provider "google" {
    version = "~> 1.18"
    credentials = "${file("tonal-vector-202523-20a4532eb10b.json")}"
    project     = "tonal-vector-202523"
    region      = "us-central1"
    zone        = "us-central1-c"
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
    target_tags = ["bastion"]
}

resource "google_compute_instance" "bastion" {
    name = "bastion-host"
    machine_type = "f1-micro"

    boot_disk {
        initialize_params {
            image = "projects/centos-cloud/global/images/centos-7-v20180911"
        }
    }

    network_interface {
        network = "default"

        access_config {
            // Ephemeral IP
        }
    }
    tags = ["bastion"]

    metadata {
        ssh-keys = "danielsantos:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvnjF2HiayRwHj1PxPBcb60Z5iYZHhnPLypNstzY68bvsmZY7xwuEPmJ/wbMw/KfFNaUJipmjdHO2vcZFwRDzpRWhy0SulU+JeOiBKCO660FeSz2dIuiqrLCuFIODAukI+Uc80+L4sjlQ0/reeSzyT2bN9SCCzURc2U5uG/SMOoKtN+UvafgH3XO/zI0GuH/DQ4oNCY1LqOmF55Ukyd/DZKHPwsIjlLLbRZXynvTFF+EzycMnQKtCvMCODpK/TycSVghmNnmuTdKdo3WBDj0xij4YpxCFKXoch4wLT+h5Lx0LsB/+hmE+cgKQdX02PB1GQWppJPyGOskJBWvqiVoLv danielsantos@BRRIOWN013037"
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
    target_tags = ["web-server"]
}

resource "google_compute_instance" "web-server" {
    name = "web-server"
    machine_type = "f1-micro"

    boot_disk {
        initialize_params {
            image = "projects/centos-cloud/global/images/centos-7-v20180911"
        }
    }

    network_interface {
        network = "default"

        access_config {
            // Ephemeral IP
        }
    }
    tags = ["web-server"]

    metadata {
        ssh-keys = "danielsantos:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvnjF2HiayRwHj1PxPBcb60Z5iYZHhnPLypNstzY68bvsmZY7xwuEPmJ/wbMw/KfFNaUJipmjdHO2vcZFwRDzpRWhy0SulU+JeOiBKCO660FeSz2dIuiqrLCuFIODAukI+Uc80+L4sjlQ0/reeSzyT2bN9SCCzURc2U5uG/SMOoKtN+UvafgH3XO/zI0GuH/DQ4oNCY1LqOmF55Ukyd/DZKHPwsIjlLLbRZXynvTFF+EzycMnQKtCvMCODpK/TycSVghmNnmuTdKdo3WBDj0xij4YpxCFKXoch4wLT+h5Lx0LsB/+hmE+cgKQdX02PB1GQWppJPyGOskJBWvqiVoLv danielsantos@BRRIOWN013037"
    }
}