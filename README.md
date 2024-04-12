## Kartaca DevOps Task

When you execute `docker-compose up`, it orchestrates several tasks:

Firstly, it launches three instances of `elasticsearch:8.13.0`, configuring them as a cluster with three nodes. As each node becomes operational, data seeding commences. This involves triggering a script housed within the `curlimages/curl:8.7.1` container, located at [`/scripts/init.sh`](https://github.com/akinbezatoglu/kartaca-devops-muhendisi-task/blob/main/scripts/init.sh). This script sends HTTP requests to the Elasticsearch nodes, thereby creating two indexes, namely `iller` and `ulkeler`, and populating them with ten documents each.

Subsequently, the process moves to building and deploying both Go and Python applications. These apps establish connections with the Elasticsearch cluster using the respective Elasticsearch clients for [Go](https://github.com/elastic/go-elasticsearch) and [Python](https://github.com/elastic/elasticsearch-py).

Meanwhile, the monitoring infrastructure is set up. This involves launching `prom/prometheus:v2.51.1` to gather metrics, while `prom/node-exporter:v1.7.0` collects metrics pertaining to the host, and `cadvisor:v0.49.1` handles container-specific metrics.

For visualization, `grafana/grafana:10.1.9` is spun up to provide an interface for visualizing the data collected by Prometheus.

Lastly, for load balancing and proxying HTTP traffic to the application, `haproxy:3.0-dev6-alpine3.19` is deployed. This ensures efficient distribution of requests across the deployed application instances.

### Grafana Dashboards as JSON

When downloading dashboards by ID as JSON from `https://grafana.com/grafana/dashboards/`, you might encounter an issue with the cAdvisor dashboard, displaying a message like `Datasource ${DS_PROMETHEUS} was not found` in Grafana. To resolve this problem, I took the following steps:

Firstly, I manually imported the Node Exporter and cAdvisor Exporter Dashboards by their respective IDs directly in the Grafana front-end.

Then, to ensure error-free usage in Grafana configuration, I exported these dashboards as JSON files by clicking "save as" in the share section of Grafana.

Finally, I added these JSON files to the project directory under [`/grafana/provisioning/dashboards/json`](https://github.com/akinbezatoglu/kartaca-devops-muhendisi-task/tree/main/grafana/provisioning/dashboards/json). This ensures that the dashboards are readily available and properly configured within the project setup.

---

#### Python App
```sh
$ curl localhost:4444
"Merhaba, Python!"
```

```sh
$ curl localhost:4444/staj
{
  "il": "istanbul",
  "ilceler": [
    "avcilar",
    "besiktas",
    "esenler"
  ],
  "nufus": 15655924
}
```
#### Golang App
```sh
$ curl localhost:5555
"Merhaba, Go!"
```

```sh
$ curl localhost:5555/staj
{
  "ulke": "turkiye",
  "nufus": 85372377,
  "baskent": "ankara"
}
```

#### HAProxy 
```sh
$ curl kartaca.localhost/pythonapp
{
  "il": "antalya",
  "ilceler": [
    "alanya",
    "kepez",
    "konyaalti"
  ],
  "nufus": 321457
}
```

```sh
$ curl kartaca.localhost/goapp
{
  "ulke": "birlesik_arap_emirlikleri",
  "nufus": 9282410,
  "baskent": "abu_dabi"
}
```

```sh
$ curl http://kartaca.localhost/grafana/api/datasources
[
  {
    "id": 1,
    "uid": "PBFA97CFB590B2093",
    "orgId": 1,
    "name": "Prometheus",
    "type": "prometheus",
    "typeName": "Prometheus",
    "typeLogoUrl": "public/app/plugins/datasource/prometheus/img/prometheus_logo.svg",
    "access": "proxy",
    "url": "http://prometheus:9090",
    "user": "",
    "database": "",
    "basicAuth": false,
    "isDefault": true,
    "jsonData": {

    },
    "readOnly": false
  }
]
```

[Node Exporter Full](https://github.com/akinbezatoglu/kartaca-devops-muhendisi-task/blob/90a06884f56b6db8ac949dd20a1ec95e62fda18d/grafana/provisioning/dashboards/json/Node%20Exporter%20Full-1712780772323.json#L23822)
```sh
$ curl http://kartaca.localhost/grafana/api/dashboards/uid/rYdddlPWk
{
  "meta": {
    "type": "db",
    "canSave": true,
    "canEdit": true,
    "canAdmin": true,
    "canStar": true,
    "canDelete": true,
    "slug": "node-exporter-full",
    "url": "/grafana/d/rYdddlPWk/node-exporter-full",
    "expires": "0001-01-01T00:00:00Z",
    "created": "2024-04-12T21:07:14Z",
    "updated": "2024-04-12T21:07:14Z",
    "updatedBy": "Anonymous",
    "createdBy": "Anonymous",
    "version": 1,
    "hasAcl": false,
    "isFolder": false,
    "folderId": 0,
    "folderUid": "",
    "folderTitle": "General",
    "folderUrl": "",
    "provisioned": true,
    "provisionedExternalId": "Node Exporter Full-1712780772323.json",
    ...
}
```

[Cadvisor exporter](https://github.com/akinbezatoglu/kartaca-devops-muhendisi-task/blob/90a06884f56b6db8ac949dd20a1ec95e62fda18d/grafana/provisioning/dashboards/json/Cadvisor%20exporter-1712780787051.json#L806)
```sh
$ curl http://kartaca.localhost/grafana/api/dashboards/uid/pMEd7m0Mz
{
  "meta": {
    "type": "db",
    "canSave": true,
    "canEdit": true,
    "canAdmin": true,
    "canStar": true,
    "canDelete": true,
    "slug": "cadvisor-exporter",
    "url": "/grafana/d/pMEd7m0Mz/cadvisor-exporter",
    "expires": "0001-01-01T00:00:00Z",
    "created": "2024-04-12T21:07:14Z",
    "updated": "2024-04-12T21:07:14Z",
    "updatedBy": "Anonymous",
    "createdBy": "Anonymous",
    "version": 1,
    "hasAcl": false,
    "isFolder": false,
    "folderId": 0,
    "folderUid": "",
    "folderTitle": "General",
    "folderUrl": "",
    "provisioned": true,
    "provisionedExternalId": "Cadvisor exporter-1712780787051.json",
    ...
}
```