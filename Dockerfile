FROM gcr.io/distroless/base

COPY server /bin/server

EXPOSE 8080

CMD ["server"]