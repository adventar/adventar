# resource "aws_instance" "app" {
#   ami           = "ami-03d2ed888c6af3175" # ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-20190918
#   subnet_id     = aws_subnet.public_1a.id
#   instance_type = "t3.micro"

#   associate_public_ip_address = true

#   vpc_security_group_ids = [
#     aws_security_group.default.id,
#     aws_security_group.http.id,
#     aws_security_group.ssh.id,
#   ]

#   tags = {
#     Name = "adventar-001"
#   }
# }
