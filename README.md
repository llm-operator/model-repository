# Model Manager

Model Manager consists of the following components;

- `server`: Serve gRPC requests and HTTP requests
- `loadedr`: Load open models to the system

`loader` currently loads open models from Hugging Face, but we can extend that to support other locations.

## Running `server` Locally

```bash
make build-server
./bin/server run --config config.yaml
```

`config.yaml` has the following content:

```yaml
httpPort: 8080
grpcPort: 8081
internalGrpcPort: 8082

objectStore:
  s3:
    pathPrefix: models

debug:
  standalone: true
  sqlitePath: /tmp/model_manager.db
```

You can then connect to the DB.

```bash
sqlite3 /tmp/model_manager.db
# Run the query inside the database.
insert into models
  (model_id, tenant_id, created_at, updated_at)
values
  ('my-model', 'fake-tenant-id', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
```

You can then hit the endpoint.

```bash
curl http://localhost:8080/v1/models

grpcurl -d '{"base_model": "base", "suffix": "suffix", "tenant_id": "fake-tenant-id"}' -plaintext localhost:8082 list llmoperator.models.server.v1.ModelsInternalService/CreateModel
```

## Running `loader` Locally

```bash
make build-loader
./bin/loader run --config config.yaml
```

`config.yaml` has the following content:

```yaml
objectStore:
  s3:
    pathPrefix: models
    baseModelPathPrefix: base-models

baseModels:
- google/gemma-2b

modelLoadInterval: 1m

debug:
  standalone: true
  sqlitePath: /tmp/model_manager.db
```
