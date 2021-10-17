project_id    = "samyn-project5"
cluster_name  = "samyn-cluster"
region        = "us-central1"
zone          = "us-central1-a"
namespace     = "lazone"
domain        = "dev2choiz.com"
# standard
gke_num_nodes = 2
machine_type  = "e2-medium"
#machine_type  = "e2-micro"
disk_size     = 25
# anthos
#gke_num_nodes = 2
#machine_type  = "e2-standard-4"
#disk_size     = 50
membership_id = "samyn-cluster-mi"
dir_output    = "./../../../local/files"

# postgresql
sql_instance = "hello-pg-instance"
db_name      = "hello-db"
