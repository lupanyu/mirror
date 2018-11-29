package filelist

import (
	"os"
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
)

type userError string

func (e userError) Error() string{
	return e.Message()
}

func (e userError) Message() string{
	return string(e)
}

const  prefix  =  "/list/"
func HandleFilelist(writer http.ResponseWriter,
	request *http.Request) error{
	if strings.Index(request.URL.Path,prefix) !=0 {
		return userError(fmt.Sprintf("path %s must start %s",request.URL.Path,prefix))
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil

}