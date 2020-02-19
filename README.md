# Set up computation cluster on AWS

## Prerequirements

- [Terraform download](https://www.terraform.io/downloads.html)
- [Ansible installation](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#from-pip)

## Creating EC2 instances

Add default VPC id to terraform 
```
terraform import aws_default_vpc.default vpc-id
```

Export credentials as variables
```
export AWS_ACCESS_KEY_ID="accesskey"
export AWS_SECRET_ACCESS_KEY="secretkey"
export AWS_DEFAULT_REGION="eu-central-1"
```

Deploy EC2 instances
```
terraform init
terraform plan
terraform apply
```

# Update code

```
ansible all -m shell -a 'source ~/.bash_profile && git clone https://github.com/boodyvo/google-hashcode-2020.git $GOPATH/src/github.com/boodyvo/google-hashcode-2020'
```

```
ansible all -m shell -a 'source ~/.bash_profile && cd $GOPATH/src/github.com/boodyvo/google-hashcode-2020 && git pull'
ansible all -m shell -a 'source ~/.bash_profile && cd $GOPATH/src/github.com/boodyvo/google-hashcode-2020/code && go run main.go'
```

# Get result

```
ansible-playbook fetch-playbook.yml
```

