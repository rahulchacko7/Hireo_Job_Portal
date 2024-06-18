FROM golang:1.22-alpine AS build-stage
WORKDIR /notification_service
COPY ./ /notification_service 
RUN mkdir -p /notification_service/build
RUN go mod download
RUN go build -v -o /notification_service/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /notification_service/build/api / 
COPY --from=build-stage /notification_service/.env / 
EXPOSE 50055
CMD [ "/api" ]