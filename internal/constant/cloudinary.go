package constant

// Import Cloudinary and other necessary libraries
//===================
import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/spf13/viper"
)

func Credentials() *cloudinary.Cloudinary {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, err := cloudinary.NewFromURL(viper.GetString("cloudinary.url"))
	if err != nil {
		fmt.Println(err)
	}
	cld.Config.URL.Secure = true

	return cld
}

func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context, imageId string, image interface{}) (string, error) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, image, uploader.UploadParams{
		PublicID:       imageId,
		UniqueFilename: *api.Bool(false),
		Overwrite:      *api.Bool(true),
		Folder:         "/Users/Profile Photos",
	})
	if err != nil {
		return "", err
	}

	// Log the delivery URL
	// fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
	return resp.SecureURL, err
}
