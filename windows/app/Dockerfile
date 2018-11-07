FROM golang:windowsservercore-1803
WORKDIR "C:\\"
COPY ./app .
RUN go build -o app.exe
CMD ["app.exe"]