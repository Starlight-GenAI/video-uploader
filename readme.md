# Video uploader service

## To start application
1. Create config.yaml at config folder by using config.example.yaml as reference
2. Add a Google Cloud service account in JSON format at root folder
3. Run
```sh
go build ./cmd/main.go  && ./main
```
4. The application will be running on 127.0.0.1:8080

## To start application using docker
1. Create config.yaml by using config.example.yaml as reference
2. Add a Google Cloud service account in JSON format by updating the Dockerfile at line 17.
3. Run
```sh
docker build . -t video-uploader
```
4. Run
```sh
docker run -p 8080:8080  -v ./config.yaml:/config video-uploader -d
```
5. The application will be running on 127.0.0.1:8080

## API document
### 1. POST /upload <br />
description: endpoint for uploading video
#### Request body
```json
{
    "video_url": "",
    "is_use_subtitle": false,
    "user_id": "gaVbn6cxnYOwoxlCfDInZwgwA262"
}
```
#### request body detail
```text
video_url: youtube url
is_use_subtitle: If true, server will use only subtitles to generate trip. If false, server will use subtitles and video in MP4 format
user_id: random string
```
#### Response
```json
{
    "queue_id": "0f33fd7594e1486aa63aa4643d4fa2ef"
}
```
#### Response detail
```text
queue_id: id for a generated trip
```