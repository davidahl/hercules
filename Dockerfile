FROM golang

COPY ./ ./hercules
WORKDIR /hercules

ENTRYPOINT ["go", "run", "main.go"]