package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type UnknownObject map[string]interface{}

func decodeDelim(expectedDelim rune, decoder *json.Decoder) {
	token, err := decoder.Token()
	if err != nil {
		dieWithError(err)
	}
	delim, ok := token.(json.Delim)
	if !ok || rune(delim) != expectedDelim {
		die("Error: Expecting array of objects in stdin\n")
	}
}

func writeSeenId(watchFile *os.File, id string) {
	_, err := fmt.Fprintln(watchFile, id)
	if err != nil {
		dieWithError(err)
	}
}

func getObjId(keyProp string, obj *UnknownObject) string {
	var key string
	switch typedObjKey := ((*obj)[keyProp]).(type) {
	case string:
		key = typedObjKey
	case float64:
		key = fmt.Sprint(typedObjKey)
	case nil:
		die("Error: Got an object without the property '%s' used as key\n", keyProp)
	default:
		die("Error: Got an object with the property '%s' with value that is not a string or number\n", keyProp)
	}

	return key
}

func main() {
	opts, err := parseArgs(os.Args)
	if err != nil {
		printUsage()
		fmt.Println()
		dieWithError(err)
	}

	watchFile, err := os.OpenFile(formatWatchPath(opts.name), os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		dieWithError(err)
	}

	seenIds, err := readLinesMap(watchFile)
	if err != nil {
		dieWithError(err)
	}

	isEmptyWatchFile := len(seenIds) == 0
	decoder := json.NewDecoder(os.Stdin)
	decodeDelim('[', decoder)
	for decoder.More() {
		var obj UnknownObject
		decoder.Decode(&obj)
		objId := getObjId(opts.key, &obj)

		if isEmptyWatchFile {
			writeSeenId(watchFile, objId)
		} else if !seenIds[objId] {
			json, _ := json.Marshal(obj)
			fmt.Println(string(json))
			writeSeenId(watchFile, objId)
		}
	}
	decodeDelim(']', decoder)
}
