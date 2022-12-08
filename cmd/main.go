package main
import (
        . "github.com/arzzon/test-version-update/pkg/a"
        log "github.com/sirupsen/logrus"	
)
func main(){
        log.WithFields(log.Fields{
        "Version": "V1",
       }).Info("")
	a1 := CreateA()
	log.WithFields(log.Fields{
        "Reading": "a",
       }).Info("A: ",a1)
}
