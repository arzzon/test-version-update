package a
import (
        . "github.com/arzzon/test-version-update/pkg/types"
        log "github.com/sirupsen/logrus"
)

func CreateA() A{
       a1 := A{
           S: "s1",
           I: 1,
	}
       log.WithFields(log.Fields{
    	"Create": "a",
       }).Info("A: ",a1)
       return a1
}

