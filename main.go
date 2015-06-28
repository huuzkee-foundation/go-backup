package main

import (
	"crypto/md5"
	"encoding/hex"
	// "flag"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var rootDir string = "./testData"
var backupDir string = "Z:/testBackup/"
var configsFile string = "./configs.json"

type FileHash struct {
	fullpath       string
	hash           string
	lastModifyTime time.Time
}

type configs struct {
	Src []string
	Dst string
}

var fileHashes []FileHash

func main() {

	//take CLI input
	/*	dirPtr := flag.String("path", ".", " string")
		svrPtr := flag.String("server", "www.softlayer.com", " string")
		frqPtr := flag.Int("bkp-interval", 60, "backup interval in hours")
		comprsPtr := flag.Bool("compression", false, "a bool")
		encrptPtr := flag.Bool("encryption", false, "a bool")
		flag.Parse()
		dir_path := *dirPtr
		server := *svrPtr
		frequency := *frqPtr
		compress := *comprsPtr
		encrypt := *encrptPtr
	*/ //I am not including bakcup run time since
	//having both frequency and backup run time
	///does not make sense

	//	fmt.Println("Hello user your inputs are  ", dir_path, server, frequency, compress, encrypt)

	t0 := time.Now()
	fmt.Printf("\nIt begins at %v", t0)

	parseConfigs()

	/*FO, err := os.Create("backup.txt")
	if err != nil {
		panic(err)
	}
	defer FO.Close()
	*/
	filepath.Walk(rootDir, VisitFile)

	/*	for _, fh := range fileHashes {
			FO.WriteString(fh.fullpath + ", " + fh.hash + ", " + fh.lastModifyTime.String() + "\n")
		}
	*/t1 := time.Now()

	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func MD5OfFile(fullpath string) []byte {
	if contents, err := ioutil.ReadFile(fullpath); err == nil {
		md5sum := md5.New()
		md5sum.Write(contents)
		return md5sum.Sum(nil)
	}
	return nil
}

func VisitFile(fullpath string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !f.IsDir() {
		fullpath, _ = filepath.Abs(fullpath)
		fmt.Println("Copying... " + fullpath)
		FileCopy(fullpath)
		hash := MD5OfFile(fullpath)
		fileHashes = append(fileHashes, FileHash{fullpath, hex.EncodeToString(hash), f.ModTime()})
	}
	return nil
}

// http://stackoverflow.com/a/21061062
func FileCopy(src string) error {
	in, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer in.Close()

	_, filename := filepath.Split(src)

	dst := backupDir + filename

	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return cerr
}

func parseConfigs() {
	file, err := ioutil.ReadFile(configsFile)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	var myConfigs configs
	err = json.Unmarshal(file, &myConfigs)
	if err != nil {
		fmt.Printf("JSON error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Results: %v\n", myConfigs)
}
