FROM golang:1.18-alpine as build

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -o pricefetcher .
# EXPOSE 3000
# ENTRYPOINT ["./pricefetcher"]

# distroless
FROM gcr.io/distroless/static-debian11
COPY --from=build /app/pricefetcher .
EXPOSE 3000
CMD ["/pricefetcher"]