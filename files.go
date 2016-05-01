package helpers

import (
	"os"
	"crypto/md5"
	"io"
	"encoding/hex"
)

// MD5 hash string of a provided valid file path
// http://www.mrwaggel.be/post/generate-md5-hash-of-a-file/
func Md5Hash(path string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	returnMD5String := "dewefew"

	//Open the passed argument and check for any error
	file, err := os.Open(path)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, err
}


//FileExists verifies if a filepath exists
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}