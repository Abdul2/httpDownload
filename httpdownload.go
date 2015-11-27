package httpdownload


import (


"fmt"
"io"
"net/http"
"os"
"strings"

)



//Function that downloads file from a url
//assumes that url has a file server
//warning : it does not check if  file with the same name exists [overrides]

func HttpDownload(url string)  {

//figure out the file name
tokens := strings.Split(url, "/")

fileTObeDownloaded := tokens[len(tokens)-1]

//create destination file and open pipe to write to

output, err := os.Create(fileTObeDownloaded)


if err != nil {

fmt.Println("cant create file ", fileTObeDownloaded, "-", err)
return
}

//release once finished
defer output.Close()

//request for the file

resp, err := http.Get(url)

if err != nil {
fmt.Println("Error while downloading", url, "-", err)
return
}

//will stream from url, so defer response closure until we are done with reading the body
defer resp.Body.Close()

//body copied into file
writtenBytes, err := io.Copy(output, resp.Body)

if err != nil {

fmt.Println("cant save file", url, "-", err)

return
}


fmt.Println("Downaloded file size " ,writtenBytes )

}
