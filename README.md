This is ReadME

This is not finished version please don't use it for now

How to use:

How to send photo from file or by it's file ID in telegram
```go
// Send as file from file system
file, _ := os.Open("img.jpg")
rp, err := t.SendPhoto(tb.SendPhotoType{
	ChatID: 280598933,
}, file)

// Send by file ID
rp, err := t.SendPhoto(tb.SendPhotoType{
	ChatID: 280598933,
	Photo:  "FileID that we want to send in string format",
    // FileID taken from Message respond after uploading file to telegram
}, nil)
```

Notice, not all functions was propely tested.