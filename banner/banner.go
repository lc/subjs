package banner

//Banner prints banner
func Banner() string {
	b := `
	

	███████╗██╗   ██╗██████╗      ██╗███████╗
	██╔════╝██║   ██║██╔══██╗     ██║██╔════╝
	███████╗██║   ██║██████╔╝     ██║███████╗
	╚════██║██║   ██║██╔══██╗██   ██║╚════██║
	███████║╚██████╔╝██████╔╝╚█████╔╝███████║
	╚══════╝ ╚═════╝ ╚═════╝  ╚════╝ ╚══════╝
                                         
	[+] Usage (all urls from file) 1º: $cat urls.txt | subj
	[+] Usage 2º: $cat urls.txt | go run subjs.go
	[+] Usage 3º: (single domain) | go run subjs.go -d https://example.com -o myfile
	
	`
	return b
}
