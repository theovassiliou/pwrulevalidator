# pwrulemanager
[[golang]]

The Password Rule Validator a micro-services application that checks a given password against a set of configurable rules. 

## Requirements

### Architecture

- API specification in OpenAPI / Swagger
- rules set is a ressource (from OpenAPI perspective)


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
