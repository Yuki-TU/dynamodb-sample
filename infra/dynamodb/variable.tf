variable "env" {
  description = "The environment in which the Network will be created"
  type        = string
  default     = "stg"
}

variable "service" {
  description = "The service name for the resources"
  type        = string
  default     = "point-app"
}


variable "region" {
  description = "The region in which the VPC will be created"
  type        = string
  default     = "ap-northeast-1"
}

locals {
  # Fully Qualified Name
  fqn = "${var.service}-${var.env}"
}
