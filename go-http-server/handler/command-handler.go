package handler

import (
	"encoding/json"
	"github.com/zerosuxx/go-http-server/go-http-server/utility"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type CommandHandler struct {
	Shell  *utility.Shell
	Input  *io.PipeReader
	Output *io.PipeWriter
}

func (handler *CommandHandler) Handle(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)

	if len(reqBody) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		log.Println("Request body can't be empty!")

		return
	}

	var commandWithArgs []string
	_ = json.Unmarshal(reqBody, &commandWithArgs)

	if len(commandWithArgs) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		log.Println("Command can't be empty!")

		return
	}

	log.Printf("Command: %v", commandWithArgs)

	go func() {
		if _, err := io.Copy(res, handler.Input); err != nil {
			log.Panic(err)
		}
	}()

	if err := handler.Shell.Run(commandWithArgs[0], commandWithArgs[1:], handler.Output); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	if err := handler.Output.Close(); err != nil {
		log.Panic(err)
	}
}

func CreateCommandHandler() *CommandHandler {
	pipeReader, pipeWriter := io.Pipe()
	return &CommandHandler{
		Shell:  &utility.Shell{},
		Input:  pipeReader,
		Output: pipeWriter,
	}
}
