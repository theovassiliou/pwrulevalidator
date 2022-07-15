# pwrulevalidator
[![Go Report Card](https://goreportcard.com/badge/github.com/theovassiliou/pwrulevalidator)](https://goreportcard.com/report/github.com/theovassiliou/pwrulevalidator)

The Password Rule Validator a micro-services application that checks a given password against a set of configurable rules. 

## Requirements

### Architecture

- API specification in OpenAPI / Swagger
- rules set aka validator is a ressource (from OpenAPI perspective)


### Features

- configuratble via a set of rules, see [[#Parameterisable Rules]]
- persistant reference of rules
- only tranfer of pw candidate withouth user id
- support of historie of passwords without clear text transmission
- support of bcrypt hashes to avoid clear text transmission

### Operation

- docker environment
- support of logging
- Makefile



## Parameterisable Rules


    simple minimal password length      x chars for users *without* a priviliged rights 
    priviliged minimal password length  y chars for users *with* a priviliged rights
    password complexity                 passwort must fulfill x out y features:
                                        upper case, lower case, special char and number
    password age	                    at most y day(s)
                                        at least x day(s)
    Passworthistorie                    password history of y


## API Specification

* Create a new validator in response to a valid POST request at /validator
* Fetch  list of all validators in response to a valid GET request at /validators
* Delete a validator in response to a valid DELETE request at /validator/{id}
* Validate a password in response to a valid POST request at /validator/{id}
* Updata a validator in response to a valid PUT request at /validator/{id}
Validate a password 

## Description

In our simple application we will have a validator which has 

* `id` - a unique identfifier for the validator
* `ruleset` - a struct containing all parameters for the validator
* `description` - a description of the validator
