package main

import (
	"fmt"
	"github.com/leilei3167/BiliDownloader/cmd"
)

const flag = `
  ___  _  _  _  ___                     
 | _ )(_)| |(_)|   \  ___ __ __ __ _ _  
 | _ \| || || || |) |/ _ \\ V  V /| ' \ 
 |___/|_||_||_||___/ \___/ \_/\_/ |_||_| 
										v1.0.0
`

func main() {
	fmt.Println(flag)
	cmd.Execute()
}
