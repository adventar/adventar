package main

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/disintegration/imaging"
)

func main() {
	lambda.Start(Handler)
}

type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	url := request.QueryStringParameters["url"]
	if url == "" {
		return Response{StatusCode: 400, Body: `{"message":"url is required"}`}, nil
	}
	digest := request.PathParameters["digest"]
	ok := verify(digest, url)
	if ok == false {
		return Response{StatusCode: 400, Body: `{"message":"Invalid digest"}`}, nil
	}

	buf, contentType, err := fetchAndResize(url)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: true,
		Body:            base64.StdEncoding.EncodeToString(buf.Bytes()),
		Headers: map[string]string{
			"Content-Type": contentType,
		},
	}

	return resp, nil
}

func fetchAndResize(url string) (*bytes.Buffer, string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-type")
	var src image.Image

	switch contentType {
	case "image/jpeg":
		src, err = jpeg.Decode(response.Body)
	case "image/gif":
		src, err = gif.Decode(response.Body)
	case "image/png":
		src, err = png.Decode(response.Body)
	default:
		return nil, "", fmt.Errorf("Invalid ContentType: %s", contentType)
	}

	if err != nil {
		return nil, "", err
	}

	dst := imaging.Resize(src, 128, 128, imaging.Lanczos)
	buf := new(bytes.Buffer)

	switch contentType {
	case "image/jpeg":
		err = jpeg.Encode(buf, dst, nil)
	case "image/gif":
		err = gif.Encode(buf, dst, nil)
	case "image/png":
		err = png.Encode(buf, dst)
	}

	if err != nil {
		return nil, "", err
	}

	return buf, contentType, nil
}

func verify(digest string, url string) bool {
	h := sha1.New()
	h.Write([]byte(url + os.Getenv("DIGEST_SALT")))

	return digest == fmt.Sprintf("%x", h.Sum(nil))
}
