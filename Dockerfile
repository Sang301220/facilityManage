#Specify the base image for the go app
FROM golang:1.18-alpine

WORKDIR /facilityManage

#Copy everything from this project into the filesystem of the container
COPY . /facilityManage
#Obtain the package needed to run main command. Alternatively use GO modules. Use && \ to exten command
RUN go build -o main 

# Start the application
CMD ["./main"]