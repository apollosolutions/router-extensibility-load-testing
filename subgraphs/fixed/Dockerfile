FROM rust:bookworm AS builder
WORKDIR /app

COPY . .
RUN --mount=type=cache,target=/usr/local/cargo/registry \
    --mount=type=cache,target=/app/target \
    cargo build --release \
    && cp target/release/fixed /app/fixed


FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y openssl
WORKDIR /app
COPY --from=builder /app/fixed .

COPY fixtures.yaml .
CMD ./fixed
