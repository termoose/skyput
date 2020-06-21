package dappdappgo

const portalUrl = "https://dappdappgo.coolhd.hu/api/add_skylink.php"

type Response struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}
