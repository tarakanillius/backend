//computeImageUrl.go
package utils

import "fmt"

// ComputeImageURL generates a URL for the product image based on its ID.
func ComputeImageURL(barcode string) string {
    var url string
    if len(barcode) <= 8 {
        url = fmt.Sprintf("https://images.openfoodfacts.org/images/products/%s/%d.jpg", barcode, 1)
    } else {
        prefix := barcode[:9]
        lastPart := barcode[9:]
        url = fmt.Sprintf("https://images.openfoodfacts.org/images/products/%s/%s/%s/%s/%d.jpg",
            prefix[:3],
            prefix[3:6],
            prefix[6:],
            lastPart,
            1)
    }
    return url
}
