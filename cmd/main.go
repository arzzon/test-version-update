package main
import (
        . "github.com/arzzon/test-version-update/v2/pkg/a"
        log "github.com/sirupsen/logrus"	
)
func main(){
        log.WithFields(log.Fields{
        "Version": "V2.2.1",
       }).Info("")
	a1 := CreateA()
	log.WithFields(log.Fields{
        "Reading": "a",
       }).Info("A: ",a1)
}
