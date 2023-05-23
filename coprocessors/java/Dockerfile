FROM openjdk:20-jdk-slim AS build-env
WORKDIR /app
COPY . .
RUN ./mvnw package

FROM openjdk:20-slim
WORKDIR /app
COPY --from=build-env /app/target .
ENTRYPOINT java -jar ./java-coprocessor-1.0.0.jar
