package fox

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/e-gov/fox/util"
	log "github.com/Sirupsen/logrus"
)

func getFileName(uuid string) string {
	return util.GetConfig().Storage.Filepath + uuid
}

// StoreFox persists the fox instance to somewhere
func StoreFox(fox Fox, uuid string) UUID {

	err := os.MkdirAll(getFileName(""), 0300)

	if err != nil {
		panic(err)
	}

	fname := getFileName(uuid)
	f, err := os.Create(fname)

	if err != nil {
		log.WithFields(log.Fields{
			"path": fname,
		}).Panic("Failed to write fox data to file")
	}

	defer f.Close()

	// Make sure the file name matches the uuid in the structure
	fox.Uuid = uuid
	if err := json.NewEncoder(f).Encode(fox); err != nil {
		panic(err)
	}

	return UUID{Uuid: uuid}
}

func ReadFox(uuid string) (Fox, error) {
	var fox Fox
	data, err := ioutil.ReadFile(getFileName(uuid))

	if err != nil {
		return fox, err
	}

	if err := json.Unmarshal(data, &fox); err != nil {
		return fox, err
	}

	return fox, nil
}

func GetFoxes() ([]Fox, error) {
	var foxes []Fox

	foxes = make([]Fox, 0)
	fname := util.GetConfig().Storage.Filepath
	files, _ := ioutil.ReadDir(fname)
	for _, f := range files {
		fox, err := ReadFox(f.Name())
		if err != nil {
			return foxes, err
		}
		foxes = append(foxes, fox)
	}

	return foxes, nil
}

func FoxExists(uuid string) bool {
	if _, err := os.Stat(getFileName(uuid)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Deletes the fox, if it exists. Does nothing if it does not
func DeleteFoxFromStorage(uuid string) {

	if !FoxExists(uuid) {
		// Exit quietly, if the fox is not there
		return
	}

	// Attempt to remove the file if it is there. Panic if it fails
	if err := os.Remove(getFileName(uuid)); err != nil {
		panic(err)
	}
}
