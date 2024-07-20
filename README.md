# Proof of Concept (POC) - Go Worker and Debezium

### Overview
Example of implementing a Go Worker (Consumer) to process data from Debezium. In this example, changes are captured from a MySQL database using Debezium, then sent to the Go Worker using RabbitMQ.

### Prequisites (Better Install with Docker Compose (check docker-compose.yml))
- Go 1.22.4 (See [Golang Installation](https://golang.org/doc/install))
- Debezium 2.7
- RabbitMQ
- MySQL 8 / Vitess 20
- Docker Compose

### Installation
1. Clone this repo
```sh
git clone 
```
2. Copy `example.env` to `.env`
```sh
cp env.sample .env
```
3. Adjust the `.env` (Mysql Configuration)
4. Build Worker Docker Image
```sh
docker build -t go-worker-dbz-worker:latest -f ./worker/Dockerfile .
```
5. Run Worker and Debezium Docker Container

If you use Mysql
```sh
docker-compose --env-file env.mysql.sample -f docker-compose.mysql.yml up -d
```
If you use Vitess
```sh
docker-compose --env-file env.vitess.sample -f docker-compose.vitess.yml up -d
```
5. Migrate Database! Create database with name "go_debezium". Then, you can manually create tables using the SQL code in the /database_migration folder.

#### See Log Stream (if not using Docker Desktop/Portainer)
```sh
docker-compose logs -f worker
```

### Sampe data from Debezium
This sample is a change data of the mapping_accounts table.
```json
{
  "schema": {
    "type": "struct",
    "fields": [
      {
        "type": "struct",
        "fields": [
          {
            "type": "int32",
            "optional": false,
            "field": "id"
          },
          {
            "type": "string",
            "optional": true,
            "field": "title"
          },
          {
            "type": "string",
            "optional": true,
            "field": "description"
          },
          {
            "type": "int32",
            "optional": true,
            "name": "io.debezium.time.Date",
            "version": 1,
2024/07/15 13:27:39 Received a message:
            "field": "doing_at"
          },
          {
            "type": "string",
            "optional": true,
            "name": "io.debezium.time.ZonedTimestamp",
            "version": 1,
            "field": "created_at"
          },
          {
            "type": "string",
            "optional": false,
            "name": "io.debezium.time.ZonedTimestamp",
            "version": 1,
            "default": "1970-01-01T00:00:00Z",
            "field": "updated_at"
          }
        ],
        "optional": true,
        "name": "trx.go_debezium.todo_lists.Value",
        "field": "before"
      },
      {
        "type": "struct",
        "fields": [
          {
            "type": "int32",
            "optional": false,
            "field": "id"
          },
          {
            "type": "string",
            "optional": true,
            "field": "title"
          },
          {
            "type": "string",
            "optional": true,
            "field": "description"
          },
          {
            "type": "int32",
            "optional": true,
            "name": "io.debezium.time.Date",
            "version": 1,
            "field": "doing_at"
          },
          {
            "type": "string",
            "optional": true,
            "name": "io.debezium.time.ZonedTimestamp",
            "version": 1,
            "field": "created_at"
          },
          {
            "type": "string",
            "optional": false,
            "name": "io.debezium.time.ZonedTimestamp",
            "version": 1,
            "default": "1970-01-01T00:00:00Z",
            "field": "updated_at"
          }
        ],
        "optional": true,
        "name": "trx.go_debezium.todo_lists.Value",
        "field": "after"
      },
      {
        "type": "struct",
        "fields": [
          {
            "type": "string",
            "optional": false,
            "field": "version"
          },
          {
            "type": "string",
            "optional": false,
            "field": "connector"
          },
          {
            "type": "string",
            "optional": false,
            "field": "name"
          },
          {
            "type": "int64",
            "optional": false,
            "field": "ts_ms"
          },
          {
            "type": "string",
            "optional": true,
            "name": "io.debezium.data.Enum",
            "version": 1,
            "parameters": {
              "allowed": "true,last,false,incremental"
            },
            "default": "false",
            "field": "snapshot"
          },
          {
            "type": "string",
            "optional": false,
            "field": "db"
          },
          {
            "type": "string",
            "optional": true,
            "field": "sequence"
          },
          {
            "type": "int64",
            "optional": true,
            "field": "ts_us"
          },
          {
            "type": "int64",
            "optional": true,
            "field": "ts_ns"
          },
          {
            "type": "string",
            "optional": true,
            "field": "table"
          },
          {
            "type": "int64",
            "optional": false,
            "field": "server_id"
          },
          {
            "type": "string",
            "optional": true,
            "field": "gtid"
          },
          {
            "type": "string",
            "optional": false,
            "field": "file"
          },
          {
            "type": "int64",
            "optional": false,
            "field": "pos"
          },
          {
            "type": "int32",
            "optional": false,
            "field": "row"
          },
          {
            "type": "int64",
            "optional": true,
            "field": "thread"
          },
          {
            "type": "string",
            "optional": true,
            "field": "query"
          }
        ],
        "optional": false,
        "name": "io.debezium.connector.mysql.Source",
        "field": "source"
      },
      {
        "type": "struct",
        "fields": [
          {
            "type": "string",
            "optional": false,
            "field": "id"
          },
          {
            "type": "int64",
            "optional": false,
            "field": "total_order"
          },
          {
            "type": "int64",
            "optional": false,
            "field": "data_collection_order"
          }
        ],
        "optional": true,
        "name": "event.block",
        "version": 1,
        "field": "transaction"
      },
      {
        "type": "string",
        "optional": false,
        "field": "op"
      },
      {
        "type": "int64",
        "optional": true,
        "field": "ts_ms"
      },
      {
        "type": "int64",
        "optional": true,
        "field": "ts_us"
      },
      {
        "type": "int64",
        "optional": true,
        "field": "ts_ns"
      }
    ],
    "optional": false,
    "name": "trx.go_debezium.todo_lists.Envelope",
    "version": 2
  },
  "payload": {
    "before": {
      "id": 1,
      "title": "Write Code",
      "description": "This is desc xxx",
      "doing_at": 19188,
      "created_at": "2024-07-15T13:19:33Z",
      "updated_at": "2024-07-15T06:25:49Z"
    },
    "after": {
      "id": 1,
      "title": "Write Code",
      "description": "This is desc updated",
      "doing_at": 19188,
      "created_at": "2024-07-15T13:19:33Z",
      "updated_at": "2024-07-15T06:27:38Z"
    },
    "source": {
      "version": "2.7.0.Final",
      "connector": "mysql",
      "name": "trx",
      "ts_ms": 1721024858000,
      "snapshot": "false",
      "db": "go_debezium",
      "sequence": null,
      "ts_us": 1721024858000000,
      "ts_ns": 1721024858000000000,
      "table": "todo_lists",
      "server_id": 1,
      "gtid": null,
      "file": "binlog.000002",
      "pos": 2867,
      "row": 0,
      "thread": 8,
      "query": null
    },
    "transaction": null,
    "op": "u",
    "ts_ms": 1721024858944,
    "ts_us": 1721024858944350,
    "ts_ns": 1721024858944350000
  }
}
```
