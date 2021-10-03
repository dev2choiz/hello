## Postgres migration

### Generate a diff migration file

- modify the model in `./pkg/models/`
- run `make migration-diff`
The migration file is generated in `./migrations/`

### Play migrations in dev environment
```shell
make migration-up
```

### Play migrations in gcp
There is a cloud function handling migrations and subscribed to the `pg-migration` pubsub topic.  
to trigger a `migration up`, run the `pubsub` jenkins job to publish message using these parameters:  
`TOPIC_ID: pg-migration`  
`MESSAGE: up`  
`ATTRIBUTES: (empty)`  
