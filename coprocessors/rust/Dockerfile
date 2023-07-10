FROM rust:bookworm AS builder
WORKDIR /app

COPY . .
RUN --mount=type=cache,target=/usr/local/cargo/registry \
    --mount=type=cache,target=/app/target \
    cargo build --release \
    && cp target/release/coprocessor /app/coprocessor


FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/coprocessor .

CMD ./coprocessor
