package imggather

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func ImgGather(inputPath string, wg *sync.WaitGroup){
	if _, err := os.Stat(inputPath + "/unarranged/"); os.IsNotExist(err) {
		mkPath := inputPath + "/unarranged/"
		mkPath = filepath.Clean(mkPath)
		err := os.Mkdir(mkPath, 0777)
		if err != nil{
			log.Fatalln(err)
		}

	}

	 //createdFolder.Write()


	files, err := ioutil.ReadDir(inputPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files{
		if(f.IsDir()){
			gatheringPath := inputPath + "/" + f.Name()
			targetFiles, err := ioutil.ReadDir(gatheringPath)
			if err != nil {
				log.Fatal(err)
			}

			folderLength := len(targetFiles)
			if folderLength < 11{
				for _, f := range targetFiles {
					//fileType := filepath.Base(f.Name())
					imgName := f.Name()
					movedPath := inputPath + "/unarranged/" + imgName

					err := os.Rename(gatheringPath + "/" + f.Name(), movedPath)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}

		}
	}

	folderCleaner(inputPath)
}


func folderCleaner(inputPath string){

	files, err := ioutil.ReadDir(inputPath)
	if err != nil {
		log.Fatalln(err)
	}


	for _, f := range files{
		if f.IsDir(){
			log.Println(f.Name())
			fEmpty, err := ioutil.ReadDir(inputPath + "/" + f.Name())
			if err != nil{
				log.Fatalln(err)
			}

			if len(fEmpty) == 0{
				os.Remove(inputPath + "/" + f.Name())
			}
		}
	}

}
