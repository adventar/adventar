## ECS settings

```
# Configure
$ ecs-cli configure --cluster adventar --region ap-northeast-1 --config-name adventar --default-launch-type FARGATE

# Create or update task def
$ ecs-cli compose create --cluster-config adventar --create-log-groups

# Run task
$ ecs-cli compose run --cluster-config adventar app /go/update_entry_site_info
```
