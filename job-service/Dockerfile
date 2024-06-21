FROM golang:1-alpine3.19 AS build-stage
WORKDIR /job_service
COPY ./ /job_service
RUN mkdir -p /job_service/build
RUN go mod download
RUN go build -v -o /job_service/build/api ./cmd


FROM scratch
COPY --from=build-stage /job_service/build/api /
COPY --from=build-stage /job_service/.env /
EXPOSE 50051
CMD ["/api"]

