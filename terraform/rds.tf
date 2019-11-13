resource "aws_db_instance" "adventar-002" {
  identifier        = "adventar-db-002"
  allocated_storage = 20
  storage_type      = "gp2"
  engine            = "mysql"
  engine_version    = "8.0.16"
  instance_class    = "db.t3.micro"
  username          = "root"
  password          = "dummy000"

  backup_retention_period         = 7
  backup_window                   = "17:00-17:30"
  maintenance_window              = "Mon:18:00-Mon:18:30"
  enabled_cloudwatch_logs_exports = ["error", "slowquery"]
  db_subnet_group_name            = aws_db_subnet_group.adventar-public.id
  auto_minor_version_upgrade      = false
  publicly_accessible             = true
  skip_final_snapshot             = true
  deletion_protection             = true

  vpc_security_group_ids = [
    aws_security_group.default.id,
    # "sg-06577841aa97c4e5b"
  ]
}

resource "aws_db_subnet_group" "adventar-public" {
  name       = "adventar-public"
  subnet_ids = [aws_subnet.public_1a.id, aws_subnet.public_1c.id]
}
