FROM golang:1-alpine3.19 AS build-stage
WORKDIR /hireojob
COPY ./ /hireojob
RUN mkdir -p /hireojob/build
RUN go mod download
RUN go build -v -o /hireojob/build/api ./cmd


FROM scratch
COPY --from=build-stage /hireojob/build/api /
COPY --from=build-stage /hireojob/template/ /template/
COPY --from=build-stage /hireojob/static /static/
COPY --from=build-stage /hireojob/Logging /Logging/
COPY --from=build-stage /hireojob/.env /
EXPOSE 8000
CMD ["/api"]

