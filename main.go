package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"rebuilder/pkg/util"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			b, err := ioutil.ReadFile("list.txt")
			if err != nil {
				util.Log()
			} else {
				if len(b) > 0 {
					fmt.Println(string(b))
					err := ioutil.WriteFile("list.txt", []byte(""), 0666)
					if err != nil {
						util.Log()
					}
					if string(b) == "exit" {
						return
					}
					cmds := strings.Split(string(b), " ")

					out, err := exec.Command("kill", cmds[0]).Output()
					if err != nil {
						util.Log()
						util.SendMail("のぞみんちょ", "info@otft.info", "【ERROR】プロセスの終了に失敗したよ", "このテキストをもとに終了しようとしたけど失敗したよ<br>"+string(b)+"<br><br>"+string(out))
						break
					}
					if len(cmds) < 2 {
						cmds = append(cmds, "", "", "")
					}
					out2, err := exec.Command(cmds[1], cmds[2:]...).Output()
					if err != nil {
						util.Log()
						util.SendMail("のぞみんちょ", "info@otft.info", "【ERROR】プロセスの終了に失敗したよ", "このテキストをもとに終了しようとしたけど失敗したよ<br>"+string(b)+"<br><br>"+string(out))
						break
					}
					log.Println(string(out2))
					util.SendMail("のぞみんちょ", "info@otft.info", "サービスを再起動したよ", "このテキストをもとに起動したよ<br>"+string(b)+"<br><br>"+string(out2))
				}
			}
		}
	}
}
