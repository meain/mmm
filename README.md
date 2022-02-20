<img src="https://r.meain.io/m/eepvoDTfPaOmQqOVIcWPxiJT" align="right" height="500" />

# mmm: matrix media manager

> Perpetual WIP project

*Use matrix as storage and get publicly accessible urls*

### What is going on

When you send an image/video/file on matrix, matrix does give you a link which looks something like `https://matrix.org/_matrix/media/r0/download/matrix.org/<id>` at which you can download the media that you shared. This is available publicly if your room is public (you can keep the room unlisted). Now, we take this id and push it through a url shortner and bam, you have a short url to a resource.

The best part is that upload, storage and auth are handled by matrix clients. We just do some redirections. :^)

*In fact the image that you seen in this README is added this way*

> Use [meain/sirus](https://github.com/meain/sirus) for automatically redirecting everything with a prefix. Check "sub" mode there.
