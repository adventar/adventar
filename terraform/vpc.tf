resource "aws_vpc" "adventar" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = "adventar"
  }
}

resource "aws_subnet" "public_1a" {
  vpc_id            = aws_vpc.adventar.id
  availability_zone = "ap-northeast-1a"
  cidr_block        = "10.0.1.0/24"

  tags = {
    Name = "adventar-public-1a"
  }
}

resource "aws_subnet" "public_1c" {
  vpc_id            = aws_vpc.adventar.id
  availability_zone = "ap-northeast-1c"
  cidr_block        = "10.0.2.0/24"

  tags = {
    Name = "adventar-public-1c"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.adventar.id

  tags = {
    Name = "adventar"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.adventar.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "adventar-public"
  }
}

resource "aws_route_table_association" "public_1a" {
  subnet_id      = aws_subnet.public_1a.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public_1c" {
  subnet_id      = aws_subnet.public_1c.id
  route_table_id = aws_route_table.public.id
}
