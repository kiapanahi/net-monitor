package internal

// type NicActivityReporter struct {
// 	Name     string
// 	Activity NicActivity
// }

// // NewNicActivityReporter creates a new NicActivityReporter
// func NewNicActivityReporter(name string) NicActivityReporter {
// 	return NicActivityReporter{
// 		Name:     name,
// 		Activity: newNicActivity(),
// 	}
// }

// func (nar NicActivityReporter) String() string {
// 	return fmt.Sprintf("Reporter for NIC: %s, with capacity: %v", nar.Name, cap(nar.Activity.pipe))
// }

// func (nar NicActivityReporter) Report() string {
// 	for item := range nar.Activity.Report() {
// 		return fmt.Sprintf("NIC: %s\nAt: %s\nReceive:\t%s, Send:\t%s",
// 			nar.Name,
// 			item.Time,
// 			strconv.FormatUint(item.BytesReceived, 10),
// 			strconv.FormatUint(item.BytesSent, 10),
// 		)
// 	}
// 	return "" // Add a return statement here
// }
