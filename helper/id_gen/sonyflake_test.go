package id_gen

//import (
//	"github.com/sony/sonyflake"
//	"testing"
//	"time"
//)
//
//func TestGenSonyFlakeId(t *testing.T) {
//	t, _ := time.Parse("2006-01-02", "2018-01-01")
//	settings := sonyflake.Settings{
//		StartTime:      t,
//		MachineID:      getMachineID,
//		CheckMachineID: checkMachineID,
//	}
//
//	sf := sonyflake.NewSonyflake(settings)
//	id, err := sf.NextID()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	fmt.Println(id)
//}
