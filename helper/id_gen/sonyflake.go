package id_gen

//func getMachineID() (uint16, error) {
//	var machineID uint16
//	var err error
//	machineID = readMachineIDFromLocalFile()
//	if machineID == 0 {
//		machineID, err = generateMachineID()
//		if err != nil {
//			return 0, err
//		}
//	}
//
//	return machineID, nil
//}
//
//func checkMachineID(machineID uint16) bool {
//	saddResult, err := saddMachineIDToRedisSet()
//	if err != nil || saddResult == 0 {
//		return true
//	}
//
//	err := saveMachineIDToLocalFile(machineID)
//	if err != nil {
//		return true
//	}
//
//	return false
//}
