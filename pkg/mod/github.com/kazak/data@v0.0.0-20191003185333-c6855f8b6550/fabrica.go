package data

import (
	api "github.com/kazak/grpcapi"
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
)

const (
	separateLengthText = 40
)

//ScanRequestBodyToChank read request line by line and move in chanel
func ScanRequestBodyToChank(scanner *bufio.Scanner, chank chan<- string) {
	scanner.Scan()
	openedString := scanner.Text()

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, openedString) {
			close(chank)
			return
		}
		chank <- text
	}
}

//StartProdactionPort create Port object
func StartProdactionPort(chText <-chan string, chPort chan<- *api.Port) {
	var startReadBody bool = false
	var str []byte

	for {
		textPart, done := <-chText

		if !done {
			close(chPort)
			return
		}
		if "{" == textPart {
			startReadBody = true
		}

		if startReadBody {
			str = append(str, textPart...)
			if (len(str) < separateLengthText) { // we can think about separate port object better
				continue
			}

			var checkJSON = make([]byte, len(str))
			copy(checkJSON, str)

			if bytes.Equal(str[len(str)-2:], []byte("},")) {
				checkJSON[len(checkJSON)-1] = '}'
			}

			if bytes.Equal(checkJSON[len(str)-1:], []byte("}")) {
				var newPorts map[string]*api.Port
				error := json.Unmarshal(checkJSON, &newPorts)

				if (error == nil) {
					for key, port := range newPorts {
						port.PortId = key
						chPort <- port
					}
				}
			}
		}
	}
}
