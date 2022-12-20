package constant

// Import Cloudinary and other necessary libraries
//===================
import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/cloudinary/cloudinary-go/v2/api"
)

func Credentials() *cloudinary.Cloudinary {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true

	return cld
}

func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, "https://cloudinary-devs.github.io/cld-docs-assets/assets/images/butterfly.jpeg", uploader.UploadParams{
		PublicID:       "quickstart_butterfly",
		UniqueFilename: *api.Bool(false),
		Overwrite:      *api.Bool(true)})
	if err != nil {
		fmt.Println("error")
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
}
