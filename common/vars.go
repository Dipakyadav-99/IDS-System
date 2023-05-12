package common

var (
	Email       = ""
	Development = true
	Version     = ""
	Banner      = "\n\nIDS - Intrusion Detection System"

	Usage = `  [buffers] | [options]
	   -c [...] [options]`
	Example = `  teler -c [...] -i /var/log/nginx/access.log
	  [kubectl logs|tail -f|cat] ... | teler -c [...] -x 25
	  `
)
