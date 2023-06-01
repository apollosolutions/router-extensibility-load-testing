FROM rust:latest AS builder
WORKDIR /app
COPY . .
RUN cargo build --release


FROM debian:buster-slim
RUN apt-get update && apt-get install -y openssl
WORKDIR /app
COPY --from=builder /app/target/release/fixed .

COPY fixtures.yaml .
CMD ./fixed
