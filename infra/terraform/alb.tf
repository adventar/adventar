resource "aws_alb" "api" {
  name            = "adventar-api"
  security_groups = [aws_security_group.default.id, aws_security_group.alb.id]
  subnets         = [aws_subnet.public_1a.id, aws_subnet.public_1c.id]

  internal                   = false
  enable_deletion_protection = true

  access_logs {
    bucket  = aws_s3_bucket.adventar.bucket
    prefix  = "alb-logs"
    enabled = true
  }
}

resource "aws_alb_target_group" "api" {
  name        = "adventar-api"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_vpc.adventar.id

  deregistration_delay = 30

  health_check {
    interval            = 30
    path                = "/health_check"
    port                = 80
    protocol            = "HTTP"
    timeout             = 5
    healthy_threshold   = 5
    unhealthy_threshold = 2
    matcher             = 200
  }
}

resource "aws_alb_listener" "api_https" {
  load_balancer_arn = aws_alb.api.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = aws_acm_certificate.api_adventar_org.arn

  default_action {
    target_group_arn = aws_alb_target_group.api.arn
    type             = "forward"
  }
}

resource "aws_alb_listener" "api_http" {
  load_balancer_arn = aws_alb.api.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.api.arn
    type             = "forward"
  }
}
