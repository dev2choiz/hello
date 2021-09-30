FROM golang:1.17
COPY ./entrypoint.sh /entrypoint.sh
COPY --from=gcr.io/cloudsql-docker/gce-proxy:1.25.0 /cloud_sql_proxy /cloud_sql_proxy
