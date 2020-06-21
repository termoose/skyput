package dappdappgo

const portalUrl = "https://dappdappgo.coolhd.hu/api/add_skylink.php"

// {'error':'wrong_skylink','msg':'The submitted skylink is invalid,
// please submit 'sia:{46 char}' or '{46 char}' via POST.
// Please do NOT submit urls with portal domains.'}

type Response struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}
