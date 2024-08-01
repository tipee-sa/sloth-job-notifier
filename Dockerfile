FROM gcr.io/distroless/static:nonroot
WORKDIR /

COPY sloth-job-notifier /sloth-job-notifier
USER nonroot:nonroot

ENTRYPOINT ["/sloth-job-notifier"]
