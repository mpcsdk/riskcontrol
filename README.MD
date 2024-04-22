# riskControl

## make image

```
make image
```

## init db

<!-- ```
psql   -U postgres -h localhost -p 5432 -c "create database riskcontrol"
``` -->

```
mysql -u ssh -p123456 -h 127.0.0.1 -P 3305 -D riskcontrol < mpcCommon/mpcdao/migratioin/riskctrl/riskcontrol_dump.sql
```

```
docker run -it -v /path/config.yaml:/server/config.yaml -p 8000:8000 -p 8001:8001 riskcontrol
```
