# GO-YT
GO-YT is a lightweight and user-friendly server built to assist georgexv.ru in the future with downloading YouTube videos.

### Requirements
GO-YT requires Go 1.18 or later to run.

### Installation
First, clone the repository:
```sh
$ git clone https://github.com/Georgee1337/go-yt.git
```
Then, build the project:
```sh
$ cd go-yt
$ go build
```
Finally, run the server:
```sh
$ ./go-yt
$ Starting GO-YT server on :7070...
```

### Usage

GO-YT provides an API endpoint `/get` for downloading YouTube videos. To use it, make a GET request with the `url` parameter set to a valid YouTube video URL:

```sh
http://localhost:7070/get?url=https://www.youtube.com/watch?v=K_9tX4eHztY
```

If the video is not downloaded yet, the server will start downloading it and return a "Download started successfully" message. If the video is already downloaded, the server will return a "Video already downloaded" message.

The videos are downloaded to the `./downloads` directory by default, but this can be changed by modifying the `outputPath` constant in `utils/utils.go`.

**This project is licensed under the GNU General Public License v3.0. See the LICENSE file for details.**



