resource "aws_security_group" "default" {
  vpc_id      = aws_vpc.adventar.id
  description = "default VPC security group"
  name        = "default"

  ingress {
    protocol  = -1
    self      = true
    from_port = 0
    to_port   = 0
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "adventar-default"
  }
}

resource "aws_security_group" "alb" {
  vpc_id      = aws_vpc.adventar.id
  name        = "adventar-alb"
  description = "Allow http and https"

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "adventar-alb"
  }
}
