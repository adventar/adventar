resource "aws_cloudwatch_event_rule" "update_entry_site_info" {
  name = "adventar-batch-update_entry_site_info"

  schedule_expression = "cron(0 * 1-25 12 ? *)"
  is_enabled          = true
}

resource "aws_cloudwatch_event_rule" "update_entry_site_info_2" {
  name = "adventar-batch-update_entry_site_info_2"

  schedule_expression = "cron(0 15-23 30 11 ? *)"
  is_enabled          = true
}

resource "aws_cloudwatch_event_target" "ecs_update_entry_site_info" {
  rule     = aws_cloudwatch_event_rule.update_entry_site_info.name
  arn      = "arn:aws:ecs:ap-northeast-1:287379415997:cluster/adventar"
  role_arn = "arn:aws:iam::287379415997:role/service-role/AWS_Events_Invoke_ECS_1528030466"

  ecs_target {
    launch_type         = "FARGATE"
    task_definition_arn = "arn:aws:ecs:ap-northeast-1:287379415997:task-definition/update_entry_site_info"
    task_count          = 1

    network_configuration {
      assign_public_ip = true
      security_groups = [
        "sg-073b60f32ebeb3cef",
      ]
      subnets = [
        "subnet-059a83efe99ee37ee",
      ]
    }
  }
}
