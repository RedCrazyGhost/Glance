/**
 @author: RedCrazyGhost
 @date: 2023/2/24

**/

package core

import (
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T) {
	v := viper.New()
	v.Set("test", rune('#'))
	err := v.SafeWriteConfigAs("./ok.yaml")
	if err != nil {
		return
	}
}
