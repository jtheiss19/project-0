FROM ubuntu
WORKDIR /app
COPY . /app
EXPOSE 8080
CMD [ "./main", "Host", "8080" ]

