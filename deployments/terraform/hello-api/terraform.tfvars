project_id    = "samyn-project4"
cluster_name  = "samyn-cluster"
region        = "europe-west1"
zone          = "europe-west1-d"
gke_num_nodes = 1
machine_type  = "e2-medium"
disk_size     = 25
namespace     = "lazone"
environment   = "staging"
app_name      = "hello"
kms_key_ring_hello = "kr-hello"
kms_key_hello      = "key-default-hello"
domain             = "dev2choiz.com"
certPrivateKeyPath = "./../../../local/files/tls/server.key"
certCrtPath        = "./../../../local/files/tls/server.crt"
