#Step #1
FROM maven:3-openjdk-11 as build-stage
WORKDIR /app
COPY . /app
#RUN mvn protobuf:compile -DprotocExecutable=/usr/bin/protoc
RUN mvn package
#Step #2
FROM openjdk:7 as 
WORKDIR /app
COPY --from=build-stage /app/target/dict-1.0-SNAPSHOT.jar  /app
#RUN java -jar /app/dict-1.0-SNAPSHOT.jar
#CMD ["java", "Main"]
CMD ["java", "-jar" ,"/app/dict-1.0-SNAPSHOT.jar" ]