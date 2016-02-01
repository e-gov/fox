package main

import(
	"os"
	"encoding/json"
	"gopkg.in/gcfg.v1"
	"io/ioutil"
    "log"
)
func getFileName(uuid string) string{
	return getConfig().Storage.Filepath + uuid
}

func getConfig() Config{
	var cfg Config

	// Read configuration
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil{
		panic(err)
	}
	return cfg
}

// Persist the fox instance to somewhere
func StoreFox(fox Fox, uuid string) UUID {
	f, err := os.Create(getFileName(uuid))

	if err != nil{
		panic(err)
	}

	defer f.Close()

	// Make sure the file name matches the uuid in the structure
	fox.Uuid = uuid
	if err:=json.NewEncoder(f).Encode(fox); err != nil{
		panic(err)
	}

	return UUID{Uuid: uuid}
}

func ReadFox(uuid string) (Fox, error){
	var fox Fox
	data, err := ioutil.ReadFile(getFileName(uuid))

	if err != nil{
		return fox, err
	}

	if err := json.Unmarshal(data, &fox); err != nil{
		return fox, err
	}

	return fox, nil
}

func GetFoxes()([]Fox, error){
    var foxes []Fox
    
    foxes = make([]Fox, 0)
    
    files, _ := ioutil.ReadDir(getConfig().Storage.Filepath)
    for _, f := range files {
	log.Print(f.Name())
	fox, err := ReadFox(f.Name())
	if err != nil{
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
func DeleteFoxFromStorage(uuid string){

	if !FoxExists(uuid){
		// Exit quietly, if the fox is not there
		return
	}

	// Attempt to remove the file if it is there. Panic if it fails
	if err:=os.Remove(getFileName(uuid)); err != nil{
		panic(err)
	}
}
