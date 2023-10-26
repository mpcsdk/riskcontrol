# riskControl
## make image

```
make image
```

## init db

```
psql   -U postgres -h localhost -p 5432 -c "create database riskcontrol"
```
```
psql   -U postgres -h localhost -p 5432 -c "create database tfa"
```

```
psql   -U postgres -h localhost -p 5432 -d tfa -f ./manifest/migration/tfa_dump.sql
```
```
psql   -U postgres -h localhost -p 5432 -d riskcontrol -f ./manifest/migration/riskcontrol_dump.sql
```


```
docker run -it -v /path/config.yaml:/server/config.yaml -p 8000:8000 -p 8001:8001 riskcontrol
```
