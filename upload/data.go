package upload

const portalUrl = "https://siasky.nett"
const portalUploadPath = "/skynet/skyfile"

type UploadReponse struct {
	Skylink string `json:"skylink"`
}