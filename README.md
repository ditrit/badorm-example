# Example of BaDorm as a stand-alone application

This is a example of how to use BaDorm as a stand-alone application. If you are interested in using BaDorm in a BaDaaS application please check the [BaDaaS example](https://github.com/ditrit/badaas-example)

## Run it

First, we need a database to store the data, in this case we will use CockroachDB:

```bash
docker compose up -d
```

After that, we can run the application:

```bash
go run .
```

And you should see something like:

```bash
2023/05/16 09:52:03 Setting up CRUD example
2023/05/16 09:52:03 Finished creating CRUD example
2023/05/16 09:52:03 Products with int = 1 are:
&{UUIDModel:{ID:1483487f-c585-4455-8d5b-2a58be27acbc CreatedAt:2023-05-16 09:50:12.025843 +0200 CEST UpdatedAt:2023-05-16 09:50:12.025843 +0200 CEST DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} String: Int:1 Float:0 Bool:false}
```

## Explore it

In [main.go](main.go) you will find the configuration required to use the BaDorm: provide a GormDB connection, start the BaDorm module and create the CRUD services.

In [example.go](example.go) you will find the actual example, where objects are created and then queried using BaDorm.

For more details, visit [BaDorm docs](https://github.com/ditrit/badaas/badorm/README.md).
