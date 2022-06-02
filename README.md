# To build application
go build -o target/email-sender ./cmd

# To run application
./target/email-sender </path/to/email_template.json> </path/to/customers.csv> </path/to/output_emails.json> </path/to/errors.csv>

if there is no arguments passed in, it will take default values as:
+ resource/email_template.json
+ resource/customers.csv
+ resource/output_emails.json
+ resource/errors.csv

# To run unit test
go test -v ./cmd

# To build docker image
docker build -t <image-name> .

# To run image
docker run -d -it --name <container-name> <image-name>

# Exec to container
docker exec -it <container-name> bash

and run application as above step