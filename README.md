# TinyStore

Golang REST API for storing images

Uses sqlite for storage  
Uses `github.com/lithammer/shortuuid` for uuid  
Uses `github.com/gin-gonic/gin` as backend

## Endpoints
### Get image location
#### Request
`GET /image/:uuid`
#### Response
    {
       "location": "/store/KAyWkWM7FJQrd7LHufHSvi.jpg"
    }


### Upload image
##### Request
`POST /upload`  
`curl -i -X POST -F "image=@image.jpg" localhost:8080/upload`
#### Response
    {
        "uuid": "KAyWkWM7FJQrd7LHufHSvi",
        "location": "/store/KAyWkWM7FJQrd7LHufHSvi.jpg",  
        "filetype": "jpg",  
        "size": "9806"  
    }


## TODO

* JPG image compression
* OAuth 2.0
* More upload options

